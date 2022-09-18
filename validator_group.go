package valgo

import (
	"strconv"
)

type ValidatorGroup struct {
	valid bool

	_locale      *locale
	errors       map[string]*valueError
	currentIndex int
}

func (group *ValidatorGroup) Is(v Validator) *ValidatorGroup {
	return v.Context().validateIs(group)
}

func (group *ValidatorGroup) Check(v Validator) *ValidatorGroup {
	return v.Context().validateCheck(group)
}

func (group *ValidatorGroup) Valid() bool {
	return group.valid
}

func (v *ValidatorGroup) AddErrorMessage(name string, message string) *ValidatorGroup {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.valid = false

	ev := v.getOrCreateValueError(name)

	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (group *ValidatorGroup) invalidate(name *string, fragment *ValidatorFragment) {
	if group.errors == nil {
		group.errors = map[string]*valueError{}
	}

	group.valid = false

	var _name string
	if name == nil {
		_name = concatString("value_", strconv.Itoa(group.currentIndex-1))
	} else {
		_name = *name
	}

	ev := group.getOrCreateValueError(_name)

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

func (group *ValidatorGroup) Errors() map[string]*valueError {
	return group.errors
}

func (group *ValidatorGroup) Error() error {
	if !group.valid {
		return &Error{
			errors: group.errors,
		}
	}
	return nil
}

func (group *ValidatorGroup) IsValid(name string) bool {
	if _, isNotValid := group.errors[name]; isNotValid {
		return false
	}

	return true
}

func (group *ValidatorGroup) getOrCreateValueError(name string) *valueError {
	if _, ok := group.errors[name]; !ok {
		group.errors[name] = &valueError{
			name:           &name,
			errorTemplates: map[string]*errorTemplate{},
			errorMessages:  []string{},
			validator:      group,
		}
	}

	ev := group.errors[name]
	ev.dirty = true

	return ev
}

func newValidatorGroup(_locale *locale) *ValidatorGroup {
	v := &ValidatorGroup{
		valid:   true,
		_locale: _locale,
	}

	return v
}
