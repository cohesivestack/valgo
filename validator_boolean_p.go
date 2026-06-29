package valgo

// The Boolean pointer validator type that keeps its validator context.
type ValidatorBoolP[T ~bool] struct {
	context *ValidatorContext
}

// Receives a boolean pointer to validate.
//
// The value also can be a custom boolean type such as `type Active bool;`
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0. When the name is
// provided but not the title, then the name is humanized to be used as the
// title as well; for example the name `phone_number` will be humanized as
// `Phone Number`
func BoolP[T ~bool](value *T, nameAndTitle ...string) *ValidatorBoolP[T] {
	return &ValidatorBoolP[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorBoolP[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the True() function
//	activated := true
//	Is(v.BoolP(&activated).Not().True()).Valid()
func (validator *ValidatorBoolP[T]) Not() *ValidatorBoolP[T] {
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
//	// Passes because input is True (False() OR True()).
//	input := true
//	isValid := v.Is(v.BoolP(&input).False().Or().True()).Valid()
func (validator *ValidatorBoolP[T]) Or() *ValidatorBoolP[T] {
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
//	// If input is True, the chain succeeds and False() is not evaluated.
//	// Otherwise, input must be False.
//	input := true
//	isValid := v.Is(v.BoolP(&input).True().OrElse().False()).Valid()
func (validator *ValidatorBoolP[T]) OrElse() *ValidatorBoolP[T] {
	validator.context.OrElse()

	return validator
}

// Validate if the value of a boolean pointer is equal to another value.
// For example:
//
//	activated := true
//	Is(v.BoolP(&activated).Equal(true))
func (validator *ValidatorBoolP[T]) EqualTo(value T, template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolEqual(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if the value of a boolean pointer is true.
// For example:
//
//	activated := true
//	Is(v.BoolP(&activated).True())
func (validator *ValidatorBoolP[T]) True(template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolTrue(*(validator.context.Value().(*T)))
		},
		ErrorKeyTrue, validator.context.Value(), template...)

	return validator
}

// Validate if the value of a boolean pointer is false.
// For example:
//
//	activated := false
//	Is(v.BoolP(&activated).False())
func (validator *ValidatorBoolP[T]) False(template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolFalse(*(validator.context.Value().(*T)))
		},
		ErrorKeyFalse, validator.context.Value(), template...)

	return validator
}

// Validate if the value of a boolean pointer is false or nil.
// For example:
//
//	var activated *bool
//	Is(v.BoolP(activated).FalseOrNil())
//	*activated = false
//	Is(v.BoolP(activated).FalseOrNil())
func (validator *ValidatorBoolP[T]) FalseOrNil(template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil || isBoolFalse(*(validator.context.Value().(*T)))
		},
		ErrorKeyFalse, validator.context.Value(), template...)

	return validator
}

// Validate if a boolean pointer is nil.
// For example:
//
//	var activated *bool
//	Is(v.BoolP(activated).Nil())
func (validator *ValidatorBoolP[T]) Nil(template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}

// Validate if a boolean pointer pass a custom function.
// For example:
//
//	activated := false
//	Is(v.BoolP(&activated).Passing((v *bool) bool {
//		return *v == someBoolFunction()
//	})
func (validator *ValidatorBoolP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if the value of a boolean pointer is present in a boolean slice.
// For example:
//
//	activated := false
//	elements := []bool{true, false, true}
//	Is(v.BoolP(&activated).InSlice(elements))
func (validator *ValidatorBoolP[T]) InSlice(slice []T, template ...string) *ValidatorBoolP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
