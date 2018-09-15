package valgo

func (valA *Value) IsGreaterThan(value interface{}) bool {
	valB := NewValue(value)

	if valA.absolute == nil || valB.absolute == nil {
		return false
	}

	if (valA.IsNumber() && valB.IsNumberType()) ||
		(valB.IsNumber() && valA.IsNumberType()) ||
		(valA.IsNumberType() && valB.IsNumberType()) {
		return valA.AsFloat64() > valB.AsFloat64()
	}

	if valA.IsString() && valB.IsString() {
		return valA.AsString() > valB.AsString()
	}

	return false
}

func (v *Validator) GreaterThan(val interface{}, template ...string) *Validator {
	if !v.assert(v.currentValue.IsGreaterThan(val)) {
		v.invalidate("greater_than",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": convertToString(val)}, template)
	}

	v.resetNegative()

	return v
}
