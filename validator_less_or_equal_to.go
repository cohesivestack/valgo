package valgo

func (valueA *Value) IsLessOrEqualTo(value interface{}) bool {
	valueB := NewValue(value)

	if valueA.absolute == nil || valueB.absolute == nil {
		return false
	}

	if (valueA.IsNumber() && valueB.IsNumberType()) ||
		(valueB.IsNumber() && valueA.IsNumberType()) ||
		(valueA.IsNumberType() && valueB.IsNumberType()) {
		return valueA.AsFloat64() <= valueB.AsFloat64()
	}

	if valueA.IsString() && valueB.IsString() {
		return valueA.AsString() <= valueB.AsString()
	}

	return false
}

func (validator *Validator) LessOrEqualTo(value interface{}, template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsLessOrEqualTo(value)) {
		validator.invalidate("less_or_equal_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}

	validator.resetNegative()

	return validator
}
