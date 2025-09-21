package valgo

import "math"

// The [ValidatorFloatP] provides functions for setting validation rules for a
// float pointer value types, or a custom type based on a float32 or float64 pointer.
type ValidatorFloatP[T ~float32 | ~float64] struct {
	context *ValidatorContext
}

// Receives a float32 pointer value to validate.
//
// The value can also be a custom float32 pointer type such as type Price *float32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Float32P[T ~float32](value *T, nameAndTitle ...string) *ValidatorFloatP[T] {
	return &ValidatorFloatP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives a float64 pointer value to validate.
//
// The value can also be a custom float64 pointer type such as type Price *float64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Float64P[T ~float64](value *T, nameAndTitle ...string) *ValidatorFloatP[T] {
	return &ValidatorFloatP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorFloatP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	n := float32(0)
//	Is(v.Float32P(&n).Not().Zero()).Valid()
func (validator *ValidatorFloatP[T]) Not() *ValidatorFloatP[T] {
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
//	input := float32(0)
//	isValid := v.Is(v.Float32P(&input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorFloatP[T]) Or() *ValidatorFloatP[T] {
	validator.context.Or()

	return validator
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32P(&quantity).EqualTo(2))
func (validator *ValidatorFloatP[T]) EqualTo(value T, template ...string) *ValidatorFloatP[T] {
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
//	quantity := float32(3)
//	Is(v.Float32P(&quantity).GreaterThan(2))
func (validator *ValidatorFloatP[T]) GreaterThan(value T, template ...string) *ValidatorFloatP[T] {
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
//	quantity := float32(3)
//	Is(v.Float32P(&quantity).GreaterOrEqualTo(3))
func (validator *ValidatorFloatP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorFloatP[T] {
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
//	quantity := float32(2)
//	Is(v.Float32P(&quantity).LessThan(3))
func (validator *ValidatorFloatP[T]) LessThan(value T, template ...string) *ValidatorFloatP[T] {
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
//	quantity := float32(2)
//	Is(v.Float32P(&quantity).LessOrEqualTo(2))
func (validator *ValidatorFloatP[T]) LessOrEqualTo(value T, template ...string) *ValidatorFloatP[T] {
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
//	n := float32(3)
//	Is(v.Float32P(&n).Between(2,6))
func (validator *ValidatorFloatP[T]) Between(min T, max T, template ...string) *ValidatorFloatP[T] {
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
//	n := float32(0)
//	Is(v.Float32P(&n).Zero())
func (validator *ValidatorFloatP[T]) Zero(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32P(&quantity).Passing((v *float32) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *ValidatorFloatP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorFloatP[T] {
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
//	quantity := float32(3)
//	validQuantities := []float32{1,3,5}
//	Is(v.Float32P(&quantity).InSlice(validQuantities))
func (validator *ValidatorFloatP[T]) InSlice(slice []T, template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric pointer value is nil.
//
// For example:
//
//	var n *float32
//	Is(v.Float32P(n).Nil())
func (validator *ValidatorFloatP[T]) Nil(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is zero or nil.
//
// For example:
//
//	n := float32(0)
//	Is(v.Float32P(&n).ZeroOrNil())
//
//	var n *float32
//	Is(v.Float32P(n).ZeroOrNil())
func (validator *ValidatorFloatP[T]) ZeroOrNil(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil || isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is positive (greater than zero).
//
// For example:
//
//	n := float32(5.5)
//	Is(v.Float32P(&n).Positive())
func (validator *ValidatorFloatP[T]) Positive(template ...string) *ValidatorFloatP[T] {
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
//	n := float32(-5.5)
//	Is(v.Float32P(&n).Negative())
func (validator *ValidatorFloatP[T]) Negative(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && *(validator.context.Value().(*T)) < 0
		},
		ErrorKeyNegative, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is NaN (Not a Number).
//
// For example:
//
//	n := float32(math.NaN())
//	Is(v.Float32P(&n).NaN())
func (validator *ValidatorFloatP[T]) NaN(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && math.IsNaN(float64(*(validator.context.Value().(*T))))
		},
		ErrorKeyNaN, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is infinite (positive or negative infinity).
//
// For example:
//
//	n := float32(math.Inf(1))
//	Is(v.Float32P(&n).Infinite())
func (validator *ValidatorFloatP[T]) Infinite(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && math.IsInf(float64(*(validator.context.Value().(*T))), 0)
		},
		ErrorKeyInfinite, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value is finite (not NaN and not infinite).
//
// For example:
//
//	n := float32(3.14)
//	Is(v.Float32P(&n).Finite())
func (validator *ValidatorFloatP[T]) Finite(template ...string) *ValidatorFloatP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && !math.IsNaN(float64(*(validator.context.Value().(*T)))) && !math.IsInf(float64(*(validator.context.Value().(*T))), 0)
		},
		ErrorKeyFinite, validator.context.Value(), template...)

	return validator
}
