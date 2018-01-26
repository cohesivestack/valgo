package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidatorAsInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("USD").Passing(func(_v *valgo.CustomValidator, _t ...string) {
		if _v.ValueAsString() != "BTC" {
			_v.Invalidate("equal_to", _t, map[string]interface{}{"Value": "BTC"})
		}
	})

	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be equal to \"BTC\"")
	}
}

func TestCustomValidatorAsValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("BTC").Passing(func(_v *valgo.CustomValidator, _t ...string) {
		if _v.ValueAsString() != "BTC" {
			_v.Invalidate("equal_to", _t, nil)
		}
	})

	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}