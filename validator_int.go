package valgo

// The [ValidatorInt] provides functions for setting validation rules for a
// int value types, or a custom type based on a int, int8, int16, int32, or int64.
type ValidatorInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64] struct {
	context *ValidatorContext
}

// Receives an int value to validate.
//
// The value can also be a custom int type such as type Age int.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int[T ~int](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int8 value to validate.
//
// The value can also be a custom int8 type such as type Age int8.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int8[T ~int8](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int16 value to validate.
//
// The value can also be a custom int16 type such as type Age int16.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int16[T ~int16](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int32 value to validate.
//
// The value can also be a custom int32 type such as type Age int32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int32[T ~int32](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int64 value to validate.
//
// The value can also be a custom int64 type such as type Age int64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int64[T ~int64](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives a rune value to validate.
//
// The value can also be a custom rune type such as type Age rune.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Rune[T ~rune](value T, nameAndTitle ...string) *ValidatorInt[T] {
	return &ValidatorInt[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorInt[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	Is(v.Int(0).Not().Zero()).Valid()
func (validator *ValidatorInt[T]) Not() *ValidatorInt[T] {
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
//	// This validator will pass because the input is Zero.
//	input := int(0)
//	isValid := v.Is(v.Int(input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorInt[T]) Or() *ValidatorInt[T] {
	validator.context.Or()

	return validator
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.Int(quantity).EqualTo(2))
func (validator *ValidatorInt[T]) EqualTo(value T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a numeric value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := int(3)
//	Is(v.Int(quantity).GreaterThan(2))
func (validator *ValidatorInt[T]) GreaterThan(value T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberGreaterThan(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a numeric value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := int(3)
//	Is(v.Int(quantity).GreaterOrEqualTo(3))
func (validator *ValidatorInt[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberGreaterOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a numeric value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.Int(quantity).LessThan(3))
func (validator *ValidatorInt[T]) LessThan(value T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberLessThan(validator.context.Value().(T), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a numeric value is less than or equal to another. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.Int(quantity).LessOrEqualTo(2))
func (validator *ValidatorInt[T]) LessOrEqualTo(value T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberLessOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if a number is within a range (inclusive).
// For example:
//
//	Is(v.Int(3).Between(2,6))
func (validator *ValidatorInt[T]) Between(min T, max T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithParams(
		func() bool {
			return isNumberBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if a numeric value is zero.
//
// For example:
//
//	Is(v.Int(0).Zero())
func (validator *ValidatorInt[T]) Zero(template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberZero(validator.context.Value().(T))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is positive (greater than zero).
//
// For example:
//
//	Is(v.Int(5).Positive())
func (validator *ValidatorInt[T]) Positive(template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) > 0
		},
		ErrorKeyPositive, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is negative (less than zero).
//
// For example:
//
//	Is(v.Int(-5).Negative())
func (validator *ValidatorInt[T]) Negative(template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) < 0
		},
		ErrorKeyNegative, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := int(2)
//	Is(v.Int(quantity).Passing((v int) bool {
//		return v == getAllowedQuantity()
//	})
func (validator *ValidatorInt[T]) Passing(function func(v T) bool, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a number is present in a numeric slice.
// For example:
//
//	quantity := int(3)
//	validQuantities := []int{1,3,5}
//	Is(v.Int(quantity).InSlice(validQuantities))
func (validator *ValidatorInt[T]) InSlice(slice []T, template ...string) *ValidatorInt[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
