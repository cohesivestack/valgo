package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLocalization(t *testing.T) {
	ResetMessages()

	SetDefaultLocale("es")
	v := IsString(" ").Not().Blank()
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent
	v = IsString(" ").Empty()
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 debe estar vacío")
}

func TestSeparatedLocalization(t *testing.T) {
	ResetMessages()

	err := SetDefaultLocale("en")
	assert.NoError(t, err)

	localized, err := Localized("es")
	assert.NoError(t, err)

	v := localized.CheckString(" ", "my_value").Not().Blank().Empty()
	assert.Contains(t, v.Errors()["my_value"].Messages(), "My value no puede estar en blanco")
	assert.Contains(t, v.Errors()["my_value"].Messages(), "My value debe estar vacío")

	// Default localization must not be changed
	v = IsString(" ").Not().Blank()
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank")
}
