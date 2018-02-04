package valgo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func isComparableType(value interface{}) bool {
	_value := reflect.ValueOf(value)
	switch _value.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return false
	}
	return true
}

func getNumberAsFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case uint:
		return float64(value.(uint)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case int:
		return float64(value.(int)), nil
	case int8:
		return float64(value.(int8)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	}
	return 0, errors.New(fmt.Sprintf("'%v' is not a number type", value))
}

func equivalentTo(valueA interface{}, valueB interface{}) bool {
	if isComparableType(valueA) && isComparableType(valueB) && valueA == valueB {
		return true
	}

	// if pass test was not true and one value is nil then just return false
	if valueA == nil || valueB == nil {
		return false
	}

	rvA := reflect.ValueOf(valueA)
	rvB := reflect.ValueOf(valueB)

	if rvA.Kind() == reflect.Ptr {
		valueA = reflect.Indirect(rvA).Interface()
	}

	if rvB.Kind() == reflect.Ptr {
		valueB = reflect.Indirect(rvB).Interface()
	}

	if aString(valueA) && aNumberType(valueB) {
		_valueA, err := strconv.ParseFloat(valueA.(string), 64)
		if err != nil {
			return false
		}
		_valueB, err := getNumberAsFloat64(valueB)
		if err != nil {
			return false
		}
		return _valueA == _valueB
	}

	if aString(valueB) && aNumberType(valueA) {
		_valueB, err := strconv.ParseFloat(valueB.(string), 64)
		if err != nil {
			return false
		}
		_valueA, err := getNumberAsFloat64(valueA)
		if err != nil {
			return false
		}
		return _valueB == _valueA
	}

	if aNumberType(valueA) && aNumberType(valueB) {
		_valueA, err := getNumberAsFloat64(valueA)
		if err != nil {
			return false
		}
		_valueB, err := getNumberAsFloat64(valueB)
		if err != nil {
			return false
		}
		return float64(_valueA) == float64(_valueB)
	}

	return reflect.DeepEqual(valueA, valueB)
}

func (validator *Validator) EquivalentTo(value interface{}, template ...string) *Validator {
	if !equivalentTo(validator.currentValue, value) {
		validator.invalidate("equivalent_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotEquivalentTo(value interface{}, template ...string) *Validator {
	if equivalentTo(validator.currentValue, value) {
		validator.invalidate("not_equivalent_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
