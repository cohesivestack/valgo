package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestANumberValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		"1",
		"10",
		"10.1",
		"10.10",
		10,
		-10,
		10.1,
		-10.1,
		uint(10),
		uint8(10),
		uint16(10),
		uint32(10),
		uint64(10),
		int(10),
		int8(10),
		int16(10),
		int32(10),
		int64(10),
		float32(10.1),
		float64(10.1)} {
		v := valgo.Is(value).ANumber()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestANumberInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{"", " ", "a", "a10", ".10.1", "@1", []int{10}} {
		v := valgo.Is(value).ANumber()

		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1)
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be a number")
		}
	}
}
