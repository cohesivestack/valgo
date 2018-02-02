package valgo

import (
	"reflect"
)

func (validator *Validator) IdenticalTo(value interface{}, template ...string) *Validator {

	isComparableType := true

	for _, v := range []interface{}{validator.currentValue, value} {
		_type := reflect.TypeOf(v)
		switch _type.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			isComparableType = false
		}
	}

	if !isComparableType || validator.currentValue != value {
		validator.invalidate("identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotIdenticalTo(value interface{}, template ...string) *Validator {
	if validator.currentValue == value {
		validator.invalidate("not_identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
