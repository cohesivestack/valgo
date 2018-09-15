package valgo

type CustomValidator struct {
	validator *Validator
}

func (cv *CustomValidator) Invalidate(
	key string, templateString []string, variables map[string]interface{}) {

	if variables == nil {
		variables = map[string]interface{}{}
	}

	if _, ok := variables["Title"]; !ok {
		variables["Title"] = cv.validator.currentTitle
	}

	cv.validator.invalidate(key, variables, templateString)
}

func (cv *CustomValidator) Value() *Value {
	return cv.validator.currentValue
}
