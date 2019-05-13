package test

import (
	"fmt"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEqualToValid(t *testing.T) {
	valgo.ResetMessages()

	_pointer := 10
	for description, values := range map[string][]interface{}{
		"integers": []interface{}{1, 1},
		"strings":  []interface{}{"a", "a"},
		"floats":   []interface{}{10.0, 10.0},
		"pointers": []interface{}{&_pointer, &_pointer},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).EqualTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestEqualToInvalid(t *testing.T) {
	valgo.ResetMessages()

	// _pointerA := 10
	// _pointerB := 10
	for description, values := range map[string][]interface{}{
		"integers":      []interface{}{1, 2},
		"strings":       []interface{}{"a", "b"},
		"floats":        []interface{}{10.0, 11.0},
		"integer_float": []interface{}{10.0, 10},
		// "integer_pointer": []interface{}{10, &_pointerA},
		// "pointers":        []interface{}{&_pointerA, &_pointerB},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).EqualTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid(), msg)
		assert.Len(t, v.Errors(), 1, msg)
		assert.Contains(t,
			v.Errors()["value_0"].Messages(),
			fmt.Sprintf("Value 0 must be equal to \"%v\"", valueB), msg)
	}
}
