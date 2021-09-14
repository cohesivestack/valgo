package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLengthValid(t *testing.T) {
	v := IsString("123").Length(3)
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringLengthInvalid(t *testing.T) {
	for _, value := range []string{"12", "1234"} {
		v := IsString(value).Length(3)
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.False(t, v.Valid(), m)
		if assert.NotEmpty(t, v.Errors(), m) {
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				`Value 0 must have a length equal to "3"`, m)
		}
	}
}

func TestStringNotLengthValid(t *testing.T) {
	for _, value := range []string{"12", "1234"} {
		v := IsString(value).Not().Length(3)
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.True(t, v.Valid(), m)
		assert.Empty(t, v.Errors(), m)
	}
}

func TestStringNotLengthInvalid(t *testing.T) {
	v := IsString("123").Not().Length(3)
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		`Value 0 must not have a length equal to "3"`)
}
