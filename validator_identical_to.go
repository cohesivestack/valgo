package valgo

func (valA *Value) IsIdenticalTo(value interface{}) bool {
	valBOriginal := value

	// Value.IsComparableType is not used since we need to check the original
	// values
	if !isComparableType(valA.original) || !isComparableType(valBOriginal) {
		return false
	}

	return valA.original == valBOriginal
}

func (v *Validator) IdenticalTo(value interface{}, template ...string) *Validator {
	if !v.assert(v.currentValue.IsIdenticalTo(value)) {
		v.invalidate("identical_to",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": convertToString(value)}, template)
	}

	v.resetNegative()

	return v
}
