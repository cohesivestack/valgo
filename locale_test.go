package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseOtherLocale(t *testing.T) {

	v := New(Options{LocaleCode: LocaleCodeEs}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent in the same validation
	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 debe estar vacío")
}

func TestChangeLocaleEntries(t *testing.T) {

	originalErrorMessage0 := (*getLocaleEn())[ErrorKeyNotBlank]
	modifiedErrorMessage0 := "{{title}} should not be blank"

	originalErrorMessage1 := (*getLocaleEn())[ErrorKeyBlank]
	modifiedErrorMessage1 := "{{title}} should be blank"

	assert.NotEqual(t, originalErrorMessage0, modifiedErrorMessage0)
	assert.NotEqual(t, originalErrorMessage1, modifiedErrorMessage1)

	locale := &Locale{
		ErrorKeyNotBlank: modifiedErrorMessage0,
		ErrorKeyBlank:    modifiedErrorMessage1,
	}

	v := New(Options{Locale: locale}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 should not be blank")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 should be blank")

	// Other entries should not be modified
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 can't be empty")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 must be empty")
}

func TestUseOtherLocaleAndChangeLocaleEntries(t *testing.T) {

	originalErrorMessage0 := (*getLocaleEs())[ErrorKeyNotBlank]
	modifiedErrorMessage0 := "{{title}} no debería estar en blanco"

	originalErrorMessage1 := (*getLocaleEs())[ErrorKeyBlank]
	modifiedErrorMessage1 := "{{title}} debería estar en blanco"

	assert.NotEqual(t, originalErrorMessage0, modifiedErrorMessage0)
	assert.NotEqual(t, originalErrorMessage1, modifiedErrorMessage1)

	locale := &Locale{
		ErrorKeyNotBlank: modifiedErrorMessage0,
		ErrorKeyBlank:    modifiedErrorMessage1,
	}

	v := New(Options{LocaleCode: LocaleCodeEs, Locale: locale}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no debería estar en blanco")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 debería estar en blanco")

	// Other entries should not be modified
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 no puede estar vacío")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 debe estar vacío")
}

func TestAddNewLocaleEntries(t *testing.T) {

	locale := &Locale{
		ErrorKeyNotBlank: "{{title}} can't be blank (XX)",
		ErrorKeyBlank:    "{{title}} must be blank (XX)",
	}

	v := New(Options{LocaleCode: "xx", Locale: locale}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank (XX)")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 must be blank (XX)")

	// For the unexisting keys, then should use the default language
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 can't be empty")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 must be empty")
}

func TestLocalesIsInValidationScope(t *testing.T) {
	originalErrorMessage0 := (*getLocaleEn())[ErrorKeyNotBlank]
	modifiedErrorMessage0 := "{{title}} should not be blank"

	assert.NotEqual(t, originalErrorMessage0, modifiedErrorMessage0)

	locale := &Locale{
		ErrorKeyNotBlank: modifiedErrorMessage0,
	}

	v := New(Options{Locale: locale}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 should not be blank")

	// New validation
	v = New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank")
}
