package valgo

func (val *Value) IsString() bool {
	if val.isString == nil {
		val.isString = boolPointer(false)
		switch val.absolute.(type) {
		case string:
			val.isString = boolPointer(true)
		}
	}

	return *val.isString
}

func (v *Validator) AString(template ...string) *Validator {
	if !v.assert(v.currentValue.IsString()) {
		v.invalidate("a_string",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
