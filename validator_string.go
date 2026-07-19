package valgo

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

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
func isStringByteMaxLength[T ~string](v T, length int) bool {
	return len(v) <= length
}
func isStringByteMinLength[T ~string](v T, length int) bool {
	return len(v) >= length
}
func isStringByteLength[T ~string](v T, length int) bool {
	return len(v) == length
}
func isStringByteLengthBetween[T ~string](v T, min int, max int) bool {
	return len(v) >= min && len(v) <= max
}
func isStringRuneMaxLength[T ~string](v T, length int) bool {
	return utf8.RuneCountInString(string(v)) <= length
}
func isStringRuneMinLength[T ~string](v T, length int) bool {
	return utf8.RuneCountInString(string(v)) >= length
}
func isStringRuneLength[T ~string](v T, length int) bool {
	return utf8.RuneCountInString(string(v)) == length
}
func isStringRuneLengthBetween[T ~string](v T, min int, max int) bool {
	l := utf8.RuneCountInString(string(v))
	return l >= min && l <= max
}

// The `ValidatorString` provides functions for setting validation rules for
// a string value type, or a custom type based on a string.
type ValidatorString[T ~string] struct {
	context *ValidatorContext
}

// Receive a string value to validate.
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

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorString[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Blank() function
//	Is(v.String("").Not().Blank()).Valid()
func (validator *ValidatorString[T]) Not() *ValidatorString[T] {
	validator.context.Not()

	return validator
}

// Or introduces a logical OR boundary in the current validator chain.
//
// Or groups adjacent validation fragments into a single OR-group that is
// evaluated left-to-right until one fragment succeeds. The OR-group succeeds
// if any fragment succeeds; it fails only if all fragments fail.
//
// Precedence: the OR-group is evaluated as a unit before the implicit AND
// that continues the chain. For example:
//
//	A.Or().B.C   == (A OR B) AND C
//
// Error reporting: if the OR-group fails, the error message for that group is
// a single message composed by joining the failing fragments' messages using
// the localized OR list format.
//
// Example:
//
//	// Passes because input equals "test" (MinLength(5) OR EqualTo("test")).
//	input := "test"
//	isValid := v.Is(v.String(input).MinLength(5).Or().EqualTo("test")).Valid()
func (validator *ValidatorString[T]) Or() *ValidatorString[T] {
	validator.context.Or()

	return validator
}

// OrElse introduces a logical OR boundary with a cut (short-circuit) in the
// validator chain.
//
// OrElse behaves like Or for building an OR-group, but with an additional rule:
// if the left side (a single fragment, or the entire OR-group accumulated to
// the left of OrElse) succeeds, validation stops and no fragments to the right
// of OrElse are evaluated.
//
// This is primarily used to express "accept X, otherwise validate the rest"
// without repeating X across multiple OR fragments.
//
// Precedence: OrElse still participates in OR-grouping precedence. For example:
//
//	A.OrElse().B.C  == A OR (B AND C)   (with a cut if A succeeds)
//
// Error reporting: if the OR-group fails, its message is composed the same way
// as Or (localized OR list join).
//
// Example:
//
//	// If input is empty, the chain succeeds and MinLength/EqualTo are not evaluated.
//	// Otherwise, input must have MinLength(5) AND EqualTo("test").
//	input := ""
//	isValid := v.Is(v.String(input).Empty().OrElse().MinLength(5).EqualTo("test")).Valid()
func (validator *ValidatorString[T]) OrElse() *ValidatorString[T] {
	validator.context.OrElse()

	return validator
}

// Validate if a string value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.String(status).Equal("running"))
func (validator *ValidatorString[T]) EqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a string value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).GreaterThan("ba"))
func (validator *ValidatorString[T]) GreaterThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringGreaterThan(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.String(section).GreaterOrEqualTo("bc"))
func (validator *ValidatorString[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringGreaterOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).LessThan("bc"))
func (validator *ValidatorString[T]) LessThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringLessThan(validator.context.Value().(T), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a string value is less than or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.String(section).LessOrEqualTo("bc"))
func (validator *ValidatorString[T]) LessOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringLessOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is empty. Return false if the length of the string
// is greater than zero, even if the string has only spaces.
//
// For checking if the string has only spaces, use the function `Blank()`
// instead.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be false
func (validator *ValidatorString[T]) Empty(template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringEmpty(validator.context.Value().(T))
		},
		ErrorKeyEmpty, validator.context.Value(), template...)

	return validator
}

// Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be true
func (validator *ValidatorString[T]) Blank(template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringBlank(validator.context.Value().(T))
		},
		ErrorKeyBlank, validator.context.Value(), template...)

	return validator
}

// Validate if a string value passes a custom function.
// For example:
//
//	status := ""
//	Is(v.String(status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorString[T]) Passing(function func(v0 T) bool, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a string is present in a string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.String(status).InSlice(validStatus))
func (validator *ValidatorString[T]) InSlice(slice []T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Validate if a string matches a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex, _ := regexp.Compile("pre-.+")
//	Is(v.String(status).MatchingTo(regex))
func (validator *ValidatorString[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringMatchingTo(validator.context.Value().(T), regex)
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the maximum length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MaxBytes(6))
//
// For character count, use `MaxLength` instead.
func (validator *ValidatorString[T]) MaxBytes(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringByteMaxLength(validator.context.Value().(T), length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// Validate the minimum length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MinBytes(6))
//
// For character count, use `MinLength` instead.
func (validator *ValidatorString[T]) MinBytes(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringByteMinLength(validator.context.Value().(T), length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).ByteLength(6))
//
// For character count, use `Length` instead.
func (validator *ValidatorString[T]) ByteLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringByteLength(validator.context.Value().(T), length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// OfByteLength validates the length (in bytes) of a string.
//
// Deprecated: use ByteLength instead. OfByteLength will be removed in v1.0.
func (validator *ValidatorString[T]) OfByteLength(length int, template ...string) *ValidatorString[T] {
	return validator.ByteLength(length, template...)
}

// Validate if the length (in bytes) of a string is within a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).ByteLengthBetween(2,6))
//
// For character count, use `LengthBetween` instead.
func (validator *ValidatorString[T]) ByteLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringByteLengthBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// OfByteLengthBetween validates if the length (in bytes) of a string is within
// a range (inclusive).
//
// Deprecated: use ByteLengthBetween instead. OfByteLengthBetween will be removed in v1.0.
func (validator *ValidatorString[T]) OfByteLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	return validator.ByteLengthBetween(min, max, template...)
}

// Validate the maximum length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).MaxLength(4))
func (validator *ValidatorString[T]) MaxLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringRuneMaxLength(validator.context.Value().(T), length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the minimum length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).MinLength(4))
func (validator *ValidatorString[T]) MinLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringRuneMinLength(validator.context.Value().(T), length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).Length(4))
func (validator *ValidatorString[T]) Length(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringRuneLength(validator.context.Value().(T), length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// OfLength validates the length (in runes/characters) of a string.
//
// Deprecated: use Length instead. OfLength will be removed in v1.0.
func (validator *ValidatorString[T]) OfLength(length int, template ...string) *ValidatorString[T] {
	return validator.Length(length, template...)
}

// Validate if the length (in runes/characters) of a string is within a range (inclusive).
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).LengthBetween(2,4))
func (validator *ValidatorString[T]) LengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringRuneLengthBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// OfLengthBetween validates if the length (in runes/characters) of a string is
// within a range (inclusive).
//
// Deprecated: use LengthBetween instead. OfLengthBetween will be removed in v1.0.
func (validator *ValidatorString[T]) OfLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	return validator.LengthBetween(min, max, template...)
}

// Validate if the value of a string is within a range (inclusive).
// For example:
//
//	slug := "ab"
//	Is(v.String(slug).Between("ab","ac"))
func (validator *ValidatorString[T]) Between(min T, max T, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}
