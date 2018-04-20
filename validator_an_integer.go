package valgo

import (
	"strconv"
)

func (value *Value) IsInteger() bool {
	if value.isInteger == nil {
		value.isInteger = boolPointer(false)
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
			value.isInteger = boolPointer(true)
		case string:
			_, err := strconv.ParseInt(value.absolute.(string), 10, 64)
			value.isInteger = boolPointer(err == nil)
		}
	}
	return *value.isInteger
}

func (validator *Validator) AnInteger(template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsInteger()) {
		validator.invalidate("an_integer",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}

	validator.resetNegative()

	return validator
}
