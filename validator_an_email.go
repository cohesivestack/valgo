package valgo

// Credit to https://github.com/badoux/checkmail/blob/master/checkmail.go
const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func (val *Value) IsAnEmail() bool {
	return val.IsMatchingTo(emailRegexPattern)
}

func (v *Validator) AnEmail(template ...string) *Validator {
	if !v.assert(v.currentValue.IsAnEmail()) {
		v.invalidate("an_email",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": v.currentValue}, template)
	}

	v.resetNegative()

	return v
}
