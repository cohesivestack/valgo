package valgo

import "reflect"

func identicalTo(valueA interface{}, valueB interface{}) bool {
	if !isComparableType(valueA) || !isComparableType(valueB) {
		return false
	}

	if reflect.TypeOf(valueA) != reflect.TypeOf(valueB) {
		return false
	}

	return valueA == valueB
}

func (validator *Validator) IdenticalTo(value interface{}, template ...string) *Validator {

	if !identicalTo(validator.currentValue, value) {
		validator.invalidate("identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotIdenticalTo(value interface{}, template ...string) *Validator {

	if identicalTo(validator.currentValue, value) {
		validator.invalidate("not_identical_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
