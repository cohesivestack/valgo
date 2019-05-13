package test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestNotError(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := valgo.IsString(value).Blank()
		assert.True(t, v.Valid())
		assert.NoError(t, v.Error())
	}
}

func TestError(t *testing.T) {
	v := valgo.IsString("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	assert.Error(t, v.Error())
}

func TestAddErrorMessageFromValidator(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.IsString("Vitalik Buterin", "name").Blank()

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")

	v.AddErrorMessage("email", "Email is invalid")

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")
}

func TestAddErrorMessageFromValgo(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.AddErrorMessage("email", "Email is invalid")

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")

	v.IsString("Vitalik Buterin", "name").Blank()

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")
}
