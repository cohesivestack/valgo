package valgo

import (
	"errors"
	"fmt"
)

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

func newValidator(_locale *locale) Validator {
	v := &validatorContext{
		valid:   true,
		_locale: _locale,
	}

	return v
}

func NewValidator() Validator {
	return newValidator(getDefaultLocale())
}

func AddErrorMessage(name string, message string) Validator {
	return NewValidator().AddErrorMessage(name, message)
}
