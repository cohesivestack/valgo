package valgo

// The [ValidatorIntP] provides functions for setting validation rules for a
// int pointer value types, or a custom type based on a int, int8, int16, int32, or int64 pointer.
type ValidatorIntP[T ~int | ~int8 | ~int16 | ~int32 | ~int64] struct {
	context *ValidatorContext
}

// Receives an int pointer value to validate.
//
// The value can also be a custom int pointer type such as type Age *int.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func IntP[T ~int](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int8 pointer value to validate.
//
// The value can also be a custom int8 pointer type such as type Age *int8.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int8P[T ~int8](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int16 pointer value to validate.
//
// The value can also be a custom int16 pointer type such as type Age *int16.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int16P[T ~int16](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int32 pointer value to validate.
//
// The value can also be a custom int32 pointer type such as type Age *int32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int32P[T ~int32](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an int64 pointer value to validate.
//
// The value can also be a custom int64 pointer type such as type Age *int64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Int64P[T ~int64](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives a rune pointer value to validate.
//
// The value can also be a custom rune pointer type such as type Age *rune.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func RuneP[T ~rune](value *T, nameAndTitle ...string) *ValidatorIntP[T] {
	return &ValidatorIntP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorIntP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	n := int(0)
//	Is(v.IntP(&n).Not().Zero()).Valid()
func (validator *ValidatorIntP[T]) Not() *ValidatorIntP[T] {
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
//	isValid := v.Is(v.IntP(&input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorIntP[T]) Or() *ValidatorIntP[T] {
	validator.context.Or()

	return validator
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.IntP(&quantity).EqualTo(2))
func (validator *ValidatorIntP[T]) EqualTo(value T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a numeric value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := int(3)
//	Is(v.IntP(&quantity).GreaterThan(2))
func (validator *ValidatorIntP[T]) GreaterThan(value T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a numeric value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := int(3)
//	Is(v.IntP(&quantity).GreaterOrEqualTo(3))
func (validator *ValidatorIntP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a numeric value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.IntP(&quantity).LessThan(3))
func (validator *ValidatorIntP[T]) LessThan(value T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a numeric value is less than or equal to another. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := int(2)
//	Is(v.IntP(&quantity).LessOrEqualTo(2))
func (validator *ValidatorIntP[T]) LessOrEqualTo(value T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if a number is within a range (inclusive).
// For example:
//
//	n := int(3)
//	Is(v.IntP(&n).Between(2,6))
func (validator *ValidatorIntP[T]) Between(min T, max T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberBetween(*(validator.context.Value().(*T)), min, max)
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
//	n := int(0)
//	Is(v.IntP(&n).Zero())
func (validator *ValidatorIntP[T]) Zero(template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is positive (greater than zero).
//
// For example:
//
//	n := int(5)
//	Is(v.IntP(&n).Positive())
func (validator *ValidatorIntP[T]) Positive(template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && *(validator.context.Value().(*T)) > 0
		},
		ErrorKeyPositive, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is negative (less than zero).
//
// For example:
//
//	n := int(-5)
//	Is(v.IntP(&n).Negative())
func (validator *ValidatorIntP[T]) Negative(template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && *(validator.context.Value().(*T)) < 0
		},
		ErrorKeyNegative, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := int(2)
//	Is(v.IntP(&quantity).Passing((v *int) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *ValidatorIntP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a number is present in a numeric slice.
// For example:
//
//	quantity := int(3)
//	validQuantities := []int{1,3,5}
//	Is(v.IntP(&quantity).InSlice(validQuantities))
func (validator *ValidatorIntP[T]) InSlice(slice []T, template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is zero or nil.
//
// For example:
//
//	n := int(0)
//	Is(v.IntP(&n).ZeroOrNil())
//
//	var n *int
//	Is(v.IntP(n).ZeroOrNil())
func (validator *ValidatorIntP[T]) ZeroOrNil(template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil || isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric pointer value is nil.
//
// For example:
//
//	var n *int
//	Is(v.IntP(n).Nil())
func (validator *ValidatorIntP[T]) Nil(template ...string) *ValidatorIntP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}
