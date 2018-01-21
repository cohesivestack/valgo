package valgo

type CustomValidator struct {
	validator *Validator
}

func (customValidator *CustomValidator) Invalidate(
	key string, values map[string]interface{}, templateString []string) {

	customValidator.validator.invalidate(key, values, templateString)
}

func (customValidator *CustomValidator) Value() interface{} {
	return customValidator.validator.currentValue
}

func (customValidator *CustomValidator) ValueAsString() string {
	return customValidator.validator.ensureString()
}
