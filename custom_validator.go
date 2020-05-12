package valgo

type CustomValidator struct {
	validator *validatorContext
}

func (cv *CustomValidator) Invalidate(
	errorKey string, values map[string]interface{}, template ...string) {

	if values == nil {
		values = map[string]interface{}{}
	}

	if _, ok := values["title"]; !ok {
		values["title"] = cv.validator.currentTitle
	}

	cv.validator.invalidate(errorKey, values, template...)
}

func (cv *CustomValidator) Value() interface{} {
	return cv.validator.currentValue
}

func (cv *CustomValidator) Title() *string {
	return cv.validator.currentTitle
}
