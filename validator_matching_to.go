package valgo

import (
	"regexp"
)

func (value *Value) IsMatchingTo(pattern string) bool {
	if !value.IsString() {
		return false
	}

	var r = regexp.MustCompile(pattern)
	return r.MatchString(value.AsString())
}

func (validator *Validator) MatchingTo(pattern string, template ...string) *Validator {
	if !validator.assert(validator.currentValue.IsMatchingTo(pattern)) {
		validator.invalidate("matching_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": pattern}, template)
	}

	validator.resetNegative()

	return validator
}
