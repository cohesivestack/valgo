package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestBlankValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []string{"", " "} {
		v := valgo.Is(value).Blank()

		assert.True(t, v.Valid())
		assert.Empty(t, v.ErrorItems(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestBlankInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.ErrorItems()) {
		assert.Len(t, v.ErrorItems(), 1)
		assert.Contains(t, v.ErrorItems()[0].Messages, "\"value0\" must be blank")
	}
}

func TestNotBlankValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("Vitalik Buterin").Not().Blank()
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorItems())
}

func TestNotBlankInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []string{" ", ""} {
		v := valgo.Is(value).Not().Blank()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.ErrorItems()) {
			assert.Len(t, v.ErrorItems(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t, v.ErrorItems()[0].Messages, "\"value0\" can't be blank", fmt.Sprintf("not assert using %s", value))
		}
	}
}
