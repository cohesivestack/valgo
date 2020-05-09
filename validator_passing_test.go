package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomValidatorAsInvalid(t *testing.T) {
	ResetMessages()

	v := IsString("USD").Passing(func(_v *CustomValidator, _t ...string) {
		if _v.Value().(string) != "BTC" {
			_v.Invalidate("equal_to", map[string]interface{}{"value": "BTC"}, _t...)
		}
	})

	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 must be equal to \"BTC\"")
	}
}

func TestCustomValidatorAsValid(t *testing.T) {
	ResetMessages()

	v := IsString("BTC").Passing(func(_v *CustomValidator, _t ...string) {
		if _v.Value().(string) != "BTC" {
			_v.Invalidate("equal_to", nil, _t...)
		}
	})

	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
