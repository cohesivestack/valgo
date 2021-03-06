package valgo

import "strconv"

type validatorContext struct {
	currentValue    interface{}
	currentTitle    *string
	currentName     *string
	currentValid    bool
	currentIndex    int
	currentDataType DataType
	currentNegative bool
	shortCircuit    bool

	_locale *locale
	valid   bool
	errors  map[string]*valueError

	_error error
}

func (v *validatorContext) isShortCircuit() bool {
	return !v.currentValid && v.shortCircuit
}

func (v *validatorContext) Valid() bool {
	return v.valid
}

func (v *validatorContext) Error() error {
	if !v.valid {
		return &Error{
			errors: v.errors,
		}
	}
	return nil
}

func (v *validatorContext) Errors() map[string]*valueError {
	return v.errors
}

func (v *validatorContext) IsValid(name string) bool {
	if _, isNotValid := v.errors[name]; isNotValid {
		return false
	}

	return true
}

func (v *validatorContext) AddErrorMessage(name string, message string) Validator {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.currentValid = false
	v.valid = false

	ev := v.getOrCreateValueError(name)
	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (v *validatorContext) assert(value bool) bool {
	return v.currentNegative != value
}

func (v *validatorContext) resetNegative() {
	v.currentNegative = false
}

func (v *validatorContext) invalidate(errorKey string, values map[string]interface{}, template ...string) {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.currentValid = false
	v.valid = false

	var name string
	if v.currentName == nil {
		name = concatString("value_", strconv.Itoa(v.currentIndex-1))
	} else {
		name = *v.currentName
	}

	ev := v.getOrCreateValueError(name)

	if v.currentNegative {
		errorKey = concatString("not_", errorKey)
	}

	if _, ok := ev.errorTemplates[errorKey]; !ok {
		ev.errorTemplates[errorKey] = &errorTemplate{
			key: errorKey,
		}
	}

	et := ev.errorTemplates[errorKey]
	if len(template) > 0 {
		et.template = &template[0]
	}
	et.values = values
}

func (v *validatorContext) getOrCreateValueError(name string) *valueError {
	if _, ok := v.errors[name]; !ok {
		v.errors[name] = &valueError{
			name:           &name,
			errorTemplates: map[string]*errorTemplate{},
			errorMessages:  []string{},
			validator:      v,
		}
	}

	ev := v.errors[name]
	ev.dirty = true

	return ev
}
