package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLocalization(t *testing.T) {
	teardown()

	SetDefaultLocale("es")
	v := Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent
	v = Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 debe estar vacío")
}

func TestSeparatedLocalization(t *testing.T) {
	teardown()

	err := SetDefaultLocale("en")
	assert.NoError(t, err)

	localized, err := Localized("es")
	assert.NoError(t, err)

	v := localized.New().Check(String(" ", "my_value").Not().Blank().Empty())
	assert.Contains(t, v.Errors()["my_value"].Messages(), "My value no puede estar en blanco")
	assert.Contains(t, v.Errors()["my_value"].Messages(), "My value debe estar vacío")

	// Default localization must not be changed
	v = Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank")
}
