package valgo

import (
	"strings"
)

func (val *Value) IsBlank() bool {
	if val.isBlank == nil {
		_val := strings.TrimSpace(val.AsString())
		val.isBlank = boolPointer(len(_val) == 0)
	}
	return *val.isBlank
}

func (v *Validator) Blank(template ...string) *Validator {
	if !v.assert(v.currentValue.IsBlank()) {
		v.invalidate("blank", map[string]interface{}{"Title": v.currentTitle}, template)
	}

	v.resetNegative()

	return v
}
