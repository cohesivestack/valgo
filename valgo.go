package valgo

import (
	"fmt"
)

var customMarshalJson func(e *Error) ([]byte, error)

func Teardown() {
	SetMarshalJSON(nil)
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
	return New().Is(v)
}

func In(name string, v *Validation) *Validation {
	return New().In(name, v)
}

func InRow(name string, index int, v *Validation) *Validation {
	return New().InRow(name, index, v)
}

func Check(v Validator) *Validation {
	return New().Check(v)
}

func AddErrorMessage(name string, message string) *Validation {
	return New().AddErrorMessage(name, message)
}

func SetMarshalJSON(customFunc func(e *Error) ([]byte, error)) {
	customMarshalJson = customFunc
}
