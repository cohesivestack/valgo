package valgo

func (val *Value) IsIntegerType() bool {
	if val.isIntegerType == nil {
		val.isIntegerType = boolPointer(false)
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
			int64:
			val.isIntegerType = boolPointer(true)
		}
	}
	return *val.isIntegerType
}

func (v *Validator) AnIntegerType(template ...string) *Validator {
	if !v.assert(v.currentValue.IsIntegerType()) {
		v.invalidate("an_integer_type",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
