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
	currentError *ErrorItem

	currentNegative bool

	_locale locale
	valid   bool
	errors  []*ErrorItem
}

func (validator *Validator) Is(value interface{}) *Validator {

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

func (validator *Validator) ErrorItems() []*ErrorItem {
	return validator.errors
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

	if validator.currentError == nil {
		validator.currentError = &ErrorItem{
			Name:  validator.currentName,
			Title: validator.currentTitle,
			Value: validator.currentValue,
		}

		validator.currentError.Messages = []string{message}
		validator.currentValid = false
		validator.valid = false

		if validator.errors == nil {
			validator.errors = []*ErrorItem{validator.currentError}
		} else {
			validator.errors = append(validator.errors, validator.currentError)
		}
	} else {
		validator.currentError.Messages = append(
			validator.currentError.Messages, message)
	}
}
