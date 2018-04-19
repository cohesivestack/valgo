package valgo

import (
	"fmt"
	"reflect"
	"strconv"
)

type Value struct {
	original interface{}
	absolute interface{}

	isString      *bool
	isNumber      *bool
	isNumberType  *bool
	isInteger     *bool
	isIntegerType *bool
	isBlank       *bool
	isEmpty       *bool

	isComparableType *bool

	asString  *string
	asFloat64 *float64
}

func boolPointer(value bool) *bool {
	return &value
}

func float64Pointer(value float64) *float64 {
	return &value
}

func stringAsPointer(value string) *string {
	return &value
}

func NewValue(value interface{}) *Value {
	_value := &Value{
		original: value,
	}

	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		_value.absolute = reflect.Indirect(rv).Interface()
	} else {
		_value.absolute = value
	}

	return _value
}

func isComparableType(value interface{}) bool {
	_value := reflect.ValueOf(value)
	switch _value.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return false
	}
	return true
}

func (value *Value) IsComparableType() bool {
	if value.isComparableType == nil {
		value.isComparableType = boolPointer(isComparableType(value.absolute))
	}
	return *value.isComparableType
}

func (value *Value) AsFloat64() float64 {
	if value.asFloat64 == nil {
		_value := value.absolute

		switch _value.(type) {
		case uint:
			value.asFloat64 = float64Pointer(float64(_value.(uint)))
		case uint8:
			value.asFloat64 = float64Pointer(float64(_value.(uint8)))
		case uint16:
			value.asFloat64 = float64Pointer(float64(_value.(uint16)))
		case uint32:
			value.asFloat64 = float64Pointer(float64(_value.(uint32)))
		case uint64:
			value.asFloat64 = float64Pointer(float64(_value.(uint64)))
		case int:
			value.asFloat64 = float64Pointer(float64(_value.(int)))
		case int8:
			value.asFloat64 = float64Pointer(float64(_value.(int8)))
		case int16:
			value.asFloat64 = float64Pointer(float64(_value.(int16)))
		case int32:
			value.asFloat64 = float64Pointer(float64(_value.(int32)))
		case int64:
			value.asFloat64 = float64Pointer(float64(_value.(int64)))
		case float32:
			value.asFloat64 = float64Pointer(float64(_value.(float32)))
		case float64:
			value.asFloat64 = float64Pointer(_value.(float64))
		case string:
			_value, err := strconv.ParseFloat(_value.(string), 64)
			if err != nil {
				_value = 0
			}
			value.asFloat64 = float64Pointer(_value)
		default:
			value.asFloat64 = float64Pointer(0)
		}
	}

	return *value.asFloat64
}

func convertToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func (value *Value) AsString() string {
	if value.asString == nil {
		value.asString = stringAsPointer(convertToString(value.absolute))
	}
	return *value.asString
}
