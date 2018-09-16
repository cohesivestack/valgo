package valgo

import (
	"errors"
	"fmt"
	"log"
)

var locales map[string]locale

var defaultLocaleCode string

func init() {
	err := SetDefaultLocale("en")
	if err != nil {
		log.Fatal(err)
	}
}

func getLocales() map[string]locale {
	if locales == nil {
		locales = map[string]locale{}
	}
	return locales
}

func SetDefaultLocale(code string) error {
	if _, exist := getLocales()[code]; exist {
		defaultLocaleCode = code
		return nil
	} else {
		return errors.New(fmt.Sprintf("There is not a locale registered with code %s", code))
	}

}

func AddOrReplaceMessages(code string, messages map[string]string) {
	_locale := locale{
		Messages: messages,
	}

	getLocales()[code] = _locale
}

func GetMessagesCopy(code string) (map[string]string, error) {
	if locale, ok := getLocales()[code]; ok {
		messages := map[string]string{}
		for k, v := range locale.Messages {
			messages[k] = v
		}
		return messages, nil
	} else {
		return nil, errors.New(fmt.Sprintf("There is not a locale with the code '%s'", code))
	}
}

func ResetMessages() {
	setDefaultEnglishMessages()
	setDefaultSpanishMessages()
	SetDefaultLocale("en")
}

func Localized(code string) (*localized, error) {
	if _locale, exist := getLocales()[code]; exist {
		return &localized{
			_locale: _locale,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Doesn't exist a registered locale with code '%s'", code))
	}
}

func newValidator(_locale locale, value interface{}) *Validator {
	validator := &Validator{
		currentValid: true,
		valid:        true,
		_locale:      _locale,
	}

	return validator.Is(value)
}

func Is(value interface{}) *Validator {
	return newValidator(getLocales()[defaultLocaleCode], value)
}

func AddErrorToNamed(name string, message string) *Validator {
	return Is(nil).Named(name).WithError(message)
}
