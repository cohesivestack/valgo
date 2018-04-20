package valgo

func (value *Value) IsNumberType() bool {
	if value.isNumberType == nil {
		value.isNumberType = boolPointer(false)
		switch value.absolute.(type) {
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
			value.isNumberType = boolPointer(true)
		}
	}

	return *value.isNumberType
}

func (validator *Validator) ANumberType(template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsNumberType()) {
		validator.invalidate("a_number_type",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}

	validator.resetNegative()

	return validator
}
