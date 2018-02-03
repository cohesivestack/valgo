package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestAnIntegerTypeValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		10,
		-10,
		uint(10),
		uint8(10),
		uint16(10),
		uint32(10),
		uint64(10),
		int(10),
		int8(10),
		int16(10),
		int32(10),
		int64(10)} {
		v := valgo.Is(value).AnIntegerType()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestAnIntegerTypeInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		"1",
		"1.1",
		true,
		10.1,
		float32(10.1),
		float64(10.1),
		&[]int{10},
		[]int{10}} {
		v := valgo.Is(value).AnIntegerType()

		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1)
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be an integer type")
		}
	}
}
