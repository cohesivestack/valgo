package valgo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotError(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := IsString(value).Blank()
		assert.True(t, v.Valid())
		assert.NoError(t, v.Error())
	}
}

func TestError(t *testing.T) {
	v := IsString("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	assert.Error(t, v.Error())
}

func TestAddErrorMessageFromValidator(t *testing.T) {
	ResetMessages()

	v := IsString("Vitalik Buterin", "name").Blank()

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Len(t, v.Errors()["name"].Messages(), 1)
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")

	v.AddErrorMessage("email", "Email is invalid")

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["name"].Messages(), 1)
	assert.Len(t, v.Errors()["email"].Messages(), 1)
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")
}

func TestAddErrorMessageFromValgo(t *testing.T) {
	ResetMessages()

	v := AddErrorMessage("email", "Email is invalid")

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Len(t, v.Errors()["email"].Messages(), 1)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")

	v.IsString("Vitalik Buterin", "name").Blank()

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["email"].Messages(), 1)
	assert.Len(t, v.Errors()["name"].Messages(), 1)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")
}

func TestMultipleErrorsInOneField(t *testing.T) {
	v := IsString("", "email").Not().Blank().AnEmail()

	assert.Len(t, v.Errors(), 1)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["email"].Messages(), 2)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email can't be blank")
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is not an email address")
}

func TestErrorMarshallJSON(t *testing.T) {
	v := IsString("", "email").Not().Blank().AnEmail().
		IsString("", "name").Not().Blank()

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, jsonMap["name"], "Name can't be blank")
	emailErrors := jsonMap["email"].([]interface{})
	assert.Contains(t, emailErrors, "Email can't be blank")
	assert.Contains(t, emailErrors, "Email is not an email address")
}
