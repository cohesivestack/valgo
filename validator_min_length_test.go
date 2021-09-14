package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMinLengthValid(t *testing.T) {
	for _, value := range []string{"123", "1234"} {
		v := IsString(value).MinLength(3)
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.True(t, v.Valid(), m)
		assert.Empty(t, v.Errors(), m)
	}
}

func TestStringMinLengthInvalid(t *testing.T) {
	v := IsString("12").MinLength(3)
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		`Value 0 must not have a length shorter than "3"`)
}

func TestStringNotMinLengthValid(t *testing.T) {
	v := IsString("12").Not().MinLength(3)
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringNotMinLengthInvalid(t *testing.T) {
	for _, value := range []string{"123", "1234"} {
		v := IsString(value).Not().MinLength(3)
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.False(t, v.Valid(), m)
		if assert.NotEmpty(t, v.Errors(), m) {
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				`Value 0 must not have a length longer than or equal to "3"`, m)
		}
	}
}
