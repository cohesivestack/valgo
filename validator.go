package valgo

import (
	"fmt"
	"strings"

	"github.com/valyala/fasttemplate"
)

type Validator struct {
	currentValue *Value
	currentTitle string
	currentName  string
	currentValid bool
	currentIndex int

	currentNegative bool

	_locale        locale
	valid          bool
	errorValidator *ErrorValidator
}

func (validator *Validator) Is(value interface{}) *Validator {

	validator.closeCurrentError()

	validator.currentIndex += 1
	validator.currentValue = NewValue(value)
	validator.currentValid = true
	validator.currentName = fmt.Sprintf("value%v", validator.currentIndex)
	validator.currentTitle = validator.currentName

	return validator
}

func (validator *Validator) Not() *Validator {
	validator.currentNegative = true

	return validator
}

func (validator *Validator) resetNegative() {
	validator.currentNegative = false
}

func (validator *Validator) Named(name string) *Validator {
	validator.currentName = name
	if !validator.currentValid {
		validator.errorValidator.currentError.Name = name
	}

	return validator
}

func (validator *Validator) Titled(title string) *Validator {
	validator.currentTitle = title

	return validator
}

func (validator *Validator) Valid() bool {
	return validator.valid
}

func (validator *Validator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *Validator {

	customValidator := CustomValidator{
		validator: validator,
	}

	if len(template) > 0 {
		function(&customValidator, template[0])
	} else {
		function(&customValidator)
	}

	validator.resetNegative()

	return validator
}

func (validator *Validator) ErrorItems() []ErrorItem {
	if validator.valid {
		return []ErrorItem{}
	}

	return validator.errorValidator.Items()
}

func (validator *Validator) Error() error {
	return validator.errorValidator
}

func (v *Validator) AddErrorToNamed(name string, message string) *Validator {
	return v.Is(nil).Named(name).WithError(message)
}

func (v *Validator) WithError(messageTemplate string) *Validator {
	v.invalidate("", map[string]interface{}{
		"Title": v.currentTitle,
		"Value": v.currentValue}, []string{messageTemplate})

	return v
}

func (validator *Validator) assert(value bool) bool {
	return validator.currentNegative != value
}

func (validator *Validator) invalidate(
	key string, values map[string]interface{}, templateString []string) {

	if validator.currentNegative {
		key = "not_" + key
	}

	validator.valid = false
	var _templateString string
	if len(templateString) > 0 {
		_templateString = templateString[0]
	} else if ts, ok := validator._locale.Messages[key]; ok {
		_templateString = ts
	} else if len(strings.TrimSpace(key)) == 0 {
		_templateString = "ERROR: MISSING MESSAGE KEY OR TEMPLATE STRING!"
	} else {
		_templateString = fmt.Sprintf(
			"ERROR: THERE IS NOT A MESSAGE WITH THE KEY \"%s\"!", key)
	}

	template := fasttemplate.New(_templateString, "{{", "}}")
	message := template.ExecuteString(values)

	if validator.errorValidator == nil {
		validator.errorValidator = &ErrorValidator{
			items: []*ErrorItem{},
		}
	}

	if validator.errorValidator.currentError == nil {
		validator.errorValidator.currentError = &ErrorItem{
			Name:     validator.currentName,
			Messages: []string{},
		}
	}

	validator.errorValidator.currentError.Messages = append(
		validator.errorValidator.currentError.Messages, message)

	validator.currentValid = false
	validator.valid = false
}

func (v *Validator) closeCurrentError() {
	if !v.currentValid {
		lastError := v.errorValidator.currentError
		v.errorValidator.currentError = nil

		// If already exist an ErrorItem with the name then reuse the error item
		for _, item := range v.errorValidator.items {
			if item.Name == lastError.Name {
				item.Messages = append(item.Messages, lastError.Messages...)
				return
			}
		}

		// Is was not found an ErrorItem with the same name then add it
		v.errorValidator.items = append(v.errorValidator.items, lastError)
	}
}
