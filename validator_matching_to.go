package valgo

import (
	"regexp"
)

func matchingTo(value interface{}, pattern string) bool {
	var _value string
	switch value.(type) {
	case string:
		_value = value.(string)
	case *string:
		_value = *value.(*string)
	default:
		return false
	}

	var r = regexp.MustCompile(pattern)
	return r.MatchString(_value)
}

func (validator *Validator) MatchingTo(pattern string, template ...string) *Validator {

	if !matchingTo(validator.currentValue, pattern) {
		validator.invalidate("matching_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": pattern}, template)
	}
	return validator
}

func (validator *Validator) NotMatchingTo(pattern string, template ...string) *Validator {

	if matchingTo(validator.currentValueAsString(), pattern) {
		validator.invalidate("not_matching_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": pattern}, template)
	}
	return validator
}
