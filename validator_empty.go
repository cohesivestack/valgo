package valgo

import (
	"reflect"
)

func (value *Value) IsEmpty() bool {
	if value.isEmpty == nil {
		value.isEmpty = boolPointer(false)

		if value.absolute == nil {
			value.isEmpty = boolPointer(true)
		} else if value.IsNumberType() {
			value.isEmpty = boolPointer(value.AsFloat64() == 0)
		} else if value.IsString() {
			value.isEmpty = boolPointer(len(value.AsString()) == 0)
		} else {
			_value := reflect.ValueOf(value.absolute)
			switch _value.Kind() {
			case reflect.Slice, reflect.Map:
				value.isEmpty = boolPointer(_value.Len() == 0)
			}
		}
	}
	return *value.isEmpty
}

func (validator *Validator) Empty(template ...string) *Validator {
	if !validator.currentValue.IsEmpty() {
		validator.invalidate(
			"empty",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}

func (validator *Validator) NotEmpty(template ...string) *Validator {
	if validator.currentValue.IsEmpty() {
		validator.invalidate(
			"not_empty",
			map[string]interface{}{"Title": validator.currentTitle}, template)
	}
	return validator
}
