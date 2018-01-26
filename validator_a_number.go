package valgo

import (
	"strconv"
)

func aNumber(value interface{}) bool {
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
	case string:
		_, err := strconv.ParseFloat(value.(string), 64)
		return err == nil
	}
	return false
}

func (validator *Validator) ANumber(template ...string) *Validator {
	if !aNumber(validator.currentValue) {
		validator.invalidate("a_number",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
