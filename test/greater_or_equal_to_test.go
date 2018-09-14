package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestGreaterOrEqualToValid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 2
	_float := 1.1
	_string := "1.1"
	for description, values := range map[string][]interface{}{
		"integers-greater-than":                []interface{}{2, 1},
		"integers-equal-to":                    []interface{}{1, 1},
		"strings-letters-greater-than":         []interface{}{"b", "a"},
		"strings-letters-equal-to":             []interface{}{"a", "a"},
		"strings-numbers-greater-than":         []interface{}{"2", "11"},
		"strings-numbers-equal-to":             []interface{}{"2", "1"},
		"string-greater-than integer":          []interface{}{"1.1", 1},
		"string-equal-to integer":              []interface{}{"1", 1},
		"string-greater-than float":            []interface{}{"1.1", 1.0},
		"string-equal-to float":                []interface{}{"1.1", 1.1},
		"float-greater-than integer":           []interface{}{1.1, 1},
		"float-equal-to integer":               []interface{}{1.0, 1},
		"pointer-integer-greater-than integer": []interface{}{&_integer, 1},
		"pointer-integer-equal-to integer":     []interface{}{&_integer, 2},
		"pointer-integer-greater-than string":  []interface{}{&_integer, "1"},
		"pointer-integer-equal-to string":      []interface{}{&_integer, "2"},
		"pointer-float-greater-than float":     []interface{}{&_float, 1.0},
		"pointer-float-equal-to float":         []interface{}{&_float, 1.1},
		"pointer-float-greater-than integer":   []interface{}{&_float, 1},
		"pointer-float-equal-to integer":       []interface{}{&_float, 1.1},
		"pointer-string-greater-than string":   []interface{}{&_string, "1.0"},
		"pointer-string-equal-to string":       []interface{}{&_string, "1.1"},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).GreaterOrEqualTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.ErrorItems(), msg)
	}
}

func TestGreaterOrEqualToInvalid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 1
	_float := 1.0
	_string := "a"
	for description, values := range map[string][]interface{}{
		"integers-less-than":                []interface{}{1, 2},
		"strings-letters-less-than":         []interface{}{"a", "b"},
		"strings-numbers-less-than":         []interface{}{"11", "2"},
		"string-less-than integer":          []interface{}{"1", 2},
		"string-less-than float":            []interface{}{"1", 2.0},
		"float-less-than integer":           []interface{}{1.0, 2},
		"pointer-integer-less-than integer": []interface{}{&_integer, 2},
		"pointer-integer-less-than string":  []interface{}{&_integer, "2"},
		"pointer-float-less-than float":     []interface{}{&_float, 1.1},
		"pointer-float-less-than integer":   []interface{}{&_float, 2},
		"pointer-string-less-than string":   []interface{}{&_string, "ab"},
		// Different types than strings and number never should be true
		"array":                 []interface{}{[]int{2}, []int{1}},
		"pointer-array":         []interface{}{&[]int{2}, &[]int{1}},
		"pointer-array array":   []interface{}{&[]int{2}, []int{1}},
		"map":                   []interface{}{map[string]int{"a": 2}, map[string]int{"a": 1}},
		"pointer-map":           []interface{}{&map[string]int{"a": 2}, &map[string]int{"a": 1}},
		"pointer-map map":       []interface{}{&map[string]int{"a": 2}, map[string]int{"a": 1}},
		"struct":                []interface{}{MyStruct{FieldInt: 2}, MyStruct{FieldInt: 1}},
		"pointer-struct":        []interface{}{&MyStruct{FieldInt: 2}, &MyStruct{FieldInt: 1}},
		"pointer-struct struct": []interface{}{&MyStruct{FieldInt: 2}, MyStruct{FieldInt: 1}},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).GreaterOrEqualTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid(), msg)
		if assert.NotEmpty(t, v.ErrorItems(), msg) {
			assert.Len(t, v.ErrorItems(), 1, msg)
			assert.Contains(t, v.ErrorItems()[0].Messages,
				fmt.Sprintf("\"value0\" must be greater or equal to \"%v\"", valueB), msg)
		}
	}
}
