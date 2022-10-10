package valgo

import (
	"reflect"
)

type ValidatorAny struct {
	context *ValidatorContext
}

func Any(value any, nameAndTitle ...string) *ValidatorAny {
	return &ValidatorAny{context: NewContext(value, nameAndTitle...)}
}

func (validator *ValidatorAny) Context() *ValidatorContext {
	return validator.context
}

func (validator *ValidatorAny) Not() *ValidatorAny {
	validator.context.Not()

	return validator
}

func (validator *ValidatorAny) EqualTo(value any, template ...string) *ValidatorAny {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value() == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorAny) Passing(function func(v any) bool, template ...string) *ValidatorAny {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value())
		},
		ErrorKeyPassing, template...)

	return validator
}

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
