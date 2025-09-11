package valgo

import (
	"fmt"
	"strconv"
	"strings"
)

// The [Validation] session in Valgo is the main structure for validating one or
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
//   - [Check](...)
//   - [InRow](...)
//   - [InCell](...)
//   - [If](...)
//   - [Do](...)
//   - [When](...)
//   - [Merge](...)
//   - [AddErrorMessage](...)
//
// the function [Is](...) is likely to be the most frequently used function in your
// validations. When [Is](...) is called, the function creates a validation and
// receives a validator at the same time.
type Validation struct {
	valid bool

	_locale         *Locale
	errors          map[string]*valueError
	currentIndex    int
	marshalJsonFunc func(e *Error) ([]byte, error)
}

// Options struct is used to specify options when creating a new [Validation]
// session with the [New()] function.
//
// It contains parameters for specifying a specific locale code, modify or add a
// locale, and set a custom JSON marshaler for [Error].

type Options struct {
	localeCodeDefaultFromFactory string             // Only specified by the factory
	localesFromFactory           map[string]*Locale // Only specified by the factory

	// A string field that represents the locale code to use by the [Validation]
	// session
	LocaleCode string
	// A map field that allows to modify or add a new [Locale]
	Locale *Locale
	// A function field that allows to set a custom JSON marshaler for [Error]
	MarshalJsonFunc func(e *Error) ([]byte, error)
}

// Add one or more validators to a [Validation] session.
func (validation *Validation) Is(validators ...Validator) *Validation {
	for _, v := range validators {
		validation = v.Context().validateIs(validation)
	}
	return validation
}

// [If](...) is similar to [Merge](...), but merge the [Validation] session
// only when the condition is true, and returns the same [Validation] instance.
// When the condition is false, no operation is performed and the original
// instance is returned unchanged.
//
// See [Merge](...) for more information.
//
//	v.If(isAdmin, v.Is(v.String(username, "username").Not().Blank()) )
func (validation *Validation) If(condition bool, _validation *Validation) *Validation {
	if condition {
		return validation.merge("", _validation)
	}
	return validation
}

// The [Do](...) function executes the given function with the current
// [Validation] instance and returns the same instance.
//
// This allows you to extend a validation chain with additional or
// conditional rules in a concise way:
//
//	v.Is(v.String(username, "username").Not().Blank()).Do(func(val *v.Validation) {
//		if isAdmin {
//			val.Is(v.String(role, "role").Equal("admin"))
//		}
//	})
func (validation *Validation) Do(function func(val *Validation)) *Validation {
	function(validation)
	return validation
}

// [When](...) is similar to [Do](...), but executes the given function
// only when the condition is true, and returns the same [Validation] instance.
// When the condition is false, no operation is performed and the original
// instance is returned unchanged.
//
// See [Do](...) for the unconditional variant.
//
//	v.Is(v.String(username, "username").Not().Blank()).When(isAdmin, func(val *v.Validation) {
//		val.Is(v.String(role, "role").Equal("admin"))
//	})
func (validation *Validation) When(condition bool, function func(val *Validation)) *Validation {
	if condition {
		function(validation)
	}
	return validation
}

// [Check](...) adds one or more validators to a [Validation] session. But unlike [Is()],
// the validators are not short-circuited.
func (validation *Validation) Check(validators ...Validator) *Validation {
	for _, v := range validators {
		validation = v.Context().validateCheck(validation)
	}
	return validation
}

// A [Validation] session provides this function which returns either true if
// all their validators are valid or false if any one of them is invalid.
//
// In the following example, even though the [Validator] for age is valid, the
// [Validator] for status is invalid, making the entire Validator session
// invalid.
func (validation *Validation) Valid() bool {
	return validation.valid
}

// Add a map namespace to a [Validation] session.
func (validation *Validation) In(name string, _validation *Validation) *Validation {
	return validation.merge(name, _validation)
}

// Add an indexed namespace to a [Validation] session.
func (validation *Validation) InRow(name string, index int, _validation *Validation) *Validation {
	return validation.merge(fmt.Sprintf("%s[%v]", name, index), _validation)
}

