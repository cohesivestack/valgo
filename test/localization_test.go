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

func TestCustomLocalization(t *testing.T) {
	valgo.ResetMessages()

	// Add new locale
	valgo.AddOrReplaceLocale("liberland", map[string]string{
		"not_blank": "Liberland say \"{{Title}}\" can't be blank",
		"empty":     "Liberland say \"{{Title}}\" must be empty",
	})

	localized, err := valgo.Localized("liberland")
	assert.NoError(t, err)

	v := localized.Is(" ").NotBlank().Empty()
	assert.Contains(t, v.Errors()[0].Messages, "Liberland say \"value0\" can't be blank")
	assert.Contains(t, v.Errors()[0].Messages, "Liberland say \"value0\" must be empty")

	// Replace existing locale
	valgo.AddOrReplaceLocale("en", map[string]string{
		"not_blank": "An improved english say \"{{Title}}\" can't be blank",
		"empty":     "An improved english say \"{{Title}}\" must be empty",
	})

	localized, err = valgo.Localized("en")
	assert.NoError(t, err)

	v = localized.Is(" ").NotBlank().Empty()
	assert.Contains(t, v.Errors()[0].Messages, "An improved english say \"value0\" can't be blank")
	assert.Contains(t, v.Errors()[0].Messages, "An improved english say \"value0\" must be empty")

}

func TestGetLocaleIsACopy(t *testing.T) {
	valgo.ResetMessages()

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	locale, err := valgo.GetLocaleCopy("en")
	assert.NoError(t, err)

	assert.True(t, len(locale) > 1)

	// Check that is a real copy
	locale["not_blank"] = "This message should not be assigned"

	v := valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank")
	assert.NotContains(t, v.Errors()[0].Messages, "This message should not be assigned")

	// Check copy can be used as base to replace a locale
	locale["not_blank"] = "This message is changed for test purposes"

	valgo.AddOrReplaceLocale("en", locale)

	v = valgo.Is(" ").NotBlank().Empty()
	assert.Contains(t, v.Errors()[0].Messages, "This message is changed for test purposes")
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")
}
