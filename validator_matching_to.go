package valgo

import (
	"regexp"
)

func IsMatchingTo(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

func (v *Validator) MatchingTo(regex *regexp.Regexp, template ...string) *Validator {
	if v.currentDataType != DataTypeString {
		panic("Empty validator requires a string as value")
	} else if !v.assert(IsMatchingTo(v.currentValue.(string), regex)) {
		v.invalidate("matching_to", map[string]interface{}{
			"Title": v.currentTitle,
			"Value": regex.String()}, template...)
	}

	v.resetNegative()

	return v
}
