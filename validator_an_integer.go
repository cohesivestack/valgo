package valgo

import (
	"strconv"
)

func (val *Value) IsInteger() bool {
	if val.isInteger == nil {
		val.isInteger = boolPointer(false)
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
			val.isInteger = boolPointer(true)
		case string:
			_, err := strconv.ParseInt(val.absolute.(string), 10, 64)
			val.isInteger = boolPointer(err == nil)
		}
	}
	return *val.isInteger
}

func (v *Validator) AnInteger(template ...string) *Validator {
	if !v.assert(v.currentValue.IsInteger()) {
		v.invalidate("an_integer",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
