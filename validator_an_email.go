package valgo

import "regexp"

const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

var emailRegex *regexp.Regexp

func init() {
	emailRegex = regexp.MustCompile(emailRegexPattern)
}

func IsAnEmail(value string) bool {
	return IsMatchingTo(value, emailRegex)
}

func (v *Validator) AnEmail(template ...string) *Validator {
	if v.currentDataType != DataTypeString {
		panic("Email validator requires a string as value")
	} else if !v.assert(IsAnEmail(v.currentValue.(string))) {
		v.invalidate("an_email", map[string]interface{}{"title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}
