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
	_val := &Value{
		original: value,
	}

	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		_val.absolute = reflect.Indirect(rv).Interface()
	} else {
		_val.absolute = value
	}

	return _val
}

func isComparableType(value interface{}) bool {
	_val := reflect.ValueOf(value)
	switch _val.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return false
	}
	return true
}

func (val *Value) IsComparableType() bool {
	if val.isComparableType == nil {
		val.isComparableType = boolPointer(isComparableType(val.absolute))
	}
	return *val.isComparableType
}

func (val *Value) AsFloat64() float64 {
	if val.asFloat64 == nil {
		_val := val.absolute

		switch _val.(type) {
		case uint:
			val.asFloat64 = float64Pointer(float64(_val.(uint)))
		case uint8:
			val.asFloat64 = float64Pointer(float64(_val.(uint8)))
		case uint16:
			val.asFloat64 = float64Pointer(float64(_val.(uint16)))
		case uint32:
			val.asFloat64 = float64Pointer(float64(_val.(uint32)))
		case uint64:
			val.asFloat64 = float64Pointer(float64(_val.(uint64)))
		case int:
			val.asFloat64 = float64Pointer(float64(_val.(int)))
		case int8:
			val.asFloat64 = float64Pointer(float64(_val.(int8)))
		case int16:
			val.asFloat64 = float64Pointer(float64(_val.(int16)))
		case int32:
			val.asFloat64 = float64Pointer(float64(_val.(int32)))
		case int64:
			val.asFloat64 = float64Pointer(float64(_val.(int64)))
		case float32:
			val.asFloat64 = float64Pointer(float64(_val.(float32)))
		case float64:
			val.asFloat64 = float64Pointer(_val.(float64))
		case string:
			_val, err := strconv.ParseFloat(_val.(string), 64)
			if err != nil {
				_val = 0
			}
			val.asFloat64 = float64Pointer(_val)
		default:
			val.asFloat64 = float64Pointer(0)
		}
	}

	return *val.asFloat64
}

func convertToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func (val *Value) AsString() string {
	if val.asString == nil {
		val.asString = stringAsPointer(convertToString(val.absolute))
	}
	return *val.asString
}
