package valgo

type GenericValidator struct {
	*validatorContext
}

func (v *GenericValidator) EqualTo(value interface{}, template ...string) *GenericValidator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(IsEqualTo(v.currentValue, value)) {
		v.invalidate("equal_to", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}
