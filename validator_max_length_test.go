package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMaxLengthValid(t *testing.T) {
	for _, value := range []string{"123", "12"} {
		v := Is(String(value).MaxLength(3))
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.True(t, v.Valid(), m)
		assert.Empty(t, v.Errors(), m)
	}
}

func TestStringMaxLengthInvalid(t *testing.T) {
	v := Is(String("1234").MaxLength(3))

	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		`Value 0 must not have a length longer than "3"`)
}

func TestStringNotMaxLengthValid(t *testing.T) {
	v := Is(String("1234").Not().MaxLength(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestStringNotMaxLengthInvalid(t *testing.T) {
	for _, value := range []string{"123", "12"} {
		v := Is(String(value).Not().MaxLength(3))
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.False(t, v.Valid(), m)
		if assert.NotEmpty(t, v.Errors(), m) {
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				`Value 0 must not have a length shorter than or equal to "3"`, m)
		}
	}
}
