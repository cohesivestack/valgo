package valgo

import (
	"fmt"
)

var customMarshalJson func(e *Error) ([]byte, error)

func TeardownTest() {
	SetMarshalJSON(nil)
	setDefaultEnglishMessages()
	setDefaultSpanishMessages()
	SetDefaultLocale("en")
}

// Create a localized [Validation] factory.
func Localized(code string) (*localized, error) {
	if _locale, exist := getLocales()[code]; exist {
		return &localized{
			_locale: _locale,
		}, nil
	} else {
		return nil, fmt.Errorf("doesn't exist a registered locale with code '%s'", code)
	}
}

// Create a new [Validation] session.
func New() *Validation {
	return newValidation(getDefaultLocale())
}

// Create a new [Validation] session and add a field validator to it.
func Is(v Validator) *Validation {
	return New().Is(v)
}

// Create a [Validation] session and add namespace to it.
func In(name string, v *Validation) *Validation {
	return New().In(name, v)
}

// Create a [Validation] session and add an indexed namespace to it.
func InRow(name string, index int, v *Validation) *Validation {
	return New().InRow(name, index, v)
}

// Create a new [Validation] session and add a field validator to it, but unlike
// [Is()] the field validator is not short-circuited.
func Check(v Validator) *Validation {
	return New().Check(v)
}

// Create a new [Validation] session and add an error message to it without
// executing a field validator. By adding this error message, the [Validation]
// session will be marked as invalid.
func AddErrorMessage(name string, message string) *Validation {
	return New().AddErrorMessage(name, message)
}

// Set a custom function to serialize the validator messages as JSON.
func SetMarshalJSON(customFunc func(e *Error) ([]byte, error)) {
	customMarshalJson = customFunc
}
