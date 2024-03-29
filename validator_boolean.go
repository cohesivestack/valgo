package valgo

func isBoolTrue[T ~bool](v T) bool {
	return bool(v)
}

func isBoolFalse[T ~bool](v T) bool {
	return !bool(v)
}

func isBoolEqual[T ~bool](v0 T, v1 T) bool {
	return v0 == v1
}

func isBoolInSlice[T ~bool](v T, slice []T) bool {
	for _, _v := range slice {
		if v == _v {
			return true
		}
	}
	return false
}

// The Boolean validator type that keeps its validator context.
type ValidatorBool[T ~bool] struct {
	context *ValidatorContext
}

// Receives a boolean value to validate.
//
// The value also can be a custom boolean type such as `type Active bool;`
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0. When the name is
// provided but not the title, then the name is humanized to be used as the
// title as well; for example the name `phone_number` will be humanized as
// `Phone Number`

func Bool[T ~bool](value T, nameAndTitle ...string) *ValidatorBool[T] {
	return &ValidatorBool[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorBool[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because `Not()` inverts the boolean value associated with the True() function
//	Is(v.Bool(true).Not().True()).Valid()
func (validator *ValidatorBool[T]) Not() *ValidatorBool[T] {
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
//	isValid := v.Is(v.Bool(input).False().Or().True()).Valid()
func (validator *ValidatorBool[T]) Or() *ValidatorBool[T] {
	validator.context.Or()

	return validator
}

// Validate if a boolean value is equal to another.
// For example:
//
//	activated := true
//	Is(v.Bool(activated).Equal(true))
func (validator *ValidatorBool[T]) EqualTo(value T, template ...string) *ValidatorBool[T] {
	validator.context.AddWithValue(
		func() bool {
			return isBoolEqual(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a boolean value is true.
// For example:
//
//	activated := true
//	Is(v.Bool(activated).True())
func (validator *ValidatorBool[T]) True(template ...string) *ValidatorBool[T] {
	validator.context.Add(
		func() bool {
			return isBoolTrue(validator.context.Value().(T))
		},
		ErrorKeyTrue, template...)

	return validator
}

// Validate if a boolean value is false.
// For example:
//
//	activated := false
//	Is(v.Bool(activated).Equal(true)).Valid()
func (validator *ValidatorBool[T]) False(template ...string) *ValidatorBool[T] {
	validator.context.Add(
		func() bool {
			return isBoolFalse(validator.context.Value().(T))
		},
		ErrorKeyFalse, template...)

	return validator
}

// Validate if a boolean value pass a custom function.
// For example:
//
//	activated := false
//	Is(v.Bool(activated).Passing((v bool) bool {
//		return v == someBoolFunction()
//	})
func (validator *ValidatorBool[T]) Passing(function func(v T) bool, template ...string) *ValidatorBool[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, template...)

	return validator
}

// Validate if the value of a boolean pointer is present in a boolean slice.
// For example:
//
//	activated := false
//	elements := []bool{true, false, true}
//	Is(v.Bool(activated).InSlice(elements))
func (validator *ValidatorBool[T]) InSlice(slice []T, template ...string) *ValidatorBool[T] {
	validator.context.AddWithValue(
		func() bool {
			return isBoolInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
