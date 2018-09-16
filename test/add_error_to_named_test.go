package test

import (
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestAddErrorToNamedFromValidator(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("Vitalik Buterin").Named("name").Blank()

	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorItems(), 1)
	assert.Contains(t, v.ErrorItems()[0].Messages, "Name must be blank")

	v.AddErrorToNamed("email", "Email is invalid")

	assert.Len(t, v.ErrorItems(), 2)
	assert.False(t, v.Valid())
	assert.Contains(t, v.ErrorItems()[0].Messages, "Name must be blank")
	assert.Contains(t, v.ErrorItems()[1].Messages, "Email is invalid")
}

func TestAddErrorToNamedFromValgo(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.AddErrorToNamed("email", "Email is invalid")

	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorItems(), 1)
	assert.Contains(t, v.ErrorItems()[0].Messages, "Email is invalid")

	v.Is("Vitalik Buterin").Named("name").Blank()

	assert.Len(t, v.ErrorItems(), 2)
	assert.False(t, v.Valid())
	assert.Contains(t, v.ErrorItems()[0].Messages, "Email is invalid")
	assert.Contains(t, v.ErrorItems()[1].Messages, "Name must be blank")
}
