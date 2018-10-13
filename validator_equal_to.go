package valgo

func IsEqualTo(valueA interface{}, valueB interface{}) bool {
	return valueA == valueB
}

func (v *Validator) EqualTo(value interface{}, template ...string) *Validator {
	if !v.assert(IsEqualTo(v.currentValue, value)) {
		v.invalidate("equal_to", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}
