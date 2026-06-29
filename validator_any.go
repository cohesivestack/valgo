package valgo

import (
	"reflect"
)

// The Any validator's type that keeps its validator context.
type ValidatorAny struct {
	context *ValidatorContext
}

// Receive a value to validate.
//
// The value can be any type;
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N` pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.
func Any(value any, nameAndTitle ...string) *ValidatorAny {
	return &ValidatorAny{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorAny) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because `Not()` inverts the boolean value associated with the `Equal()` function
//	Is(v.Any("a").Not().Equal("a")).Valid()
func (validator *ValidatorAny) Not() *ValidatorAny {
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
//	isValid := v.Is(v.Any(input).Passing(func(v any) bool { return v == "other" }).Or().Passing(func(v any) bool { return v == "test" })).Valid()
func (validator *ValidatorAny) Or() *ValidatorAny {
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
//	isValid := v.Is(v.Any(input).Nil().OrElse().Passing(func(v any) bool { return v != nil && *v.(*string) == "test" })).Valid()
func (validator *ValidatorAny) OrElse() *ValidatorAny {
	validator.context.OrElse()

	return validator
}

// Validate if a value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.Any(status).Equal("running"))
//
// DEPRECATED: 'any' is not safely comparable. Use the Comparable validator instead.
// This function will be removed in Valgo v1.0.0.
func (validator *ValidatorAny) EqualTo(value any, template ...string) *ValidatorAny {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value() == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a value passes a custom function.
// For example:
//
//	status := ""
//	Is(v.Any(status).Passing((v any) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorAny) Passing(function func(v any) bool, template ...string) *ValidatorAny {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value())
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a value is nil.
// For example:
//
//	var status *string
//	Is(v.Any(status).Nil())
func (validator *ValidatorAny) Nil(template ...string) *ValidatorAny {
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
