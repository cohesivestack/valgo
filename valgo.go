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

func New() *ValidatorGroup {
	return newValidatorGroup(getDefaultLocale())
}

func Is(v Validator) *ValidatorGroup {
	group := New()
	return v.Context().validateIs(group)
}

func Check(v Validator) *ValidatorGroup {
	group := New()
	return v.Context().validateCheck(group)
}

func AddErrorMessage(name string, message string) *ValidatorGroup {
	return New().AddErrorMessage(name, message)
}
