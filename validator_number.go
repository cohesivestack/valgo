package valgo

type TypeNumber interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~float32 |
		~float64
}

type ValidatorNumber[T TypeNumber] struct {
	context *ValidatorContext
}

func Number[T TypeNumber](value T, nameAndTitle ...string) *ValidatorNumber[T] {
	return &ValidatorNumber[T]{context: NewContext(value, nameAndTitle...)}
}

func (validator *ValidatorNumber[T]) Context() *ValidatorContext {
	return validator.context
}

func (validator *ValidatorNumber[T]) EqualTo(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) GreaterThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) > value
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) GreaterOrEqualThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) >= value
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) LessThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) < value
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) LessOrEqualThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) <= value
		},
		ErrorKeyLessOrEqualThan, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) Zero(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(T) == 0
		},
		ErrorKeyZero, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) Passing(function func(v T) bool, template ...string) *ValidatorNumber[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, template...)

	return validator
}

func (validator *ValidatorNumber[T]) InSlice(slice []T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			for _, v := range slice {
				if validator.context.Value() == v {
					return true
				}
			}
			return false
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

func (validator *ValidatorNumber[T]) Not() *ValidatorNumber[T] {
	validator.context.Not()

	return validator
}
