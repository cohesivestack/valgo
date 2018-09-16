package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

// Invalid and valid email addresses,
// credit to @cjaoude: https://gist.github.com/cjaoude/fd9910626629b53c4d25

func TestAnEmailValid(t *testing.T) {
	valgo.ResetMessages()

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
		v := valgo.Is(value).AnEmail()

		assert.True(t, v.Valid())
		assert.Empty(t, v.ErrorItems(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestAnEmailInvalid(t *testing.T) {
	valgo.ResetMessages()

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
		v := valgo.Is(value).AnEmail()

		if assert.NotEmpty(t, v.ErrorItems(), fmt.Sprintf("not assert using %s", value)) {
			assert.Len(t, v.ErrorItems(), 1)
			assert.Contains(t, v.ErrorItems()[0].Messages, "Value 0 is not an email address")
		}
	}
}
