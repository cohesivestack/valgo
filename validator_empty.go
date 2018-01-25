package valgo

import (
	"reflect"
)

func empty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case uint:
		return value.(uint) == 0
	case uint8:
		return value.(uint8) == 0
	case uint16:
		return value.(uint16) == 0
	case uint32:
		return value.(uint32) == 0
	case uint64:
		return value.(uint64) == 0
	case int:
		return value.(int) == 0
	case int8:
		return value.(int8) == 0
	case int16:
		return value.(int16) == 0
	case int32:
		return value.(int32) == 0
	case int64:
		return value.(int64) == 0
	case float32:
		return value.(float32) == 0
	case float64:
		return value.(float64) == 0
	case string:
		return len(value.(string)) == 0
	default:
		_value := reflect.ValueOf(value)
		switch _value.Kind() {
		case reflect.Slice, reflect.Map:
			return _value.Len() == 0
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
