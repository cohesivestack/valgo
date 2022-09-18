package valgo

import (
	"regexp"
	"strings"
)

type ValidatorString[T ~string] struct {
	context *ValidatorContext
}

func String[T ~string](value T, nameAndTitle ...string) *ValidatorString[T] {
	return &ValidatorString[T]{context: NewContext(value, nameAndTitle...)}
}

func (validator *ValidatorString[T]) Context() *ValidatorContext {
	return validator.context
}

func (validator *ValidatorString[T]) Not() *ValidatorString[T] {
	validator.context.Not()

	return validator
}

func (validator *ValidatorString[T]) EqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorString[T]) GreaterThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) > value
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

func (validator *ValidatorString[T]) GreaterOrEqualThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) >= value
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

func (validator *ValidatorString[T]) LessThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) < value
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

func (validator *ValidatorString[T]) LessOrEqualThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) <= value
		},
		ErrorKeyLessOrEqualThan, value, template...)

	return validator
}

func (validator *ValidatorString[T]) Empty(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return len(validator.context.Value().(T)) == 0
		},
		ErrorKeyEmpty, template...)

	return validator
}

func (validator *ValidatorString[T]) Blank(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return len(strings.TrimSpace(validator.context.Value().(string))) == 0
		},
		ErrorKeyBlank, template...)

	return validator
}

func (validator *ValidatorString[T]) Passing(function func(v0 T) bool, template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, template...)

	return validator
}

func (validator *ValidatorString[T]) InSlice(slice []T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			for _, v := range slice {
				if validator.context.Value() == v {
					return true
				}
			}
			return false
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

func (validator *ValidatorString[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return regex.MatchString(validator.context.Value().(string))
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex},
		template...)

	return validator
}

func (validator *ValidatorString[T]) MaxLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return len(validator.context.Value().(T)) <= length
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

func (validator *ValidatorString[T]) MinLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return len(validator.context.Value().(T)) >= length
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

func (validator *ValidatorString[T]) Length(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return len(validator.context.Value().(T)) == length
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}
