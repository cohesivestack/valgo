package valgo

func (v *Validator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *Validator {

	if v.isShortCircuit() {
		return v
	}

	customValidator := CustomValidator{
		validator: v,
	}

	function(&customValidator, template...)

	v.resetNegative()

	return v
}
