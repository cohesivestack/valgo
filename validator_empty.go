package valgo

import (
	"reflect"
)

func (val *Value) IsEmpty() bool {
	if val.isEmpty == nil {
		val.isEmpty = boolPointer(false)

		if val.absolute == nil {
			val.isEmpty = boolPointer(true)
		} else if val.IsNumberType() {
			val.isEmpty = boolPointer(val.AsFloat64() == 0)
		} else if val.IsString() {
			val.isEmpty = boolPointer(len(val.AsString()) == 0)
		} else {
			_val := reflect.ValueOf(val.absolute)
			switch _val.Kind() {
			case reflect.Slice, reflect.Map:
				val.isEmpty = boolPointer(_val.Len() == 0)
			}
		}
	}
	return *val.isEmpty
}

func (v *Validator) Empty(template ...string) *Validator {
	if !v.assert(v.currentValue.IsEmpty()) {
		v.invalidate(
			"empty",
			map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
