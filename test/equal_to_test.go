package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEqualTo(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(10).EqualTo(11)
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be equal to \"11\"")
	}

	v = valgo.Is(10).EqualTo(10)
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestNotEqualTo(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(10).NotEqualTo(10)
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be equal to \"10\"")
	}

	v = valgo.Is(10).NotEqualTo(11)
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
