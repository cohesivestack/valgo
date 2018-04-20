package valgo

func (valueA *Value) IsIdenticalTo(value interface{}) bool {
	valueBOriginal := value

	// Value.IsComparableType is not used since we need to check the original
	// values
	if !isComparableType(valueA.original) || !isComparableType(valueBOriginal) {
		return false
	}

	return valueA.original == valueBOriginal
}

func (validator *Validator) IdenticalTo(value interface{}, template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsIdenticalTo(value)) {
		validator.invalidate("identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}

	validator.resetNegative()

	return validator
}
