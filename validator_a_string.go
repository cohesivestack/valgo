package valgo

func aString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	}
	return false
}

func (validator *Validator) AString(template ...string) *Validator {
	if !aString(validator.currentValue) {
		validator.invalidate("a_string",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
