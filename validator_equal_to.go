package valgo

func (validator *Validator) EqualTo(value interface{}, template ...string) *Validator {
	if validator.currentValue != value {
		validator.invalidate("equal_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotEqualTo(value interface{}, template ...string) *Validator {
	if validator.currentValue == value {
		validator.invalidate("not_equal_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
