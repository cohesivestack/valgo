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

// Validate if a value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.Any(status).Equal("running"))
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
	validator.context.Add(
		func() bool {
			return function(validator.context.Value())
		},
		ErrorKeyPassing, template...)

	return validator
}

// Validate if a value is nil.
// For example:
//
//	var status *string
//	Is(v.Any(status).Nil())
func (validator *ValidatorAny) Nil(template ...string) *ValidatorAny {
	validator.context.Add(
		func() bool {
			val := validator.context.Value()
			// In Golang nil sometimes is not equal to raw nil, such as it's explained
			// here: https://dev.to/arxeiss/in-go-nil-is-not-equal-to-nil-sometimes-jn8
			// So, seems using reflection is the only option here
			return val == nil ||
				(reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil())
		},
		ErrorKeyNil, template...)

	return validator
}