// Add an indexed namespace to a [Validation] session where the target is a
// single, scalar value (e.g., entries of a primitive slice). This is useful
// for validating arrays or slices of primitives. Example:
//
//	validation := valgo.InCell("tag_priority", 0,
//		valgo.Is(valgo.String("", "tag_priority", "Tag priority").Not().Blank()),
//	)
//
// The example above validates the value at tag_priority[0].
func (validation *Validation) InCell(name string, index int, _validation *Validation) *Validation {

	fieldName := fmt.Sprintf("%s[%v]", name, index)

	for _, _err := range _validation.Errors() {
		for _, _errMsg := range _err.Messages() {
			if err, ok := validation.Errors()[fieldName]; ok {
				for _, errMsg := range err.Messages() {
					if _errMsg == errMsg {
						continue
					}
				}
			}
			validation.AddErrorMessage(fieldName, _errMsg)
		}
	}
	return validation
}

// Using [Merge](...) you can merge two [Validation] sessions. When two
// validations are merged, errors with the same value name will be merged. It is
// useful for reusing validation logic.
//
// The following example merges the [Validation] session returned by the
// validatePreStatus function. Since both [Validation] sessions validate a value
// with the name status, the error returned will return two error messages, and
// without duplicate the Not().Blank() error message rule.
func (validation *Validation) Merge(_validation *Validation) *Validation {
	return validation.merge("", _validation)
}

func (validation *Validation) merge(prefix string, _validation *Validation) *Validation {

	var _prefix string
	if len(strings.TrimSpace(prefix)) > 0 {
		_prefix = prefix + "."
	}

LOOP1:
	for _field, _err := range _validation.Errors() {
		for field, err := range validation.Errors() {
			if _prefix+_field == field {
			LOOP2:
				for _, _errMsg := range _err.Messages() {
					for _, errMsg := range err.Messages() {
						if _errMsg == errMsg {
							continue LOOP2
						}
					}
					validation.AddErrorMessage(_prefix+_field, _errMsg)
				}
				continue LOOP1
			}
		}
		for _, _errMsg := range _err.Messages() {
			validation.AddErrorMessage(_prefix+_field, _errMsg)
		}
	}
	return validation
}

