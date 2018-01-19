package valgo

func (validator *Validator) Empty(template ...string) *Validator {
	value := validator.ensureString()
	if len(value) > 0 {
		validator.valid = false
		validator.invalidate("empty", map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}

func (validator *Validator) NotEmpty(template ...string) *Validator {
	value := validator.ensureString()
	if len(value) == 0 {
		validator.valid = false
		validator.invalidate("not_empty", map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
