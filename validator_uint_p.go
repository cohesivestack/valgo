package valgo

// The [ValidatorUintP] provides functions for setting validation rules for a
// uint pointer value types, or a custom type based on a uint, uint8, uint16, uint32, or uint64 pointer.
type ValidatorUintP[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64] struct {
	context *ValidatorContext
}

// Receives an uint pointer value to validate.
//
// The value can also be a custom uint pointer type such as type Age *uint.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func UintP[T ~uint](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint8 pointer value to validate.
//
// The value can also be a custom uint8 pointer type such as type Age *uint8.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint8P[T ~uint8](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint16 pointer value to validate.
//
// The value can also be a custom uint16 pointer type such as type Age *uint16.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint16P[T ~uint16](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint32 pointer value to validate.
//
// The value can also be a custom uint32 pointer type such as type Age *uint32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint32P[T ~uint32](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint64 pointer value to validate.
//
// The value can also be a custom uint64 pointer type such as type Age *uint64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint64P[T ~uint64](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives a byte pointer value to validate.
//
// The value can also be a custom byte pointer type such as type Age *byte.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func ByteP[T ~byte](value *T, nameAndTitle ...string) *ValidatorUintP[T] {
	return &ValidatorUintP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorUintP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	n := uint(0)
//	Is(v.UintP(&n).Not().Zero()).Valid()
func (validator *ValidatorUintP[T]) Not() *ValidatorUintP[T] {
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
//	input := uint(0)
//	isValid := v.Is(v.UintP(&input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorUintP[T]) Or() *ValidatorUintP[T] {
	validator.context.Or()

	return validator
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := uint(2)
//	Is(v.UintP(&quantity).EqualTo(2))
func (validator *ValidatorUintP[T]) EqualTo(value T, template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(3)
//	Is(v.UintP(&quantity).GreaterThan(2))
func (validator *ValidatorUintP[T]) GreaterThan(value T, template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(3)
//	Is(v.UintP(&quantity).GreaterOrEqualTo(3))
func (validator *ValidatorUintP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(2)
//	Is(v.UintP(&quantity).LessThan(3))
func (validator *ValidatorUintP[T]) LessThan(value T, template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(2)
//	Is(v.UintP(&quantity).LessOrEqualTo(2))
func (validator *ValidatorUintP[T]) LessOrEqualTo(value T, template ...string) *ValidatorUintP[T] {
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
//	n := uint(3)
//	Is(v.UintP(&n).Between(2,6))
func (validator *ValidatorUintP[T]) Between(min T, max T, template ...string) *ValidatorUintP[T] {
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
//	n := uint(0)
//	Is(v.UintP(&n).Zero())
func (validator *ValidatorUintP[T]) Zero(template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(2)
//	Is(v.UintP(&quantity).Passing((v *uint) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *ValidatorUintP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorUintP[T] {
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
//	quantity := uint(3)
//	validQuantities := []uint{1,3,5}
//	Is(v.UintP(&quantity).InSlice(validQuantities))
func (validator *ValidatorUintP[T]) InSlice(slice []T, template ...string) *ValidatorUintP[T] {
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
//	n := uint(0)
//	Is(v.UintP(&n).ZeroOrNil())
//
//	var n *uint
//	Is(v.UintP(n).ZeroOrNil())
func (validator *ValidatorUintP[T]) ZeroOrNil(template ...string) *ValidatorUintP[T] {
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
//	var n *uint
//	Is(v.UintP(n).Nil())
func (validator *ValidatorUintP[T]) Nil(template ...string) *ValidatorUintP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}
