package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestIdenticalToValid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	_struct := MyStruct{FieldInt: 10}
	_map := map[string]int{"a": 10}
	_array := []int{10}
	for description, values := range map[string][]interface{}{
		"integers":        []interface{}{1, 1},
		"strings":         []interface{}{"a", "a"},
		"float":           []interface{}{10.0, 10.0},
		"pointer-integer": []interface{}{&_integer, &_integer},
		"pointer-float":   []interface{}{&_float, &_float},
		"pointer-string":  []interface{}{&_string, &_string},
		"pointer-array":   []interface{}{&_array, &_array},
		"pointer-map":     []interface{}{&_map, &_map},
		"struct":          []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
		"pointer-struct":  []interface{}{&_struct, &_struct},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).IdenticalTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestIdenticalToInvalid(t *testing.T) {
	valgo.ResetMessages()

	_integerA := 10
	_integerB := 10
	_floatA := 10.0
	_floatB := 10.0
	_stringA := "a"
	_stringB := "a"
	for description, values := range map[string][]interface{}{
		"integer":                         []interface{}{1, 2},
		"int32 int64":                     []interface{}{int32(1), int64(1)},
		"uint int":                        []interface{}{uint(1), int(1)},
		"string":                          []interface{}{"ab", "a"},
		"string integer":                  []interface{}{"1", 1},
		"float":                           []interface{}{1.1, 1.2},
		"float32 float64":                 []interface{}{float32(32), float64(64)},
		"float integer":                   []interface{}{10.0, 10},
		"pointer-integer integer":         []interface{}{&_integerA, 10},
		"pointer-integer pointer-integer": []interface{}{&_integerA, &_integerB},
		"pointer-float float":             []interface{}{&_floatA, 10.0},
		"pointer-float pointer-float":     []interface{}{&_floatA, &_floatB},
		"pointer-string string":           []interface{}{&_stringA, "ab"},
		"pointer-string pointer-string":   []interface{}{&_stringA, &_stringB},
		"array":          []interface{}{[]int{10}, []int{11}},
		"pointer-array":  []interface{}{&[]int{10}, &[]int{10}},
		"map":            []interface{}{map[string]int{"a": 10}, map[string]int{"a": 10}},
		"pointer-map":    []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 10}},
		"struct":         []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
		"pointer-struct": []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 10}},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).IdenticalTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid(), msg)
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" must be equal to \"%v\"", valueB), msg)
		}
	}
}

func TestNotIdenticalToValid(t *testing.T) {
	valgo.ResetMessages()

	_integerA := 10
	_integerB := 10
	_floatA := 10.0
	_floatB := 10.0
	_stringA := "a"
	_stringB := "a"
	for description, values := range map[string][]interface{}{
		"integer":                         []interface{}{1, 2},
		"int32 int64":                     []interface{}{int32(1), int64(1)},
		"uint int":                        []interface{}{uint(1), int(1)},
		"string":                          []interface{}{"ab", "a"},
		"string integer":                  []interface{}{"1", 1},
		"float":                           []interface{}{1.1, 1.2},
		"float32 float64":                 []interface{}{float32(32), float64(64)},
		"float integer":                   []interface{}{10.0, 10},
		"pointer-integer integer":         []interface{}{&_integerA, 10},
		"pointer-integer pointer-integer": []interface{}{&_integerA, &_integerB},
		"pointer-float float":             []interface{}{&_floatA, 10.0},
		"pointer-float pointer-float":     []interface{}{&_floatA, &_floatB},
		"pointer-string string":           []interface{}{&_stringA, "ab"},
		"pointer-string pointer-string":   []interface{}{&_stringA, &_stringB},
		"array":          []interface{}{[]int{10}, []int{11}},
		"pointer-array":  []interface{}{&[]int{10}, &[]int{10}},
		"map":            []interface{}{map[string]int{"a": 10}, map[string]int{"a": 10}},
		"pointer-map":    []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 10}},
		"struct":         []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
		"pointer-struct": []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 10}},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).Not().IdenticalTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestNotIdenticalToInvalid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	_struct := MyStruct{FieldInt: 10}
	_map := map[string]int{"a": 10}
	_array := []int{10}
	for description, values := range map[string][]interface{}{
		"integers":        []interface{}{1, 1},
		"strings":         []interface{}{"a", "a"},
		"float":           []interface{}{10.0, 10.0},
		"pointer-integer": []interface{}{&_integer, &_integer},
		"pointer-float":   []interface{}{&_float, &_float},
		"pointer-string":  []interface{}{&_string, &_string},
		"pointer-array":   []interface{}{&_array, &_array},
		"pointer-map":     []interface{}{&_map, &_map},
		"struct":          []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
		"pointer-struct":  []interface{}{&_struct, &_struct},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).Not().IdenticalTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid(), msg)
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" can't be equal to \"%v\"", valueB), msg)
		}
	}
}
