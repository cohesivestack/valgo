package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestAnIntegerValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{"1", "10", "-10", 10, -10} {
		v := valgo.Is(value).AnInteger()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestAnIntegerInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{"10.1", "-10.1", "", " ", "a", "a10", ".10", "@1", []int{10}, 10.1} {
		v := valgo.Is(value).AnInteger()

		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1)
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be an integer number")
		}
	}
}
