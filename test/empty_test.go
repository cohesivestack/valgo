package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	for _, value := range []string{"Vitalik Buterin", " "} {
		v = valgo.Is(value).Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty", fmt.Sprintf("not assert using %s", value))
		}
	}
}

func TestNotEmpty(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").NotEmpty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}

	for _, value := range []string{"Vitalik Buterin", " "} {
		v = valgo.Is(value).NotEmpty()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}
