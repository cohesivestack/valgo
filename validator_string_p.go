package valgo

import (
	"regexp"
)

// The String pointer validator type that keeps its validator context.
type ValidatorStringP[T ~string] struct {
	context *ValidatorContext
}

// Receives a string pointer to validate.
//
// The value also can be a custom boolean type such as `type Status *string;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`

func StringP[T ~string](value *T, nameAndTitle ...string) *ValidatorStringP[T] {
	return &ValidatorStringP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorStringP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated to the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Blank() function
//	status := ""
//	Is(v.StringP(&status).Not().Blank()).Valid()
func (validator *ValidatorStringP[T]) Not() *ValidatorStringP[T] {
	validator.context.Not()

	return validator
}

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the string is equals "test".
//	input := "test"
//	isValid := v.Is(v.StringP(&input).MinLength(5).Or().EqualTo("test")).Valid()
func (validator *ValidatorStringP[T]) Or() *ValidatorStringP[T] {
	validator.context.Or()

	return validator
}

// Validate if the value of a string pointer is equal to a another value.
// For example:
//
//	status := "running"
//	Is(v.StringP(&status).Equal("running"))
func (validator *ValidatorStringP[T]) EqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a string value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.StringP(&section).GreaterThan("ba"))
func (validator *ValidatorStringP[T]) GreaterThan(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringGreaterThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.StringP(&section).GreaterOrEqualTo("bc"))

func (validator *ValidatorStringP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringGreaterOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.StringP(&section).LessThan("bc"))
func (validator *ValidatorStringP[T]) LessThan(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringLessThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a string value is less or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.StringP(&section).LessOrEqualTo("bc"))
func (validator *ValidatorStringP[T]) LessOrEqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringLessOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is empty. Empty will be false if the length
// of the string is greater than zero, even if the string has only spaces.
// For checking if the string has only spaces, uses the function `Blank()`
// instead.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).Empty()) // Will be true
//	status = " "
//	Is(v.StringP(&status).Empty()) // Will be false
func (validator *ValidatorStringP[T]) Empty(template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringEmpty(*(validator.context.Value().(*T)))
		},
		ErrorKeyEmpty, validator.context.Value(), template...)

	return validator
}

// Validate if a string value is empty or nil. Empty will be false if the length
// of the string is greater than zero, even if the string has only spaces.
// For checking if the string has only spaces, uses the function `BlankOrNil()`
// instead.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).EmptyOrNil()) // Will be true
//	status = " "
//	Is(v.StringP(&status).EmptyOrNil()) // Will be false
//	var _status *string
//	Is(v.StringP(_status).EmptyOrNil()) // Will be true
func (validator *ValidatorStringP[T]) EmptyOrNil(template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil || isStringEmpty(*(validator.context.Value().(*T)))
		},
		ErrorKeyEmpty, validator.context.Value(), template...)

	return validator
}

// Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).Blank()) // Will be true
//	status = " "
//	Is(v.StringP(&status).Blank()) // Will be true
func (validator *ValidatorStringP[T]) Blank(template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringBlank(*(validator.context.Value().(*T)))
		},
		ErrorKeyBlank, validator.context.Value(), template...)

	return validator
}

// Validate if a string value is blank or nil. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).BlankOrNil()) // Will be true
//	status = " "
//	Is(v.StringP(&status).BlankOrNil()) // Will be true
//	var _status *string
//	Is(v.StringP(_status).BlankOrNil()) // Will be true
func (validator *ValidatorStringP[T]) BlankOrNil(template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil || isStringBlank(*(validator.context.Value().(*T)))
		},
		ErrorKeyBlank, validator.context.Value(), template...)

	return validator
}

// Validate if a string pointer pass a custom function.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorStringP[T]) Passing(function func(v0 *T) bool, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if the value of a string pointer is present in a string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.StringP(&status).InSlice(validStatus))
func (validator *ValidatorStringP[T]) InSlice(slice []T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Validate if the value of a string pointer match a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex, _ := regexp.Compile("pre-.+")
//	Is(v.StringP(&status).MatchingTo(regex))
func (validator *ValidatorStringP[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringMatchingTo(*(validator.context.Value().(*T)), regex)
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the maximum length (in bytes) of a string pointer's value.
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).MaxBytes(6))
//
// For character count, use `MaxLength` instead.
func (validator *ValidatorStringP[T]) MaxBytes(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringByteMaxLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the minimum length (in bytes) of a string pointer's value.
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).MinBytes(6))
//
// For character count, use `MinLength` instead.
func (validator *ValidatorStringP[T]) MinBytes(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringByteMinLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the length (in bytes) of a string pointer's value.
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).OfByteLength(6))
//
// For character count, use `OfLength` instead.
func (validator *ValidatorStringP[T]) OfByteLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringByteLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if the length (in bytes) of a string pointer's value is in a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).OfByteLengthBetween(2,6))
//
// For character count, use `OfLengthBetween` instead.
func (validator *ValidatorStringP[T]) OfByteLengthBetween(min int, max int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringByteLengthBetween(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the maximum length (in runes/characters) of a string pointer's value.
// For example:
//
//	slug := "虎視眈々" // 4 runes, len(slug) = 12 bytes
//	Is(v.StringP(&slug).MaxLength(4))
func (validator *ValidatorStringP[T]) MaxLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringRuneMaxLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the minimum length (in runes/characters) of a string pointer's value.
// For example:
//
//	slug := "虎視眈々" // 4 runes, len(slug) = 12 bytes
//	Is(v.StringP(&slug).MinLength(4))
func (validator *ValidatorStringP[T]) MinLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringRuneMinLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate the length (in runes/characters) of a string pointer's value.
// For example:
//
//	slug := "虎視眈々" // 4 runes, len(slug) = 12 bytes
//	Is(v.StringP(&slug).OfLength(4))
func (validator *ValidatorStringP[T]) OfLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringRuneLength(*(validator.context.Value().(*T)), length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if the length (in runes/characters) of a string pointer's value is in a range (inclusive).
// For example:
//
//	slug := "虎視眈々" // 4 runes, len(slug) = 12 bytes
//	Is(v.StringP(&slug).OfLengthBetween(2,4))
func (validator *ValidatorStringP[T]) OfLengthBetween(min int, max int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			if validator.context.Value().(*T) == nil {
				return false
			}
			return isStringRuneLengthBetween(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if the value of a string pointer's value is in a range (inclusive).
// For example:
//
//	slug := "ab"
//	Is(v.StringP(&slug).Between("ab","ac"))
func (validator *ValidatorStringP[T]) Between(min T, max T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isStringBetween(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if a string pointer is nil.
// For example:
//
//	var status *string
//	Is(v.StringP(status).Nil())
func (validator *ValidatorStringP[T]) Nil(template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}
