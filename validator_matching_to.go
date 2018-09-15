package valgo

import (
	"regexp"
)

func (val *Value) IsMatchingTo(pattern string) bool {
	if !val.IsString() {
		return false
	}

	var r = regexp.MustCompile(pattern)
	return r.MatchString(val.AsString())
}

func (v *Validator) MatchingTo(pattern string, template ...string) *Validator {
	if !v.assert(v.currentValue.IsMatchingTo(pattern)) {
		v.invalidate("matching_to",
			map[string]interface{}{
				"Title": v.currentTitle,
				"Value": pattern}, template)
	}

	v.resetNegative()

	return v
}
