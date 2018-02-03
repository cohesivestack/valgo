package valgo

func aNumberType(value interface{}) bool {
	switch value.(type) {
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
		return true
	}
	return false
}

func (validator *Validator) ANumberType(template ...string) *Validator {
	if !aNumberType(validator.currentValue) {
		validator.invalidate("a_number_type",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
