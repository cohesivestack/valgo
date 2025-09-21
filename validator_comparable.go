package valgo

// The Comparable validator's type that keeps its validator context.
// T can be any Go type (pointer, struct, etc.) that is comparable.
type ValidatorComparable[T comparable] struct {
	context *ValidatorContext
}

// Receive a value of type T to validate, where T must be comparable.
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.
//
// Example:
//
//	v.Is(v.Comparable(user).Not().Nil())
func Comparable[T comparable](value T, nameAndTitle ...string) *ValidatorComparable[T] {
	return &ValidatorComparable[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorComparable[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because `Not()` inverts the boolean value associated with `EqualTo()`
//	v.Is(v.Comparable("a").Not().EqualTo("a")).Valid()
func (validator *ValidatorComparable[T]) Not() *ValidatorComparable[T] {
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
//	// This validator will pass because the string is equals "running".
//	status := "running"
//	isValid := v.Is(v.Comparable(status).EqualTo("paused").Or().EqualTo("running")).Valid()
func (validator *ValidatorComparable[T]) Or() *ValidatorComparable[T] {
	validator.context.Or()
	return validator
}

// Validate if a value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.Comparable(status).EqualTo("running"))
func (validator *ValidatorComparable[T]) EqualTo(value T, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value() == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a value passes a custom function.
// The function receives a typed T value, enabling compile-time type safety.
//
// For example:
//
//	type Status string
//	status := Status("running")
//	isValid := v.Is(
//	  v.Comparable(status).Passing(func(s Status) bool {
//	    return s == "running" || s == "paused"
//	  }),
//	).Valid()
func (validator *ValidatorComparable[T]) Passing(function func(v T) bool, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			// Value is stored as interface{} inside the context; assert back to T.
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a value is present in a slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.Comparable(status).InSlice(validStatus))
func (validator *ValidatorComparable[T]) InSlice(slice []T, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			v := validator.context.Value().(T)
			for _, s := range slice {
				if v == s {
					return true
				}
			}
			return false
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
