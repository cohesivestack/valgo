package valgo

import (
	"regexp"
	"strings"
)

// ValidatorString The `ValidatorString` provides functions for setting validation rules for
// a string value type, or a custom type based on a string.
type ValidatorString[T ~string] struct {
	context *ValidatorContext
}

// String Receive a string value to validate.
//
// The value can also be a custom string type such as type Status string;.
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N` pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.
func String[T ~string](value T, nameAndTitle ...string) *ValidatorString[T] {
	return &ValidatorString[T]{context: NewContext(value, nameAndTitle...)}
}

// Context Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorString[T]) Context() *ValidatorContext {
	return validator.context
}

// Not Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Blank() function
//	Is(v.String("").Not().Blank()).Valid()
func (validator *ValidatorString[T]) Not() *ValidatorString[T] {
	validator.context.Not()

	return validator
}

// EqualTo Validate if a string value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.String(status).Equal("running"))
func (validator *ValidatorString[T]) EqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringEqualTo(val, value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// GreaterThan Validate if a string value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).GreaterThan("ba"))
func (validator *ValidatorString[T]) GreaterThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringGreaterThan(val, value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// GreaterOrEqualTo Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.String(section).GreaterOrEqualTo("bc"))
func (validator *ValidatorString[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringGreaterOrEqualTo(val, value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// LessThan Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).LessThan("bc"))
func (validator *ValidatorString[T]) LessThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringLessThan(val, value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// LessOrEqualTo Validate if a string value is less than or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.String(section).LessOrEqualTo("bc"))
func (validator *ValidatorString[T]) LessOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringLessOrEqualTo(val, value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Empty Validate if a string value is empty. Return false if the length of the string
// is greater than zero, even if the string has only spaces.
//
// For checking if the string has only spaces, use the function `Blank()`
// instead.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be false
func (validator *ValidatorString[T]) Empty(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringEmpty(val)
		},
		ErrorKeyEmpty, template...)

	return validator
}

// Blank Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be true
func (validator *ValidatorString[T]) Blank(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringBlank(val)
		},
		ErrorKeyBlank, template...)

	return validator
}

// Passing Validate if a string value passes a custom function.
// For example:
//
//	status := ""
//	Is(v.String(status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorString[T]) Passing(function func(v0 T) bool, template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return function(val)
		},
		ErrorKeyPassing, template...)

	return validator
}

// InSlice Validate if a string is present in a string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.String(status).InSlice(validStatus))
func (validator *ValidatorString[T]) InSlice(slice []T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringInSlice(val, slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// MatchingTo Validate if a string matches a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex := regexp.MustCompile("pre-.+")
//	Is(v.String(status).MatchingTo(regex))
func (validator *ValidatorString[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringMatchingTo(val, regex)
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex},
		template...)

	return validator
}

// MaxLength Validate the maximum length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MaxLength(6))
func (validator *ValidatorString[T]) MaxLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringMaxLength(val, length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// MinLength Validate the minimum length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MinLength(6))
func (validator *ValidatorString[T]) MinLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringMinLength(val, length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// OfLength Validate the length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).OfLength(6))
func (validator *ValidatorString[T]) OfLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringLength(val, length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// OfLengthBetween Validate if the length of a string is within a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).OfLengthBetween(2,6))
func (validator *ValidatorString[T]) OfLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringLengthBetween(val, min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Between Validate if the value of a string is within a range (inclusive).
// For example:
//
//	slug := "ab"
//	Is(v.String(slug).Between("ab","ac"))
func (validator *ValidatorString[T]) Between(min T, max T, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(T)
			if !ok {
				return false
			}

			return isStringBetween(val, min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

func isStringEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 == v1
}

func isStringGreaterThan[T ~string](v0 T, v1 T) bool {
	return v0 > v1
}

func isStringGreaterOrEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 >= v1
}

func isStringLessThan[T ~string](v0 T, v1 T) bool {
	return v0 < v1
}

func isStringLessOrEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 <= v1
}

func isStringBetween[T ~string](v T, min T, max T) bool {
	return v >= min && v <= max
}

func isStringEmpty[T ~string](v T) bool {
	return len(v) == 0
}

func isStringBlank[T ~string](v T) bool {
	return len(strings.TrimSpace(string(v))) == 0
}

func isStringInSlice[T ~string](v T, slice []T) bool {
	for _, _v := range slice {
		if v == _v {
			return true
		}
	}

	return false
}

func isStringMatchingTo[T ~string](v T, regex *regexp.Regexp) bool {
	return regex.MatchString(string(v))
}

func isStringMaxLength[T ~string](v T, length int) bool {
	return len(v) <= length
}

func isStringMinLength[T ~string](v T, length int) bool {
	return len(v) >= length
}

func isStringLength[T ~string](v T, length int) bool {
	return len(v) == length
}

func isStringLengthBetween[T ~string](v T, min int, max int) bool {
	return len(v) >= min && len(v) <= max
}
