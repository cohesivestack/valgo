package valgo

type Int64Validator struct {
	*validatorContext
}

func IsInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return NewValidator().IsInt64(value, nameAndTitle...)
}

func CheckInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return NewValidator().CheckInt64(value, nameAndTitle...)
}

func (v *validatorContext) IsInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return v.isInt64(true, value, nameAndTitle...)
}

func (v *validatorContext) CheckInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return v.isInt64(false, value, nameAndTitle...)
}

func (l *localized) IsInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return l.NewValidator().IsInt64(value, nameAndTitle...)
}

func (l *localized) CheckInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return l.NewValidator().CheckInt64(value, nameAndTitle...)
}

func (v *validatorContext) isInt64(shortCircuit bool, value int64, nameAndTitle ...string) *Int64Validator {
	v.currentDataType = DataTypeInt64
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
	return &Int64Validator{v}
}

func (v *Int64Validator) EqualTo(value int64, template ...string) *Int64Validator {
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

func (v *Int64Validator) GreaterThan(value int64, template ...string) *Int64Validator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(v.currentValue.(int64) > value) {
		v.invalidate("greater_than", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}

func (v *Int64Validator) GreaterOrEqualThan(value int64, template ...string) *Int64Validator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(v.currentValue.(int64) >= value) {
		v.invalidate("greater_or_equal_than", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}

func (v *Int64Validator) LessThan(value int64, template ...string) *Int64Validator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(v.currentValue.(int64) < value) {
		v.invalidate("less_than", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}

func (v *Int64Validator) LessOrEqualThan(value int64, template ...string) *Int64Validator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(v.currentValue.(int64) <= value) {
		v.invalidate("less_or_equal_than", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}

func (v *Int64Validator) Not() *Int64Validator {
	v.currentNegative = true

	return v
}

func (v *Int64Validator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *Int64Validator {

	if v.isShortCircuit() {
		return v
	}

	customValidator := CustomValidator{
		validator: v.validatorContext,
	}

	function(&customValidator, template...)

	v.resetNegative()

	return v
}

func (v *Int64Validator) InSlice(slice []interface{}, template ...string) *Int64Validator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(IsInSlice(v.currentValue, slice)) {
		v.invalidate("in_slice", map[string]interface{}{
			"title": v.currentTitle,
			"value": v.currentValue}, template...)
	}

	v.resetNegative()

	return v
}
