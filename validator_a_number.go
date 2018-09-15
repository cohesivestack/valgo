package valgo

import (
	"strconv"
)

func (val *Value) IsNumber() bool {
	if val.isNumber == nil {
		val.isNumber = boolPointer(false)
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
			val.isNumber = boolPointer(true)
		case string:
			_, err := strconv.ParseFloat(val.absolute.(string), 64)
			val.isNumber = boolPointer(err == nil)
		}
	}
	return *val.isNumber
}

func (v *Validator) ANumber(template ...string) *Validator {
	if !v.assert(v.currentValue.IsNumber()) {
		v.invalidate("a_number",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
