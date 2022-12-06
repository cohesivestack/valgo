package valgo

import (
	"regexp"
)

// ValidatorStringP The String pointer validator type that keeps its validator context.
type ValidatorStringP[T ~string] struct {
	context *ValidatorContext
}

// StringP Receives a string pointer to validate.
//
// The value also can be a custom boolean type such as `type Status *string;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`.
func StringP[T ~string](value *T, nameAndTitle ...string) *ValidatorStringP[T] {
	return &ValidatorStringP[T]{context: NewContext(value, nameAndTitle...)}
}

// Context Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorStringP[T]) Context() *ValidatorContext {
	return validator.context
}

// Not Invert the logical value associated to the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Blank() function
//	status := ""
//	Is(v.StringP(&status).Not().Blank()).Valid()
func (validator *ValidatorStringP[T]) Not() *ValidatorStringP[T] {
	validator.context.Not()

	return validator
}

// EqualTo Validate if the value of a string pointer is equal to a another value.
// For example:
//
//	status := "running"
//	Is(v.StringP(&status).Equal("running"))
func (validator *ValidatorStringP[T]) EqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringEqualTo(*val, value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// GreaterThan Validate if a string value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.StringP(&section).GreaterThan("ba"))
func (validator *ValidatorStringP[T]) GreaterThan(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringGreaterThan(*val, value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// GreaterOrEqualTo Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.StringP(&section).GreaterOrEqualTo("bc"))
func (validator *ValidatorStringP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringGreaterOrEqualTo(*val, value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// LessThan Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.StringP(&section).LessThan("bc"))
func (validator *ValidatorStringP[T]) LessThan(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringLessThan(*val, value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// LessOrEqualTo Validate if a string value is less or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.StringP(&section).LessOrEqualTo("bc"))
func (validator *ValidatorStringP[T]) LessOrEqualTo(value T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringLessOrEqualTo(*val, value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Empty Validate if a string value is empty. Empty will be false if the length
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
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringEmpty(*val)
		},
		ErrorKeyEmpty, template...)

	return validator
}

// EmptyOrNil Validate if a string value is empty or nil. Empty will be false if the length
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
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val == nil || isStringEmpty(*val)
		},
		ErrorKeyEmpty, template...)

	return validator
}

// Blank Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).Blank()) // Will be true
//	status = " "
//	Is(v.StringP(&status).Blank()) // Will be true
func (validator *ValidatorStringP[T]) Blank(template ...string) *ValidatorStringP[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringBlank(*val)
		},
		ErrorKeyBlank, template...)

	return validator
}

// BlankOrNil Validate if a string value is blank or nil. Blank will be true if the length
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
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val == nil || isStringBlank(*val)
		},
		ErrorKeyBlank, template...)

	return validator
}

// Passing Validate if a string pointer pass a custom function.
// For example:
//
//	status := ""
//	Is(v.StringP(&status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorStringP[T]) Passing(function func(v0 *T) bool, template ...string) *ValidatorStringP[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return function(val)
		},
		ErrorKeyPassing, template...)

	return validator
}

// InSlice Validate if the value of a string pointer is present in a string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.StringP(&status).InSlice(validStatus))
func (validator *ValidatorStringP[T]) InSlice(slice []T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithValue(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringInSlice(*val, slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// MatchingTo Validate if the value of a string pointer match a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex := regexp.MustCompile("pre-.+")
//	Is(v.StringP(&status).MatchingTo(regex))
func (validator *ValidatorStringP[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringMatchingTo(*val, regex)
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex},
		template...)

	return validator
}

// MaxLength Validate if the maximum length of a string pointer's value.
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).MaxLength(6))
func (validator *ValidatorStringP[T]) MaxLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringMaxLength(*val, length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// MinLength Validate the minimum length of a string pointer's value
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).MinLength(6))
func (validator *ValidatorStringP[T]) MinLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringMinLength(*val, length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// OfLength Validate the length of a string pointer's value.
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).OfLength(6))
func (validator *ValidatorStringP[T]) OfLength(length int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringLength(*val, length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// OfLengthBetween Validate if the length of a string pointer's value is in a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).OfLengthBetween(2,6))
func (validator *ValidatorStringP[T]) OfLengthBetween(min int, max int, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringLengthBetween(*val, min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Between Validate if the value of a string is in a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.StringP(&slug).Between(2,6))
func (validator *ValidatorStringP[T]) Between(min T, max T, template ...string) *ValidatorStringP[T] {
	validator.context.AddWithParams(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val != nil && isStringBetween(*val, min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Nil Validate if a string pointer is nil.
// For example:
//
//	var status *string
//	Is(v.StringP(status).Nil())
func (validator *ValidatorStringP[T]) Nil(template ...string) *ValidatorStringP[T] {
	validator.context.Add(
		func() bool {
			val, ok := validator.context.Value().(*T)
			if !ok {
				return false
			}

			return val == nil
		},
		ErrorKeyNil, template...)

	return validator
}
