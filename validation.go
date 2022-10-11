package valgo

import (
	"strconv"
)

type Validation struct {
	valid bool

	_locale      *locale
	errors       map[string]*valueError
	currentIndex int
}

func (validation *Validation) Is(v Validator) *Validation {
	return v.Context().validateIs(validation)
}

func (validation *Validation) Check(v Validator) *Validation {
	return v.Context().validateCheck(validation)
}

func (validation *Validation) Valid() bool {
	return validation.valid
}

func (validation *Validation) Merge(_validation *Validation) *Validation {

LOOP1:
	for _field, _err := range _validation.Errors() {
		for field, err := range validation.Errors() {
			if _field == field {
				for _, _errMsg := range _err.messages {
					for _, errMsg := range err.messages {
						if _errMsg == errMsg {
							continue LOOP1
						}
					}
					validation.AddErrorMessage(_field, _errMsg)
				}
				continue LOOP1
			}
		}
		for _, _errMsg := range _err.messages {
			validation.AddErrorMessage(_field, _errMsg)
		}
	}
	return validation
}

func (v *Validation) AddErrorMessage(name string, message string) *Validation {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.valid = false

	ev := v.getOrCreateValueError(name)

	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (validation *Validation) invalidate(name *string, fragment *ValidatorFragment) {
	if validation.errors == nil {
		validation.errors = map[string]*valueError{}
	}

	validation.valid = false

	var _name string
	if name == nil {
		_name = concatString("value_", strconv.Itoa(validation.currentIndex-1))
	} else {
		_name = *name
	}

	ev := validation.getOrCreateValueError(_name)

	errorKey := fragment.errorKey

	if !fragment.boolOperation {
		errorKey = concatString("not_", errorKey)
	}

	if _, ok := ev.errorTemplates[errorKey]; !ok {
		ev.errorTemplates[errorKey] = &errorTemplate{
			key: errorKey,
		}
	}

	et := ev.errorTemplates[errorKey]
	if len(fragment.template) > 0 {
		et.template = &fragment.template[0]
	}
	et.params = fragment.templateParams
}

func (session *Validation) Errors() map[string]*valueError {
	return session.errors
}

func (validation *Validation) Error() error {
	if !validation.valid {
		return &Error{
			errors: validation.errors,
		}
	}
	return nil
}

func (validation *Validation) IsValid(name string) bool {
	if _, isNotValid := validation.errors[name]; isNotValid {
		return false
	}

	return true
}

func (validation *Validation) getOrCreateValueError(name string) *valueError {
	if _, ok := validation.errors[name]; !ok {
		validation.errors[name] = &valueError{
			name:           &name,
			errorTemplates: map[string]*errorTemplate{},
			errorMessages:  []string{},
			validator:      validation,
		}
	}

	ev := validation.errors[name]
	ev.dirty = true

	return ev
}

func newValidation(_locale *locale) *Validation {
	v := &Validation{
		valid:   true,
		_locale: _locale,
	}

	return v
}
