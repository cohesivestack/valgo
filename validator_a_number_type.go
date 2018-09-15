package valgo

func (val *Value) IsNumberType() bool {
	if val.isNumberType == nil {
		val.isNumberType = boolPointer(false)
		switch val.absolute.(type) {
		case uint,
			uint8,
			uint16,
			uint32,
			uint64,
			int,
			int8,
			int16,
			int32,
			int64,
			float32,
			float64:
			val.isNumberType = boolPointer(true)
		}
	}

	return *val.isNumberType
}

func (v *Validator) ANumberType(template ...string) *Validator {
	if !v.assert(v.currentValue.IsNumberType()) {
		v.invalidate("a_number_type",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