// Add an error message to the [Validation] session without executing a field
// validator. By adding this error message, the [Validation] session will be
// marked as invalid.
func (v *Validation) AddErrorMessage(name string, message string) *Validation {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.valid = false

	ev := v.getOrCreateValueError(name, nil)

	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (v *Validation) mergeError(prefix string, err *Error) *Validation {

	if err != nil && len(err.errors) > 0 {
		if v.errors == nil {
			v.errors = map[string]*valueError{}
		}

		v.valid = false

		var _prefix string
		if len(strings.TrimSpace(prefix)) > 0 {
			_prefix = prefix + "."
		}

		for name, _ev := range err.errors {
			for _, message := range _ev.Messages() {
				v.AddErrorMessage(_prefix+name, message)
			}
		}
	}

	return v
}

// MergeError allows merging Valgo errors from an already validated [Validation] session.
// The function takes an Valgo [Error] pointer as an argument and returns a [Validation] pointer.
func (v *Validation) MergeError(err *Error) *Validation {
	return v.mergeError("", err)
}

// MergeErrorIn allows merging Valgo errors from already validated [Validation] sessions
// within a map namespace. The function takes a namespace name and an [Error] pointer
// as arguments and returns a [Validation] pointer.
func (v *Validation) MergeErrorIn(name string, err *Error) *Validation {
	return v.mergeError(name, err)
}

// MergeErrorInRow allows merging Valgo errors from already validated [Validation] sessions
// within an indexed namespace. The function takes a namespace name, an index, and an [Error] pointer
// as arguments and returns a [Validation] pointer.
//
// DEPRECATED: This method is deprecated in favor of MergeErrorInIndex().
// The MergeErrorInIndex() method is a generic name to cover errors added by
// InRow() and InCell() validations.
func (v *Validation) MergeErrorInRow(name string, index int, err *Error) *Validation {
	return v.mergeError(fmt.Sprintf("%s[%v]", name, index), err)
}

// MergeErrorInRow allows merging Valgo errors from already validated [Validation] sessions
// within an indexed namespace. These are errors added by InRow() and InCell() validations.
// The function takes a namespace name, an index, and an [Error] pointer
// as arguments and returns a [Validation] pointer.
func (v *Validation) MergeErrorInIndex(name string, index int, err *Error) *Validation {
	return v.mergeError(fmt.Sprintf("%s[%v]", name, index), err)
}

func (validation *Validation) invalidate(name *string, title *string, fragment *validatorFragment) {
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

	ev := validation.getOrCreateValueError(_name, title)

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

// Return a map with the information for each invalid field validator
// in the Validation session.
func (session *Validation) Errors() map[string]*valueError {
	return session.errors
}

// Error returns the validation errors as a standard Go error interface.
//
// DEPRECATED: This method is deprecated in favor of ToError() or ToValgoError().
// The Error() method name conflicts with Go's error interface implementation
// convention, where Error() typically implements the error interface for a type.
//
// Use ToError() for standard error handling or ToValgoError() for detailed
// validation error information.
func (validation *Validation) Error(marshalJsonFun ...func(e *Error) ([]byte, error)) error {
	return validation.ToError(marshalJsonFun...)
}

// ToError returns the validation errors as a standard Go error interface.
//
// This method is useful for idiomatic error handling and integration with
// Go's native error system. It returns the same underlying error value as
// ToValgoError() but typed as the error interface.
//
// Example:
//
//	val := Is(String("", "name").Not().Blank())
//	if err := val.ToError(); err != nil {
//	    log.Printf("Validation failed: %v", err)
//	    return err
//	}
//
// An optional JSON marshaling function can be passed to customize how the
// validation errors are serialized into JSON. If no function is provided,
// a default marshaling behavior is used.
func (validation *Validation) ToError(marshalJsonFun ...func(e *Error) ([]byte, error)) error {
	// We cannot simply return validation.ToValgoError(marshalJsonFun...) because
	// when ToValgoError returns nil, it's a nil *Error (concrete type), not a nil
	// error interface. This causes issues with error checking functions like
	// assert.NoError() which expect a proper nil error interface.
	if err := validation.ToValgoError(marshalJsonFun...); err != nil {
		return err
	}
	return nil
}

// ToValgoError returns the validation errors as a *valgo.Error type, providing
// access to rich, structured error details. It's essentially a shortcut to
// `ToError().(*valgo.Error)`.
//
// This method returns the underlying *valgo.Error type directly, exposing
// detailed validation information such as per-field messages, templates,
// and localized titles. It's the single source of truth for validation errors.
//
// Example:
//
//	val := Is(String("", "name").Not().Blank())
//	if errInfo := val.ToValgoError(); errInfo != nil {
//	    for field, valueError := range errInfo.Errors() {
//	        fmt.Printf("Field '%s': %v\n", field, valueError.Messages())
//	    }
//	}
//
// An optional JSON marshaling function can be passed to customize how the
// validation errors are serialized into JSON. If no function is provided,
// a default marshaling behavior is used.
func (validation *Validation) ToValgoError(marshalJsonFun ...func(e *Error) ([]byte, error)) *Error {
	if !validation.valid {
		fn := validation.marshalJsonFunc
		if len(marshalJsonFun) > 0 {
			fn = marshalJsonFun[0]
		}
		return &Error{
			errors:          validation.errors,
			marshalJsonFunc: fn,
		}
	}
	return nil
}

// Return true if a specific field validator is valid.
func (validation *Validation) IsValid(name string) bool {
	if _, isNotValid := validation.errors[name]; isNotValid {
		return false
	}

	return true
}

func (validation *Validation) getOrCreateValueError(name string, title *string) *valueError {
	if _, ok := validation.errors[name]; !ok {
		validation.errors[name] = &valueError{
			name:           &name,
			title:          title,
			errorTemplates: map[string]*errorTemplate{},
			errorMessages:  []string{},
			validator:      validation,
		}
	}

	ev := validation.errors[name]
	ev.dirty = true

	return ev
}

func newValidation(options ...Options) *Validation {
	v := &Validation{
		valid: true,
	}

	if len(options) == 0 {
		v._locale = getLocale(localeCodeDefault)
	} else {
		_options := options[0]

		// If the factory has default locale specified, we try to use it as fallback
		if options[0].localeCodeDefaultFromFactory != "" {
			// Skipping default option will return nil, so we can use the factory
			// locale default
			v._locale = getLocaleAndSkipDefaultOption(_options.LocaleCode, options[0].localesFromFactory)
			if v._locale == nil {
				v._locale = getLocale(options[0].localeCodeDefaultFromFactory, options[0].localesFromFactory)
			}
		} else {
			v._locale = getLocale(_options.LocaleCode, options[0].localesFromFactory)
		}

		// If locale entries were specified, then we merge it with the calculated
		// Locale from the options localeCode
		if _options.Locale != nil {
			v._locale.merge(_options.Locale)
		}
		v.marshalJsonFunc = _options.MarshalJsonFunc
	}

	return v
}
