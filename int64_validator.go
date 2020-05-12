package valgo

type Int64Validator struct {
	*validatorContext
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
