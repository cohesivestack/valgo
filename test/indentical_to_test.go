package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestIdenticalToValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		10,
		1,
		"a",
		"10",
		&MyStruct{FieldInt: 10},
		&[]int{10, 20},
		&map[string]int{"a": 1}} {
		msg := fmt.Sprintf("not assert using %+v", value)
		v := valgo.Is(value).IdenticalTo(value)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestIdenticalToInvalid(t *testing.T) {
	valgo.ResetMessages()

	funcToTest := func(value1 interface{}, value2 interface{}) {
		msg := fmt.Sprintf("not assert using %+v", value2)
		v := valgo.Is(value1).IdenticalTo(value2)
		assert.False(t, v.Valid(), msg)
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" must be identical to \"%v\"", value2),
				msg)
		}
	}

	for key, value := range map[interface{}]interface{}{
		1:          "1",
		"a":        "ab",
		&[]int{10}: &[]int{10},
		10.0:       10,
	} {
		funcToTest(key, value)
	}

	structA := MyStruct{
		FieldInt: 1,
	}
	structB := structA

	funcToTest(structA, structB)

	arrayA := []int{10}
	arrayB := arrayA
	funcToTest(arrayA, arrayB)

	mapA := map[string]int{"a": 1}
	mapB := mapA
	funcToTest(mapA, mapB)

}

func TestNotIdenticalToValid(t *testing.T) {
	valgo.ResetMessages()

	funcToTest := func(value1 interface{}, value2 interface{}) {
		msg := fmt.Sprintf("not assert using %+v", value2)
		v := valgo.Is(value1).NotIdenticalTo(value2)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}

	for key, value := range map[interface{}]interface{}{
		1:          "1",
		"a":        "ab",
		&[]int{10}: &[]int{10},
		10.0:       10,
	} {
		funcToTest(key, value)
	}

	structA := MyStruct{
		FieldInt: 1,
	}
	structB := structA

	funcToTest(structA, structB)

	arrayA := []int{10}
	arrayB := arrayA
	funcToTest(arrayA, arrayB)

	mapA := map[string]int{"a": 1}
	mapB := mapA
	funcToTest(mapA, mapB)

}

func TestNotIdenticalToInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		10,
		1,
		"a",
		"10",
		&MyStruct{FieldInt: 10},
		&[]int{10, 20},
		&map[string]int{"a": 1}} {
		msg := fmt.Sprintf("not assert using %+v", value)
		v := valgo.Is(value).NotIdenticalTo(value)
		assert.False(t, v.Valid(), msg)
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" can't be identical to \"%v\"", value),
				msg)
		}
	}

}
