package valgo

//go:generate go run generator/main.go

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

func isNumberEqualTo[T TypeNumber](v0 T, v1 T) bool {
	return v0 == v1
}

func isNumberGreaterThan[T TypeNumber](v0 T, v1 T) bool {
	return v0 > v1
}
func isNumberGreaterOrEqualTo[T TypeNumber](v0 T, v1 T) bool {
	return v0 >= v1
}
func isNumberLessThan[T TypeNumber](v0 T, v1 T) bool {
	return v0 < v1
}
func isNumberLessOrEqualTo[T TypeNumber](v0 T, v1 T) bool {
	return v0 <= v1
}
func isNumberBetween[T TypeNumber](v T, min T, max T) bool {
	return v >= min && v <= max
}
func isNumberZero[T TypeNumber](v T) bool {
	return v == 0
}
func isNumberInSlice[T TypeNumber](v T, slice []T) bool {
	for _, _v := range slice {
		if v == _v {
			return true
		}
	}
	return false
}

type ValidatorNumber[T TypeNumber] struct {
	context *ValidatorContext
}

// Receives a number value to validate.
//
// The value also can be any golang number type (int64, int32, float32, uint,
// etc.) or a custom number type such as `type Level int32;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`

func Number[T TypeNumber](value T, nameAndTitle ...string) *ValidatorNumber[T] {
	return &ValidatorNumber[T]{context: NewContext(value, nameAndTitle...)}
}

// This function returns the context for the Valgo Validator session's
// validator. The function should not be called unless you are creating a custom
// validator by extending this validator.
func (validator *ValidatorNumber[T]) Context() *ValidatorContext {
	return validator.context
}

// Reverse the logical value associated to the next validation function.
// For example:
//
//	// It will return false because Not() inverts to Zero()
//	Is(v.Number(0).Not().Zero()).Valid()
func (validator *ValidatorNumber[T]) EqualTo(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) GreaterThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberGreaterThan(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberGreaterOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) LessThan(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberLessThan(validator.context.Value().(T), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

func (validator *ValidatorNumber[T]) LessOrEqualTo(value T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithValue(
		func() bool {
			return isNumberLessOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if the value of a number is in a range (inclusive).
// For example:
//
//	Is(v.Number(3).Between(2,6))
func (validator *ValidatorNumber[T]) Between(min T, max T, template ...string) *ValidatorNumber[T] {
	validator.context.AddWithParams(
		func() bool {
			return isNumberBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

func (validator *ValidatorNumber[T]) Zero(template ...string) *ValidatorNumber[T] {
	validator.context.Add(
		func() bool {
			return isNumberZero(validator.context.Value().(T))
		},
		ErrorKeyZero, template...)

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
			return isNumberInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

func (validator *ValidatorNumber[T]) Not() *ValidatorNumber[T] {
	validator.context.Not()

	return validator
}
