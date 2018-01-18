package valgo

import (
	"strings"
)

func (validator *Validator) Blank() *Validator {
	value := strings.Trim(validator.ensureString(), " ")

	if len(value) > 0 {
		validator.valid = false
		validator.invalidate("blank", map[string]interface{}{"Title": validator.currentTitle})
	}
	return validator
}

func (validator *Validator) NotBlank() *Validator {
	value := strings.Trim(validator.ensureString(), " ")

	if len(value) == 0 {
		validator.valid = false
		validator.invalidate("not_blank", map[string]interface{}{
			"Title": validator.currentTitle})
	}
	return validator
}
