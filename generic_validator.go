package valgo

type GenericValidator struct {
	*validatorContext
}

func Is(value interface{}, nameAndTitle ...string) *GenericValidator {
	return NewValidator().Is(value, nameAndTitle...)
}

func Check(value interface{}, nameAndTitle ...string) *GenericValidator {
	return NewValidator().Check(value, nameAndTitle...)
}

func (v *validatorContext) Is(value interface{}, nameAndTitle ...string) *GenericValidator {
	return v.is(true, value, nameAndTitle...)
}

func (v *validatorContext) Check(value interface{}, nameAndTitle ...string) *GenericValidator {
	return v.is(false, value, nameAndTitle...)
}

func (v *validatorContext) is(shortCircuit bool, value interface{}, nameAndTitle ...string) *GenericValidator {
	v.currentValue = value
	v.currentIndex += 1
	v.currentValid = true
	v.shortCircuit = shortCircuit

	sizeNameAndTitle := len(nameAndTitle)
	if sizeNameAndTitle > 0 {
		v.currentName = &nameAndTitle[0]
		if sizeNameAndTitle > 1 {
			v.currentTitle = &nameAndTitle[1]
		}
	}
	return &GenericValidator{v}
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
