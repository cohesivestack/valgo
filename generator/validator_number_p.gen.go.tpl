// Code generated by Valgo; DO NOT EDIT.
package valgo
{{ range . }}
// The {{ .Type }} pointer validator type that keeps its validator context.
type Validator{{ .Name }}P[T ~{{ .Type }}] struct {
	context *ValidatorContext
}

// Receives the {{ .Type }} pointer to validate.
//
// The value also can be a custom {{ .Type }} type such as `type Level *{{ .Type }};`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func {{ .Name }}P[T ~{{ .Type }}](value *T, nameAndTitle ...string) *Validator{{ .Name }}P[T] {
	return &Validator{{ .Name }}P[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *Validator{{ .Name }}P[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the boolean value associated with the next validator function.
// For example:
//
//	// It will return false because Not() inverts the boolean value associated with the Zero() function
//	n := {{ .Type }}(0)
//	Is(v.{{ .Name }}P(&n).Not().Zero()).Valid()
func (validator *Validator{{ .Name }}P[T]) Not() *Validator{{ .Name }}P[T] {
	validator.context.Not()

	return validator
}

// Validate if the {{ .Type }} pointer value is equal to another value. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}P(quantity).Equal({{ .Type }}(2)))
func (validator *Validator{{ .Name }}P[T]) EqualTo(value T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}EqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is greater than another value. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := {{ .Type }}(3)
//	Is(v.{{ .Name }}P(&quantity).GreaterThan({{ .Type }}(2)))
func (validator *Validator{{ .Name }}P[T]) GreaterThan(value T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}GreaterThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is greater than or equal to another value. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := {{ .Type }}(3)
//	Is(v.{{ .Name }}P(&quantity).GreaterOrEqualTo({{ .Type }}(3)))
func (validator *Validator{{ .Name }}P[T]) GreaterOrEqualTo(value T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}GreaterOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is less than another value. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}P(&quantity).LessThan({{ .Type }}(3)))
func (validator *Validator{{ .Name }}P[T]) LessThan(value T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}LessThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is less than or equal to another value. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}P(&quantity).LessOrEqualTo({{ .Type }}(2)))
func (validator *Validator{{ .Name }}P[T]) LessOrEqualTo(value T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}LessOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if the value of the {{ .Type }} pointer is within a range (inclusive).
// For example:
//
//	n := {{ .Type }}(3)
//	Is(v.{{ .Name }}P(&n).Between({{ .Type }}(2),{{ .Type }}(6)))
func (validator *Validator{{ .Name }}P[T]) Between(min T, max T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}Between(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is zero.
//
// For example:
//
//	n := {{ .Type }}(0)
//	Is(v.{{ .Name }}P(&n).Zero())
func (validator *Validator{{ .Name }}P[T]) Zero(template ...string) *Validator{{ .Name }}P[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}Zero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is zero or nil.
//
// For example:
//
//	var _quantity *{{ .Type }}
//	Is(v.{{ .Name }}P(_quantity).ZeroOrNil()) // Will be true
func (validator *Validator{{ .Name }}P[T]) ZeroOrNil(template ...string) *Validator{{ .Name }}P[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil || is{{ .Name }}Zero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is nil.
//
// For example:
//
//	var quantity *{{ .Type }}
//	Is(v.{{ .Name }}P(quantity).Nil()) // Will be true
func (validator *Validator{{ .Name }}P[T]) Nil(template ...string) *Validator{{ .Name }}P[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value passes a custom function.
// For example:
//
//	quantity := {{ .Type }}(2)
//	Is(v.{{ .Name }}P(&quantity).Passing((v *{{ .Type }}) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *Validator{{ .Name }}P[T]) Passing(function func(v *T) bool, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, template...)

	return validator
}

// Validate if the {{ .Type }} pointer value is present in a numeric slice.
// For example:
//
//	quantity := {{ .Type }}(3)
//	validQuantities := []{{ .Type }}{1,3,5}
//	Is(v.{{ .Name }}P(&quantity).InSlice(validQuantities))
func (validator *Validator{{ .Name }}P[T]) InSlice(slice []T, template ...string) *Validator{{ .Name }}P[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && is{{ .Name }}InSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
{{ end }}