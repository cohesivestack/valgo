package valgo

func IsInSlice(value interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func (v *Validator) InSlice(slice []interface{}, template ...string) *Validator {
	if !v.assert(IsInSlice(v.currentValue, slice)) {
		v.invalidate("in_slice", map[string]interface{}{
			"title": v.currentTitle,
			"value": v.currentValue}, template...)
	}

	v.resetNegative()

	return v
}
