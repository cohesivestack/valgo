package valgo

func anIntegerType(value interface{}) bool {
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
		int64:
		return true
	}
	return false
}

func (validator *Validator) AnIntegerType(template ...string) *Validator {
	if !anIntegerType(validator.currentValue) {
		validator.invalidate("an_integer_type",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
