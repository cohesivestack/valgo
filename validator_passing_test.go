package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassingAsInvalid(t *testing.T) {
	ResetMessages()

	v := Is(String("USD").Passing(func(v string) bool {
		return v == "BTC"
	}))

	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 is not valid")
	}
}

func TestCustomValidatorAsValid(t *testing.T) {
	ResetMessages()

	v := Is(String("BTC").Passing(func(v string) bool {
		return v == "BTC"
	}))

	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
