package valgo

import (
	"strings"
)

func blank(value interface{}) bool {
	_value := strings.TrimSpace(convertToString(value))
	return len(_value) == 0
}

func (validator *Validator) Blank(template ...string) *Validator {
	if !blank(validator.currentValue) {
		validator.invalidate("blank", map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}

func (validator *Validator) NotBlank(template ...string) *Validator {
	if blank(validator.currentValue) {
		validator.invalidate("not_blank", map[string]interface{}{
			"Title": validator.currentTitle}, template)
	}
	return validator
}
