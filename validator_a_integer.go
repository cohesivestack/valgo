package valgo

import (
	"strconv"
)

func anInteger(value interface{}) bool {
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
	case string:
		_, err := strconv.ParseInt(value.(string), 10, 64)
		return err == nil
	}
	return false
}

func (validator *Validator) AnInteger(template ...string) *Validator {
	if !anInteger(validator.currentValue) {
		validator.invalidate("an_integer",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
