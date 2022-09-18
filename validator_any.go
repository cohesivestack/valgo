package valgo

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

func (validator *ValidatorAny) InSlice(slice []any, template ...string) *ValidatorAny {
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
