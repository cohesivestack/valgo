package valgo

import "reflect"

// The Typed validator's type that keeps its validator context.
// T can be any Go type (pointer, struct, slice, map, etc.).
type ValidatorTyped[T any] struct {
	context *ValidatorContext
}

// Receive a value of type T to validate.
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
//	v.Is(v.Typed(user).Not().Nil())
func Typed[T any](value T, nameAndTitle ...string) *ValidatorTyped[T] {
	return &ValidatorTyped[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorTyped[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because `Not()` inverts the boolean value associated with `EqualTo()`
//	v.Is(v.Typed("a").Not().EqualTo("a")).Valid()
func (validator *ValidatorTyped[T]) Not() *ValidatorTyped[T] {
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
//	// Passes because input passes the second Passing() check.
//	input := "test"
//	isValid := v.Is(v.Typed(input).Passing(func(s string) bool { return len(s) > 5 }).Or().Passing(func(s string) bool { return s == "test" })).Valid()
func (validator *ValidatorTyped[T]) Or() *ValidatorTyped[T] {
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
//	// If input is nil, the chain succeeds and Passing() is not evaluated.
//	// Otherwise, input must pass the custom function.
//	var input *string
//	isValid := v.Is(v.Typed(input).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" })).Valid()
func (validator *ValidatorTyped[T]) OrElse() *ValidatorTyped[T] {
	validator.context.OrElse()
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
//	  v.Typed(status).Passing(func(s Status) bool {
//	    return s == "running" || s == "paused"
//	  }),
//	).Valid()
func (validator *ValidatorTyped[T]) Passing(function func(v T) bool, template ...string) *ValidatorTyped[T] {
	validator.context.AddWithValue(
		func() bool {
			// Value is stored as interface{} inside the context; assert back to T.
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a value is nil.
// Works for nil-able kinds: pointers, slices, maps, chans, funcs, and interfaces.
// For non-nil-able types, this will return false.
//
// For example:
//
//	var s *string
//	v.Is(v.Typed(s).Nil())
func (validator *ValidatorTyped[T]) Nil(template ...string) *ValidatorTyped[T] {
	validator.context.AddWithValue(
		func() bool {
			val := validator.context.Value()
			// In Golang nil sometimes is not equal to raw nil, such as it's explained
			// here: https://dev.to/arxeiss/in-go-nil-is-not-equal-to-nil-sometimes-jn8
			// So, seems using reflection is the only option here
			return val == nil ||
				(reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil())
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}
