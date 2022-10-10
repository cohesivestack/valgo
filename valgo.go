package valgo

import (
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
		return nil, fmt.Errorf("doesn't exist a registered locale with code '%s'", code)
	}
}

func New() *Validation {
	return newValidation(getDefaultLocale())
}

func Is(v Validator) *Validation {
	validation := New()
	return v.Context().validateIs(validation)
}

func Check(v Validator) *Validation {
	validation := New()
	return v.Context().validateCheck(validation)
}

func AddErrorMessage(name string, message string) *Validation {
	return New().AddErrorMessage(name, message)
}
