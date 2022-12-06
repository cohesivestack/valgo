// Code generated by Valgo; DO NOT EDIT.
package valgo
{{ range . }}

// The {{ .Type }} validator type that keeps its validator context.
type Validator{{ .Name }}[T ~{{ .Type }}] struct {
	context *ValidatorContext
}

// Receives the {{ .Type }} value to validate.
//
// The value also can be a custom {{ .Type }} type such as `type Level {{ .Type }};`
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone {{ .Name }}`

func {{ .Name }}[T ~{{ .Type }}](value T, nameAndTitle ...string) *Validator{{ .Name }}[T] {
	return &Validator{{ .Name }}[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *Validator{{ .Name }}[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	Is(v.{{ .Name }}({{ .Type }}(0)).Not().Zero()).Valid()
func (validator *Validator{{ .Name }}[T]) Not() *Validator{{ .Name }}[T] {
	validator.context.Not()

	return validator
}

// Validate if the {{ .Type }} value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}(quantity).Equal({{ .Type }}(2)))
func (validator *Validator{{ .Name }}[T]) EqualTo(value T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}EqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if the {{ .Type }} value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := {{ .Type }}(3)
//	Is(v.{{ .Name }}(quantity).GreaterThan({{ .Type }}(2)))
func (validator *Validator{{ .Name }}[T]) GreaterThan(value T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}GreaterThan(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if the {{ .Type }} value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := {{ .Type }}(3)
//	Is(v.{{ .Name }}(quantity).GreaterOrEqualTo({{ .Type }}(3)))
func (validator *Validator{{ .Name }}[T]) GreaterOrEqualTo(value T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}GreaterOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if the {{ .Type }} value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}(quantity).LessThan({{ .Type }}(3)))
func (validator *Validator{{ .Name }}[T]) LessThan(value T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}LessThan(validator.context.Value().(T), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if the {{ .Type }} value is less than or equal to another. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}(quantity).LessOrEqualTo({{ .Type }}(2)))
func (validator *Validator{{ .Name }}[T]) LessOrEqualTo(value T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}LessOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if the {{ .Type }} is within a range (inclusive).
// For example:
//
//	Is(v.{{ .Name }}({{ .Type }}(3)).Between({{ .Type }}(2),{{ .Type }}(6)))
func (validator *Validator{{ .Name }}[T]) Between(min T, max T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithParams(
		func() bool {
			return is{{ .Name }}Between(validator.context.Value().(T), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Validate if the {{ .Type }} value is zero.
//
// For example:
//
//	Is(v.{{ .Name }}({{ .Type }}(0)).Zero())
func (validator *Validator{{ .Name }}[T]) Zero(template ...string) *Validator{{ .Name }}[T] {
	validator.context.Add(
		func() bool {
			return is{{ .Name }}Zero(validator.context.Value().(T))
		},
		ErrorKeyZero, template...)

	return validator
}

// Validate if the {{ .Type }} value passes a custom function.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}(quantity).Passing((v {{ .Type }}) bool {
//		return v == getAllowedQuantity()
//	})
func (validator *Validator{{ .Name }}[T]) Passing(function func(v T) bool, template ...string) *Validator{{ .Name }}[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, template...)

	return validator
}

// Validate if the {{ .Type }} value is present in the {{ .Type }} slice.
// For example:
//
//	quantity := {{ .Type }}(3)
//	validQuantities := []{{ .Type }}{1,3,5}
//	Is(v.{{ .Name }}(quantity).InSlice(validQuantities))
func (validator *Validator{{ .Name }}[T]) InSlice(slice []T, template ...string) *Validator{{ .Name }}[T] {
	validator.context.AddWithValue(
		func() bool {
			return is{{ .Name }}InSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

func is{{ .Name }}EqualTo[T ~{{ .Type }}](v0 T, v1 T) bool {
	return v0 == v1
}

func is{{ .Name }}GreaterThan[T ~{{ .Type }}](v0 T, v1 T) bool {
	return v0 > v1
}
func is{{ .Name }}GreaterOrEqualTo[T ~{{ .Type }}](v0 T, v1 T) bool {
	return v0 >= v1
}
func is{{ .Name }}LessThan[T ~{{ .Type }}](v0 T, v1 T) bool {
	return v0 < v1
}
func is{{ .Name }}LessOrEqualTo[T ~{{ .Type }}](v0 T, v1 T) bool {
	return v0 <= v1
}
func is{{ .Name }}Between[T ~{{ .Type }}](v T, min T, max T) bool {
	return v >= min && v <= max
}
func is{{ .Name }}Zero[T ~{{ .Type }}](v T) bool {
	return v == 0
}
func is{{ .Name }}InSlice[T ~{{ .Type }}](v T, slice []T) bool {
	for _, _v := range slice {
		if v == _v {
			return true
		}
	}
	return false
}
{{ end }}