package valgo

import (
	"strconv"
)

func (value *Value) IsNumber() bool {
	if value.isNumber == nil {
		value.isNumber = boolPointer(false)
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
			value.isNumber = boolPointer(true)
		case string:
			_, err := strconv.ParseFloat(value.absolute.(string), 64)
			value.isNumber = boolPointer(err == nil)
		}
	}
	return *value.isNumber
}

func (validator *Validator) ANumber(template ...string) *Validator {
	if !validator.currentValue.IsNumber() {
		validator.invalidate("a_number",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
