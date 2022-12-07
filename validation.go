package valgo

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Validation The [Validation] session in Valgo is the main structure for validating one or
// more values. It is called Validation in code.
//
// A [Validation] session will contain one or more Validators, where each [Validator]
// will have the responsibility to validate a value with one or more rules.
//
// There are multiple functions to create a [Validation] session, depending on the
// requirements:
//
//   - [New]()
//   - [Is](...)
//   - [In](...)
//   - [InRow](...)
//   - [Check](...)
//   - [AddErrorMessage](...)
//
// the function [Is](...) is likely to be the most frequently used function in your
// validations. When [Is](...) is called, the function creates a validation and
// receives a validator at the same time.
type Validation struct {
	valid bool

	_locale      *locale
	errors       map[string]*valueError
	currentIndex int

	sync.RWMutex
}

// Is Add a field validator to a [Validation] session.
func (v *Validation) Is(vr Validator) *Validation {
	return vr.Context().validateIs(v)
}

// Check Add a field validator to a [Validation] session. But unlike [Is()] the
// field validator is not short-circuited.
func (v *Validation) Check(vr Validator) *Validation {
	return vr.Context().validateCheck(v)
}

// Valid A [Validation] session provides this function which returns either true if
// all their validators are valid or false if any one of them is invalid.
//
// In the following example, even though the [Validator] for age is valid, the
// [Validator] for status is invalid, making the entire Validator session
// invalid.
func (v *Validation) Valid() bool {
	return v.valid
}

// In Add a map namespace to a [Validation] session.
func (v *Validation) In(name string, _validation *Validation) *Validation {
	return v.merge(name, _validation)
}

// InRow Add an indexed namespace to a [Validation] session.
func (v *Validation) InRow(name string, index int, _validation *Validation) *Validation {
	return v.merge(fmt.Sprintf("%s[%v]", name, index), _validation)
}

// Merge Using [Merge](...) you can merge two [Validation] sessions. When two
// validations are merged, errors with the same value name will be merged. It is
// useful for reusing validation logic.
//
// The following example merges the [Validation] session returned by the
// validatePreStatus function. Since both [Validation] sessions validate a value
// with the name status, the error returned will return two error messages, and
// without duplicate the Not().Blank() error message rule.
func (v *Validation) Merge(_validation *Validation) *Validation {
	return v.merge("", _validation)
}

//nolint:gocognit // by initial design. should be refactored to be simplified
func (v *Validation) merge(prefix string, _validation *Validation) *Validation {
	var _prefix string
	if len(strings.TrimSpace(prefix)) > 0 {
		_prefix = prefix + "."
	}

LOOP1:
	for _field, _err := range _validation.Errors() {
		for field, err := range v.Errors() {
			if _prefix+_field == field {
			LOOP2:
				for _, _errMsg := range _err.Messages() {
					for _, errMsg := range err.Messages() {
						if _errMsg == errMsg {
							continue LOOP2
						}
					}
					v.AddErrorMessage(_prefix+_field, _errMsg)
				}

				continue LOOP1
			}
		}

		for _, _errMsg := range _err.Messages() {
			v.AddErrorMessage(_prefix+_field, _errMsg)
		}
	}

	return v
}

// AddErrorMessage Add an error message to the [Validation] session without executing a field
// validator. By adding this error message, the [Validation] session will be
// marked as invalid.
func (v *Validation) AddErrorMessage(name string, message string) *Validation {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.valid = false

	ev := v.getOrCreateValueError(name)

	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (v *Validation) invalidate(name *string, fragment *validatorFragment) {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.valid = false

	var _name string
	if name == nil {
		_name = concatString("value_", strconv.Itoa(v.currentIndex-1))
	} else {
		_name = *name
	}

	ev := v.getOrCreateValueError(_name)

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

// Errors Return a map with the information for each invalid field validator in the
// [Validation] session.
//
//nolint:revive // by design. should be exported as can be annoying to use
func (v *Validation) Errors() map[string]*valueError {
	return v.errors
}

// ErrorByKey Return a map with the information for each invalid field validator in the
// [Validation] session.
//
//nolint:revive // by design. should be exported as can be annoying to use
func (v *Validation) ErrorByKey(key string) *valueError {
	v.RLock()
	err := v.errors[key]
	v.RUnlock()

	return err
}

// Error Return a map with the information for each invalid field validator in the
// [Validation] session.
func (v *Validation) Error() error {
	if v.valid {
		return nil
	}

	return &Error{errors: v.errors}
}

// IsValid Return true if a specific field validator is valid.
func (v *Validation) IsValid(name string) bool {
	if _, isNotValid := v.errors[name]; isNotValid {
		return false
	}

	return true
}

func (v *Validation) getOrCreateValueError(name string) *valueError {
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

func newValidation(_locale *locale) *Validation {
	v := &Validation{
		valid:   true,
		_locale: _locale,
	}

	return v
}
