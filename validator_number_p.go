package valgo

//go:generate go run generator/main.go

// The Numeric pointer validator type that keeps its validator context.
type ValidatorNumberP[T TypeNumber] struct {
	context *ValidatorContext
}

// Receives a numeric pointer to validate.
//
// The value can be any golang numeric pointer type (*int64, *int32, *float32, *uint,
// etc.) or a custom numeric type such as `type Level *int32;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func NumberP[T TypeNumber](value *T, nameAndTitle ...string) *ValidatorNumberP[T] {
	return &ValidatorNumberP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorNumberP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	n := 0
//	Is(v.NumberP(&n).Not().Zero()).Valid()
func (validator *ValidatorNumberP[T]) Not() *ValidatorNumberP[T] {
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
//	// Passes because input is Zero (GreaterThan(5) OR Zero()).
//	input := 0
//	isValid := v.Is(v.NumberP(&input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorNumberP[T]) Or() *ValidatorNumberP[T] {
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
//	// If input is Zero, the chain succeeds and GreaterThan/LessThan are not evaluated.
//	// Otherwise, input must be within (5, 10).
//	input := 0
//	isValid := v.Is(v.NumberP(&input).Zero().OrElse().GreaterThan(5).LessThan(10)).Valid()
func (validator *ValidatorNumberP[T]) OrElse() *ValidatorNumberP[T] {
	validator.context.OrElse()

	return validator
}

// Validate if a numeric pointer value is equal to another value. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(quantity).Equal(2))
func (validator *ValidatorNumberP[T]) EqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a numeric pointer value is greater than another value. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := 3
//	Is(v.NumberP(&quantity).GreaterThan(2))
func (validator *ValidatorNumberP[T]) GreaterThan(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a numeric pointer value is greater than or equal to another value. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := 3
//	Is(v.NumberP(&quantity).GreaterOrEqualTo(3))
func (validator *ValidatorNumberP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a numeric pointer value is less than another value. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).LessThan(3))
func (validator *ValidatorNumberP[T]) LessThan(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a numeric pointer value is less than or equal to another value. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).LessOrEqualTo(2))
func (validator *ValidatorNumberP[T]) LessOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if the value of a numeric pointer is within a range (inclusive).
// For example:
//
//	n := 3
//	Is(v.NumberP(&n).Between(2,6))
func (validator *ValidatorNumberP[T]) Between(min T, max T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberBetween(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max, "value": validator.context.Value()},
		template...)

	return validator
}

// Validate if a numeric pointer value is zero.
//
// For example:
//
//	n := 0
//	Is(v.NumberP(&n).Zero())
func (validator *ValidatorNumberP[T]) Zero(template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric pointer value is zero or nil.
//
// For example:
//
//	var _quantity *int
//	Is(v.NumberP(_quantity).ZeroOrNil()) // Will be true
func (validator *ValidatorNumberP[T]) ZeroOrNil(template ...string) *ValidatorNumberP[T] {
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
//	var quantity *int
//	Is(v.NumberP(quantity).Nil()) // Will be true
func (validator *ValidatorNumberP[T]) Nil(template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric pointer value passes a custom function.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).Passing((v *int) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *ValidatorNumberP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a numeric pointer value is present in a numeric slice.
// For example:
//
//	quantity := 3
//	validQuantities := []int{1,3,5}
//	Is(v.NumberP(&quantity).InSlice(validQuantities))
func (validator *ValidatorNumberP[T]) InSlice(slice []T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
