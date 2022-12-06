// Package valgo Valgo is a type-safe, expressive, and extensible validator library for
// Golang. Valgo is built with generics, so Go 1.18 or higher is required.
//
// Valgo differs from other Golang validation libraries in that the rules are
// written in functions and not in struct tags. This allows greater flexibility
// and freedom when it comes to where and how data is validated.
//
// Additionally, Valgo supports customizing and localizing validation messages.
package valgo

var customMarshalJSON func(e *Error) ([]byte, error)

// Localized Create a localized [Validation] factory.
//
//nolint:revive // by design. should be exported as can be annoying to use
func Localized(code string) (*localized, error) {
	_locale := getLocale(code)
	if _locale == nil {
		return nil, localeDoesNotExist(code)
	}

	return &localized{_locale: _locale}, nil
}

// New This function allows you to create a new Validation session without a
// Validator. This is useful for conditional validation or reusing validation
// logic.
//
// The following example conditionally adds a validator rule for the month_day
// value.
func New() *Validation {
	return newValidation(getDefaultLocale())
}

// Is The [Is](...) function allows you to pass a [Validator] with the value and
// the rules for validating it. At the same time, create a [Validation] session,
// which lets you add more Validators in order to verify more values.
//
// As shown in the following example, we are passing to the function [Is](...)
// the [Validator] for the full_name value. The function returns a [Validation]
// session that allows us to add more Validators to validate more values; in the
// example case the values age and status:.
func Is(v Validator) *Validation {
	return New().Is(v)
}

// In The [In](...) function executes one or more validators in a namespace, so the
// value names in the error result are prefixed with this namespace. This is
// useful for validating nested structures.
//
// In the following example we are validating the Person and the nested
// Address structure. We can distinguish the errors of the nested Address
// structure in the error results.
func In(name string, v *Validation) *Validation {
	return New().In(name, v)
}

// InRow The [InRow](...) function executes one or more validators in a namespace
// similar to the [In](...) function, but with indexed namespace. So, the value
// names in the error result are prefixed with this indexed namespace. It is
// useful for validating nested lists in structures.
//
// In the following example we validate the Person and the nested list
// Addresses. The error results can distinguish the errors of the nested list
// Addresses.
func InRow(name string, index int, v *Validation) *Validation {
	return New().InRow(name, index, v)
}

// Check The [Check](...) function, similar to the [Is](...) function, however with
// [Check](...)` the Rules of the [Validator] parameter are not short-circuited,
// which means that regardless of whether a previous rule was valid, all rules
// are checked.
//
// This example shows two rules that fail due to the empty value in the full_name
// [Validator], and since the [Validator] is not short-circuited, both error
// messages are added to the error result.
func Check(v Validator) *Validation {
	return New().Check(v)
}

// AddErrorMessage Create a new [Validation] session and add an error message to it without
// executing a field validator. By adding this error message, the [Validation]
// session will be marked as invalid.
func AddErrorMessage(name string, message string) *Validation {
	return New().AddErrorMessage(name, message)
}

// SetMarshalJSON Set a custom function to serialize the validator messages as JSON.
func SetMarshalJSON(customFunc func(e *Error) ([]byte, error)) {
	customMarshalJSON = customFunc
}
