package valgo

// The [ValidatorUint] provides functions for setting validation rules for a
// uint value types, or a custom type based on a uint, uint8, uint16, uint32, or uint64.
type ValidatorUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64] struct {
	context *ValidatorContext
}

// Receives an uint value to validate.
//
// The value can also be a custom uint type such as type Age uint.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint[T ~uint](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint8 value to validate.
//
// The value can also be a custom uint8 type such as type Age uint8.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint8[T ~uint8](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint16 value to validate.
//
// The value can also be a custom uint16 type such as type Age uint16.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint16[T ~uint16](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint32 value to validate.
//
// The value can also be a custom uint32 type such as type Age uint32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint32[T ~uint32](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an uint64 value to validate.
//
// The value can also be a custom uint64 type such as type Age uint64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Uint64[T ~uint64](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Receives an byte value to validate.
//
// The value can also be a custom byte type such as type Age byte.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Byte[T ~byte](value T, nameAndTitle ...string) *ValidatorUint[T] {
	return &ValidatorUint[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorUint[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	Is(v.Uint(0).Not().Zero()).Valid()
func (validator *ValidatorUint[T]) Not() *ValidatorUint[T] {
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
// a single message composed by joining the failing fragments’ messages using
// the localized OR list format.
//
// Example:
//
//	// Passes because input is Zero (GreaterThan(5) OR Zero()).
//	input := uint(0)
//	isValid := v.Is(v.Uint(input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorUint[T]) Or() *ValidatorUint[T] {
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
// This is primarily used to express “accept X, otherwise validate the rest”
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
//	// If input is Zero, the chain succeeds and GreaterThan/LessThan are not evaluated.
//	// Otherwise, input must be within (5, 10).
//	input := uint(0)
//	isValid := v.Is(v.Uint(input).Zero().OrElse().GreaterThan(5).LessThan(10)).Valid()
func (validator *ValidatorUint[T]) OrElse() *ValidatorUint[T] {
	validator.context.OrElse()

	return validator
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := uint(2)
//	Is(v.Uint(quantity).EqualTo(2))
func (validator *ValidatorUint[T]) EqualTo(value T, template ...string) *ValidatorUint[T] {
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
//	quantity := uint(3)
//	Is(v.Uint(quantity).GreaterThan(2))
func (validator *ValidatorUint[T]) GreaterThan(value T, template ...string) *ValidatorUint[T] {
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
//	quantity := uint(3)
//	Is(v.Uint(quantity).GreaterOrEqualTo(3))
func (validator *ValidatorUint[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorUint[T] {
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
//	quantity := uint(2)
//	Is(v.Uint(quantity).LessThan(3))
func (validator *ValidatorUint[T]) LessThan(value T, template ...string) *ValidatorUint[T] {
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
//	quantity := uint(2)
//	Is(v.Uint(quantity).LessOrEqualTo(2))
func (validator *ValidatorUint[T]) LessOrEqualTo(value T, template ...string) *ValidatorUint[T] {
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
//	Is(v.Uint(uint(3)).Between(2,6))
func (validator *ValidatorUint[T]) Between(min T, max T, template ...string) *ValidatorUint[T] {
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
//	Is(v.Uint(uint(0)).Zero())
func (validator *ValidatorUint[T]) Zero(template ...string) *ValidatorUint[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberZero(validator.context.Value().(T))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := uint(2)
//	Is(v.Uint(quantity).Passing((v uint) bool {
//		return v == getAllowedQuantity()
//	})
func (validator *ValidatorUint[T]) Passing(function func(v T) bool, template ...string) *ValidatorUint[T] {
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
//	quantity := uint(3)
//	validQuantities := []uint{1,3,5}
//	Is(v.Uint(quantity).InSlice(validQuantities))
func (validator *ValidatorUint[T]) InSlice(slice []T, template ...string) *ValidatorUint[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
