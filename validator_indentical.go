package valgo

import (
	"reflect"
)

func indenticalTo(valueA interface{}, valueB interface{}) bool {
	for _, v := range []interface{}{valueA, valueB} {
		_type := reflect.TypeOf(v)
		switch _type.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			return false
		}
	}

	if valueA != valueB {
		return false
	}
	return true
}

func (validator *Validator) IdenticalTo(value interface{}, template ...string) *Validator {

	if !indenticalTo(validator.currentValue, value) {
		validator.invalidate("identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotIdenticalTo(value interface{}, template ...string) *Validator {

	if indenticalTo(validator.currentValue, value) {
		validator.invalidate("not_identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
