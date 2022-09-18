package valgo

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotError(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := Is(String(value).Blank())
		assert.True(t, v.Valid())
		assert.NoError(t, v.Error())
	}
}

func TestError(t *testing.T) {
	v := Is(String("Vitalik Buterin").Blank())
	assert.False(t, v.Valid())
	assert.Error(t, v.Error())
}

func TestAddErrorMessageFromValidator(t *testing.T) {
	ResetMessages()

	v := Is(String("Vitalik Buterin", "name").Blank())

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

	v.Is(String("Vitalik Buterin", "name").Blank())

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["email"].Messages(), 1)
	assert.Len(t, v.Errors()["name"].Messages(), 1)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email is invalid")
	assert.Contains(t, v.Errors()["name"].Messages(), "Name must be blank")
}

func TestMultipleErrorsInOneFieldWithIs(t *testing.T) {
	r, _ := regexp.Compile("a")
	v := Is(String("", "email").Not().Blank().MatchingTo(r))

	assert.Len(t, v.Errors(), 1)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["email"].Messages(), 1)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email can't be blank")
}

func TestMultipleErrorsInOneFieldWithCheck(t *testing.T) {
	r, _ := regexp.Compile("a")
	v := Check(String("", "email").Not().Blank().MatchingTo(r))

	assert.Len(t, v.Errors(), 1)
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors()["email"].Messages(), 2)
	assert.Contains(t, v.Errors()["email"].Messages(), "Email can't be blank")
	assert.Contains(t, v.Errors()["email"].Messages(), "Email must match to \"a\"")
}

func TestErrorMarshallJSONWithIs(t *testing.T) {
	r, _ := regexp.Compile("a")
	v := Is(String("", "email").Not().Blank().MatchingTo(r)).
		Is(String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["name"])
	assert.Equal(t, "Email can't be blank", jsonMap["email"])

}

func TestErrorMarshallJSONWithCheck(t *testing.T) {
	r, _ := regexp.Compile("a")
	v := Check(String("", "email").Not().Blank().MatchingTo(r)).
		Check(String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["name"])
	emailErrors := jsonMap["email"].([]interface{})
	assert.Contains(t, emailErrors, "Email can't be blank")
	assert.Contains(t, emailErrors, "Email must match to \"a\"")
}

func TestIsValidByName(t *testing.T) {
	v := Is(String("Steve", "firstName").Not().Blank()).
		Is(String("", "lastName").Not().Blank())

	assert.True(t, v.IsValid("firstName"))
	assert.False(t, v.IsValid("lastName"))
}
