package valgo

func (valA *Value) IsLessOrEqualTo(value interface{}) bool {
	valB := NewValue(value)

	if valA.absolute == nil || valB.absolute == nil {
		return false
	}

	if (valA.IsNumber() && valB.IsNumberType()) ||
		(valB.IsNumber() && valA.IsNumberType()) ||
		(valA.IsNumberType() && valB.IsNumberType()) {
		return valA.AsFloat64() <= valB.AsFloat64()
	}

	if valA.IsString() && valB.IsString() {
		return valA.AsString() <= valB.AsString()
	}

	return false
}

func (v *Validator) LessOrEqualTo(value interface{}, template ...string) *Validator {
	if !v.assert(v.currentValue.IsLessOrEqualTo(value)) {
		v.invalidate("less_or_equal_to",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": convertToString(value)}, template)
	}

	v.resetNegative()

	return v
}
