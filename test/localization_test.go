package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLocalization(t *testing.T) {
	valgo.ResetMessages()

	valgo.SetDefaultLocale("es")
	v := valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" no puede estar en blanco")

	// Default localization must be persistent
	v = valgo.Is(" ").Empty()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" debe estar vacío")
}

func TestSeparatedLocalization(t *testing.T) {
	valgo.ResetMessages()

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	localized, err := valgo.Localized("es")
	assert.NoError(t, err)

	v := localized.Is(" ").NotBlank().Empty()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" no puede estar en blanco")
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" debe estar vacío")

	// Default localization must not be changed
	v = valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank")
}
