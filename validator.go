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

func (v *Validator) Is(value interface{}) *Validator {

	v.closeCurrentError()

	v.currentIndex += 1
	v.currentValue = NewValue(value)
	v.currentValid = true
	v.currentName = fmt.Sprintf("value%v", v.currentIndex)
	v.currentTitle = v.currentName

	return v
}

func (v *Validator) Not() *Validator {
	v.currentNegative = true

	return v
}

func (v *Validator) resetNegative() {
	v.currentNegative = false
}

func (v *Validator) Named(name string) *Validator {
	v.currentName = name
	if !v.currentValid {
		v.errorValidator.currentError.Name = name
	}

	return v
}

func (v *Validator) Titled(title string) *Validator {
	v.currentTitle = title

	return v
}

func (v *Validator) Valid() bool {
	return v.valid
}

func (v *Validator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *Validator {

	customValidator := CustomValidator{
		validator: v,
	}

	if len(template) > 0 {
		function(&customValidator, template[0])
	} else {
		function(&customValidator)
	}

	v.resetNegative()

	return v
}

func (v *Validator) ErrorItems() []ErrorItem {
	if v.valid {
		return []ErrorItem{}
	}

	return v.errorValidator.Items()
}

func (v *Validator) Error() error {
	return v.errorValidator
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

func (v *Validator) assert(value bool) bool {
	return v.currentNegative != value
}

func (v *Validator) invalidate(
	key string, values map[string]interface{}, templateString []string) {

	if v.currentNegative {
		key = "not_" + key
	}

	v.valid = false
	var _templateString string
	if len(templateString) > 0 {
		_templateString = templateString[0]
	} else if ts, ok := v._locale.Messages[key]; ok {
		_templateString = ts
	} else if len(strings.TrimSpace(key)) == 0 {
		_templateString = "ERROR: MISSING MESSAGE KEY OR TEMPLATE STRING!"
	} else {
		_templateString = fmt.Sprintf(
			"ERROR: THERE IS NOT A MESSAGE WITH THE KEY \"%s\"!", key)
	}

	template := fasttemplate.New(_templateString, "{{", "}}")
	message := template.ExecuteString(values)

	if v.errorValidator == nil {
		v.errorValidator = &ErrorValidator{
			items: []*ErrorItem{},
		}
	}

	if v.errorValidator.currentError == nil {
		v.errorValidator.currentError = &ErrorItem{
			Name:     v.currentName,
			Messages: []string{},
		}
	}

	v.errorValidator.currentError.Messages = append(
		v.errorValidator.currentError.Messages, message)

	v.currentValid = false
	v.valid = false
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
