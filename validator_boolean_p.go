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

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the input is equals false.
//	input := true
//	isValid := v.Is(v.BoolP(&input).Nil().Or().EqualTo(false)).Valid()
func (validator *ValidatorBoolP[T]) Or() *ValidatorBoolP[T] {
	validator.context.Or()

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
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolTrue(*(validator.context.Value().(*T)))
		},
		ErrorKeyTrue, template...)

	return validator
}

// Validate if the value of a boolean pointer is false.
// For example:
//
//	activated := false
//	Is(v.BoolP(&activated).False())
func (validator *ValidatorBoolP[T]) False(template ...string) *ValidatorBoolP[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) != nil && isBoolFalse(*(validator.context.Value().(*T)))
		},
		ErrorKeyFalse, template...)

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
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil || isBoolFalse(*(validator.context.Value().(*T)))
		},
		ErrorKeyFalse, template...)

	return validator
}

// Validate if a boolean pointer is nil.
// For example:
//
//	var activated *bool
//	Is(v.BoolP(activated).Nil())
func (validator *ValidatorBoolP[T]) Nil(template ...string) *ValidatorBoolP[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, template...)

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
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, template...)

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
