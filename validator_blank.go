package valgo

import (
	"strings"
)

func (value *Value) IsBlank() bool {
	if value.isBlank == nil {
		_value := strings.TrimSpace(value.AsString())
		value.isBlank = boolPointer(len(_value) == 0)
	}
	return *value.isBlank
}

func (validator *Validator) Blank(template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsBlank()) {
		validator.invalidate("blank", map[string]interface{}{"Title": validator.currentTitle}, template)
	}

	validator.resetNegative()

	return validator
}
