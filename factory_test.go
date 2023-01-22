package valgo

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryUseLocaleCodeDefault(t *testing.T) {

	factory := Factory(FactoryOptions{
		LocaleCodeDefault: LocaleCodeEs,
	})

	v := factory.New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent in the same validation
	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 debe estar vacío")

	// Default localization must be persistent in the other validations
	v = factory.New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be possible to change at Validation level
	v = factory.New(Options{LocaleCode: LocaleCodeEn}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank")
}

func TestFactoryChangeLocaleEntries(t *testing.T) {

	originalErrorMessage0 := getLocaleEn().Messages[ErrorKeyNotBlank]
	modifiedErrorMessage0 := "{{title}} should not be blank"

	originalErrorMessage1 := getLocaleEn().Messages[ErrorKeyBlank]
	modifiedErrorMessage1 := "{{title}} should be blank"

	assert.NotEqual(t, originalErrorMessage0, modifiedErrorMessage0)
	assert.NotEqual(t, originalErrorMessage1, modifiedErrorMessage1)

	locale := &Locale{Messages: map[string]string{
		ErrorKeyNotBlank: modifiedErrorMessage0,
		ErrorKeyBlank:    modifiedErrorMessage1,
	}}

	factory := Factory(FactoryOptions{
		Locales: map[string]*Locale{
			localeCodeDefault: locale,
		},
	})

	v := factory.New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 should not be blank")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 should be blank")

	// Other entries should not be modified
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 can't be empty")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 must be empty")
}

func TestFactoryUseOtherLocaleAndChangeLocaleEntries(t *testing.T) {

	originalErrorMessage0 := getLocaleEs().Messages[ErrorKeyNotBlank]
	modifiedErrorMessage0 := "{{title}} no debería estar en blanco"

	originalErrorMessage1 := getLocaleEs().Messages[ErrorKeyBlank]
	modifiedErrorMessage1 := "{{title}} debería estar en blanco"

	assert.NotEqual(t, originalErrorMessage0, modifiedErrorMessage0)
	assert.NotEqual(t, originalErrorMessage1, modifiedErrorMessage1)

	locale := &Locale{Messages: map[string]string{
		ErrorKeyNotBlank: modifiedErrorMessage0,
		ErrorKeyBlank:    modifiedErrorMessage1,
	}}

	factory := Factory(FactoryOptions{
		Locales: map[string]*Locale{
			LocaleCodeEs: locale,
		},
	})

	v := factory.New(Options{LocaleCode: LocaleCodeEs}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 no debería estar en blanco")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 debería estar en blanco")

	// Other entries should not be modified
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 no puede estar vacío")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 debe estar vacío")
}

func TestFactoryAddNewLocaleEntries(t *testing.T) {

	locale := &Locale{Messages: map[string]string{
		ErrorKeyNotBlank: "{{title}} can't be blank (XX)",
		ErrorKeyBlank:    "{{title}} must be blank (XX)",
	}}

	factory := Factory(FactoryOptions{
		Locales: map[string]*Locale{
			"xx": locale,
		},
	})

	v := factory.New(Options{LocaleCode: "xx"}).Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank (XX)")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 must be blank (XX)")

	// For the unexisting keys, then should use the default language
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 can't be empty")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 must be empty")

	// Use new locale Entries but changing the default in the Factory
	factory = Factory(FactoryOptions{
		LocaleCodeDefault: LocaleCodeEs,
		Locales: map[string]*Locale{
			"xx": locale,
		},
	})

	v = factory.New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank (XX)")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 must be blank (XX)")

	// For the unexisting keys, then should use the default language
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 no puede estar vacío")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 debe estar vacío")

	// Use new locale Entries but changing the default in the Factory to be the
	// same new unexisting locale. That will use the Valgo default locale ("en")
	factory = Factory(FactoryOptions{
		LocaleCodeDefault: "xx",
		Locales: map[string]*Locale{
			"xx": locale,
		},
	})

	v = factory.New().Is(String(" ").Not().Blank())
	assert.Contains(t, v.Errors()["value_0"].Messages(), "Value 0 can't be blank (XX)")

	v = v.Is(String("a").Blank())
	assert.Contains(t, v.Errors()["value_1"].Messages(), "Value 1 must be blank (XX)")

	// For the unexisting keys, then should use the default language
	v = v.Is(String("").Not().Empty())
	assert.Contains(t, v.Errors()["value_2"].Messages(), "Value 2 can't be empty")

	v = v.Is(String(" ").Empty())
	assert.Contains(t, v.Errors()["value_3"].Messages(), "Value 3 must be empty")
}

func TestFactoryCustomErrorMarshallJSON(t *testing.T) {

	customFunc := func(e *Error) ([]byte, error) {

		errors := map[string]interface{}{}

		for k, v := range e.errors {
			if len(v.Messages()) == 1 {
				errors[k] = v.Messages()[0]
			} else {
				errors[k] = v.Messages()
			}
		}

		// Add root level errors to customize errors interface
		return json.Marshal(map[string]map[string]interface{}{"errors": errors})
	}

	factory := Factory(FactoryOptions{
		MarshalJsonFunc: customFunc,
	})

	r, _ := regexp.Compile("a")
	v := factory.New().
		Check(String("", "email").Not().Blank().MatchingTo(r)).
		Check(String("", "name").Not().Blank())

	jsonByte, err := json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMap := map[string]map[string]interface{}{}
	err = json.Unmarshal(jsonByte, &jsonMap)
	assert.NoError(t, err)

	assert.Equal(t, "Name can't be blank", jsonMap["errors"]["name"])
	emailErrors := jsonMap["errors"]["email"].([]interface{})
	assert.Contains(t, emailErrors, "Email can't be blank")
	assert.Contains(t, emailErrors, "Email must match to \"a\"")

	customFuncAtValidationLevel := func(e *Error) ([]byte, error) {

		errors := map[string]string{"errors": "overridden"}

		// Add root level errors to customize errors interface
		return json.Marshal(errors)
	}

	// Marshal JSON should be overridden at Validation level
	v = factory.New(Options{MarshalJsonFunc: customFuncAtValidationLevel}).
		Check(String("", "email").Not().Blank().MatchingTo(r)).
		Check(String("", "name").Not().Blank())

	jsonByte, err = json.Marshal(v.Error())
	assert.NoError(t, err)

	jsonMapAtValidationLevel := map[string]string{}
	err = json.Unmarshal(jsonByte, &jsonMapAtValidationLevel)
	assert.NoError(t, err)

	assert.Equal(t, "overridden", jsonMapAtValidationLevel["errors"])
}
