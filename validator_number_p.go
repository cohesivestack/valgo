package valgo

//go:generate go run generator/main.go

type ValidatorNumberP[T TypeNumber] struct {
	context *ValidatorContext
}

func NumberP[T TypeNumber](value *T, nameAndTitle ...string) *ValidatorNumberP[T] {
	return &ValidatorNumberP[T]{context: NewContext(value, nameAndTitle...)}
}

func (validator *ValidatorNumberP[T]) Context() *ValidatorContext {
	return validator.context
}

func (validator *ValidatorNumberP[T]) EqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) GreaterThan(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberGreaterOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) LessThan(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessThan(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) LessOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberLessOrEqualTo(*(validator.context.Value().(*T)), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if the value of a number is in a range (inclusive).
// For example:
//
//	Is(v.Number(3).Between(2,6))
func (validator *ValidatorNumberP[T]) Between(min T, max T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberBetween(*(validator.context.Value().(*T)), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

func (validator *ValidatorNumberP[T]) Zero(template ...string) *ValidatorNumberP[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) ZeroOrNil(template ...string) *ValidatorNumberP[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil || isNumberZero(*(validator.context.Value().(*T)))
		},
		ErrorKeyZero, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) Nil(template ...string) *ValidatorNumberP[T] {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(*T) == nil
		},
		ErrorKeyNil, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorNumberP[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(*T))
		},
		ErrorKeyPassing, template...)

	return validator
}

func (validator *ValidatorNumberP[T]) InSlice(slice []T, template ...string) *ValidatorNumberP[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*T) != nil && isNumberInSlice(*(validator.context.Value().(*T)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

func (validator *ValidatorNumberP[T]) Not() *ValidatorNumberP[T] {
	validator.context.Not()

	return validator
}
