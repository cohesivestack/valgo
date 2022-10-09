package valgo

type ValidatorComplex128 struct {
	context *ValidatorContext
}

func Complex128(value complex128, nameAndTitle ...string) *ValidatorComplex128 {
	return &ValidatorComplex128{context: NewContext(value, nameAndTitle...)}
}

func (validator *ValidatorComplex128) Context() *ValidatorContext {
	return validator.context
}

func (validator *ValidatorComplex128) EqualTo(value complex128, template ...string) *ValidatorComplex128 {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(complex128) == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

func (validator *ValidatorComplex128) Zero(value complex128, template ...string) *ValidatorComplex128 {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(complex128) == 0
		},
		ErrorKeyZero, value, template...)

	return validator
}

func (validator *ValidatorComplex128) Passing(function func(v complex128) bool, template ...string) *ValidatorComplex128 {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(complex128))
		},
		ErrorKeyPassing, template...)

	return validator
}

func (validator *ValidatorComplex128) InSlice(slice []complex128, template ...string) *ValidatorComplex128 {
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

func (validator *ValidatorComplex128) Not() *ValidatorComplex128 {
	validator.context.Not()

	return validator
}
