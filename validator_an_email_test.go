package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Invalid and valid email addresses,
// credit to @cjaoude: https://gist.github.com/cjaoude/fd9910626629b53c4d25

func TestAnEmailValid(t *testing.T) {
	ResetMessages()

	for _, value := range []string{
		"email@example.com",
		"firstname.lastname@example.com",
		"email@subdomain.example.com",
		"firstname+lastname@example.com",
		"email@123.123.123.123",
		"1234567890@example.com",
		"email@example-one.com",
		"_______@example.com",
		"email@example.name",
		"email@example.museum",
		"email@example.co.jp",
		"firstname-lastname@example.com",
	} {
		v := IsString(value).AnEmail()

		msg := fmt.Sprintf("not assert using %s", value)

		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestAnEmailInvalid(t *testing.T) {
	ResetMessages()

	for _, value := range []string{
		"plainaddress",
		"#@%^%#$@#$@#.com",
		"@example.com",
		"Joe Smith <email@example.com>",
		"email@example@example.com",
		"あいうえお@example.com",
		"email@example.com (Joe Smith)",
		"email@-example.com",
		"email@example..com"} {
		v := IsString(value).AnEmail()

		msg := fmt.Sprintf("not assert using %s", value)

		assert.False(t, v.Valid(), msg)
		assert.Len(t, v.Errors(), 1, msg)
		assert.Contains(t,
			v.Errors()["value_0"].Messages(),
			"Value 0 is not an email address", msg)
	}
}
