package valgo

func (value *Value) IsIntegerType() bool {
	if value.isIntegerType == nil {
		value.isIntegerType = boolPointer(false)
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
			int64:
			value.isIntegerType = boolPointer(true)
		}
	}
	return *value.isIntegerType
}

func (validator *Validator) AnIntegerType(template ...string) *Validator {
	if !validator.currentValue.IsIntegerType() {
		validator.invalidate("an_integer_type",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
