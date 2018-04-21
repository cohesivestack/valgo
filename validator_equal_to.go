package valgo

import (
	"reflect"
)

func (valueA *Value) IsEqualTo(value interface{}) bool {
	valueB := NewValue(value)
	if valueA.IsComparableType() && valueB.IsComparableType() && valueA.absolute == valueB.absolute {
		return true
	}

	// if previous test was not true and one value is nil then just return false
	if valueA.absolute == nil || valueB.absolute == nil {
		return false
	}

	if valueA.IsNumber() && valueB.IsNumber() {
		return valueA.AsFloat64() == valueB.AsFloat64()
	}

	return reflect.DeepEqual(valueA.absolute, valueB.absolute)
}

func (validator *Validator) EqualTo(value interface{}, template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsEqualTo(value)) {
		validator.invalidate("equivalent_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}

	validator.resetNegative()

	return validator
}
