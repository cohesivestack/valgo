package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestBlank(t *testing.T) {
	v := valgo.Is("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be blank")
	}

	for _, value := range []string{"", " "} {
		v = valgo.Is(value).Blank()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestNotBlank(t *testing.T) {
	v := valgo.Is("Vitalik Buterin").NotBlank()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	for _, value := range []string{" ", ""} {
		v = valgo.Is(value).NotBlank()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank", fmt.Sprintf("not assert using %s", value))
		}
	}
}
