package valgo

import (
	"reflect"
)

func empty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return value.(float64) == 0
	case string:
		return len(value.(string)) == 0
	default:
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			return len(value.([]interface{})) == 0
		case reflect.Map:
			return len(value.(map[interface{}]interface{})) == 0
		}
	}

	return false
}

func (validator *Validator) Empty(template ...string) *Validator {
	if !empty(validator.currentValue) {
		validator.invalidate(
			"empty",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}

func (validator *Validator) NotEmpty(template ...string) *Validator {
	if empty(validator.currentValue) {
		validator.invalidate(
			"not_empty",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
