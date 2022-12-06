package valgo_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotError(t *testing.T) {
	t.Parallel()

	for _, value := range []string{"", " "} {
		v := valgo.Is(valgo.String(value).Blank())
		assert.True(t, v.Valid())
		assert.NoError(t, v.Error())
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	v := valgo.Is(valgo.String("Vitalik Buterin").Blank())
	assert.False(t, v.Valid())
	assert.Error(t, v.Error())
}

func TestAddErrorMessageFromValidator(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.Is(valgo.String("Vitalik Buterin", "name").Blank())

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Len(t, v.ErrorByKey("name").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("name").Messages(), "Name must be blank")

	v.AddErrorMessage("email", "Email is invalid")

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorByKey("name").Messages(), 1)
	assert.Len(t, v.ErrorByKey("email").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("name").Messages(), "Name must be blank")
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email is invalid")
}

func TestAddErrorMessageFromValgo(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.AddErrorMessage("email", "Email is invalid")

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Len(t, v.ErrorByKey("email").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email is invalid")

	v.Is(valgo.String("Vitalik Buterin", "name").Blank())

	assert.Len(t, v.Errors(), 2)
	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorByKey("email").Messages(), 1)
	assert.Len(t, v.ErrorByKey("name").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email is invalid")
	assert.Contains(t, v.ErrorByKey("name").Messages(), "Name must be blank")
}

func TestMultipleErrorsInOneFieldWithIs(t *testing.T) {
	t.Parallel()

	r := regexp.MustCompile("a")
	v := valgo.Is(valgo.String("", "email").Not().Blank().MatchingTo(r))

	assert.Len(t, v.Errors(), 1)
	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorByKey("email").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email can't be blank")
}

func TestMultipleErrorsInOneFieldWithCheck(t *testing.T) {
	t.Parallel()

	r := regexp.MustCompile("a")
	v := valgo.Check(valgo.String("", "email").Not().Blank().MatchingTo(r))

	assert.Len(t, v.Errors(), 1)
	assert.False(t, v.Valid())
	assert.Len(t, v.ErrorByKey("email").Messages(), 2)
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email can't be blank")
	assert.Contains(t, v.ErrorByKey("email").Messages(), "Email must match to \"a\"")
}

func TestErrorMarshallJSONWithIs(t *testing.T) {
	t.Parallel()

	r := regexp.MustCompile("a")
	v := valgo.Is(valgo.String("", "email").Not().Blank().MatchingTo(r)).
		Is(valgo.String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string][]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["name"][0])
	assert.Equal(t, "Email can't be blank", jsonMap["email"][0])
}

func TestErrorMarshallJSONWithCheck(t *testing.T) {
	t.Parallel()

	r := regexp.MustCompile("a")
	v := valgo.Check(valgo.String("", "email").Not().Blank().MatchingTo(r)).
		Check(valgo.String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string][]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["name"][0])
	assert.Contains(t, jsonMap["email"], "Email can't be blank")
	assert.Contains(t, jsonMap["email"], "Email must match to \"a\"")
}

func TestIsValidByName(t *testing.T) {
	t.Parallel()

	v := valgo.Is(valgo.String("Steve", "firstName").Not().Blank()).
		Is(valgo.String("", "lastName").Not().Blank())

	assert.True(t, v.IsValid("firstName"))
	assert.False(t, v.IsValid("lastName"))
}

func TestCustomErrorMarshallJSON(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	customFunc := func(e *valgo.Error) ([]byte, error) {
		errors := map[string]interface{}{}

		for k, v := range e.Errors() {
			if len(v.Messages()) == 1 {
				errors[k] = v.Messages()[0]
			} else {
				errors[k] = v.Messages()
			}
		}

		// Add root level errors to customize errors interface
		return json.Marshal(map[string]map[string]interface{}{"errors": errors})
	}

	valgo.SetMarshalJSON(customFunc)

	r := regexp.MustCompile("a")

	v := valgo.Check(valgo.String("", "email").Not().Blank().MatchingTo(r)).
		Check(valgo.String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	t.Log(string(jsonByte))

	jsonMap := map[string]map[string]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["errors"]["name"])

	emailErrors, ok := jsonMap["errors"]["email"].([]interface{})
	require.True(t, ok)

	assert.Contains(t, emailErrors, "Email can't be blank")
	assert.Contains(t, emailErrors, "Email must match to \"a\"")
}
