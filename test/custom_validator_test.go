package test

import (
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidatorAsInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("USD").Passing(func(_v *valgo.CustomValidator, _t ...string) {
		if _v.Value().AsString() != "BTC" {
			_v.Invalidate("equal_to", _t, map[string]interface{}{"Value": "BTC"})
		}
	})

	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.ErrorItems()) {
		assert.Len(t, v.ErrorItems(), 1)
		assert.Contains(t, v.ErrorItems()[0].Messages, "\"value0\" must be equal to \"BTC\"")
	}
}

func TestCustomValidatorAsValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("BTC").Passing(func(_v *valgo.CustomValidator, _t ...string) {
		if _v.Value().AsString() != "BTC" {
			_v.Invalidate("equal_to", _t, nil)
		}
	})

	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorItems())
}
