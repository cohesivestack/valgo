package valgo

func (value *Value) IsString() bool {
	if value.isString == nil {
		value.isString = boolPointer(false)
		switch value.absolute.(type) {
		case string:
			value.isString = boolPointer(true)
		}
	}

	return *value.isString
}

func (validator *Validator) AString(template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsString()) {
		validator.invalidate("a_string",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}

	validator.resetNegative()

	return validator
}
