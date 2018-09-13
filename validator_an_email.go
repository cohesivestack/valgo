package valgo

// Credit to https://github.com/badoux/checkmail/blob/master/checkmail.go
const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func (value *Value) IsAnEmail() bool {
	return value.IsMatchingTo(emailRegexPattern)
}

func (validator *Validator) AnEmail(template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsAnEmail()) {
		validator.invalidate("an_email",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": validator.currentValue}, template)
	}

	validator.resetNegative()

	return validator
}
