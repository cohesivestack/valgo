package valgo

func IsEmpty(value string) bool {
	return len(value) == 0
}

func (v *Validator) Empty(template ...string) *Validator {
	if v.currentDataType != DataTypeString {
		panic("Empty validator requires a string as value")
	} else if !v.assert(IsEmpty(v.currentValue.(string))) {
		v.invalidate("empty", map[string]interface{}{"Title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}
