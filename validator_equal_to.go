package valgo

import (
	"reflect"
)

func (val *Value) IsEqualTo(value interface{}) bool {
	valB := NewValue(value)
	if val.IsComparableType() && valB.IsComparableType() && val.absolute == valB.absolute {
		return true
	}

	// if previous test was not true and one value is nil then just return false
	if val.absolute == nil || valB.absolute == nil {
		return false
	}

	if val.IsNumber() && valB.IsNumber() {
		return val.AsFloat64() == valB.AsFloat64()
	}

	return reflect.DeepEqual(val.absolute, valB.absolute)
}

func (v *Validator) EqualTo(value interface{}, template ...string) *Validator {
	if !v.assert(v.currentValue.IsEqualTo(value)) {
		v.invalidate("equivalent_to",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": convertToString(value)}, template)
	}

	v.resetNegative()

	return v
}
