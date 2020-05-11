package valgo

import (
	"strings"
)

func IsBlank(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func (v *Validator) Blank(template ...string) *Validator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Blank validator requires a string as value")
	} else if !v.assert(IsBlank(v.currentValue.(string))) {
		v.invalidate("blank", map[string]interface{}{"title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}
