package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLocalization(t *testing.T) {
	valgo.SetDefaultLocale("es")
	v := valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" no puede estar en blanco")
}

func TestSeparatedLocalization(t *testing.T) {
	valgo.SetDefaultLocale("en")

	localized, err := valgo.Localized("es")
	assert.NoError(t, err)

	v := localized.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" no puede estar en blanco")

	v = valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank")
}
