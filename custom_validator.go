package valgo

type CustomValidator struct {
	validator *Validator
}

func (customValidator *CustomValidator) Invalidate(
	key string, templateString []string, variables map[string]interface{}) {

	if variables == nil {
		variables = map[string]interface{}{}
	}

	if _, ok := variables["Title"]; !ok {
		variables["Title"] = customValidator.validator.currentTitle
	}

	customValidator.validator.invalidate(key, variables, templateString)
}

func (customValidator *CustomValidator) Value() interface{} {
	return customValidator.validator.currentValue
}

func (customValidator *CustomValidator) ValueAsString() string {
	return customValidator.validator.currentValueAsString()
}
