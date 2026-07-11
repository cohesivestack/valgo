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
	invalidateMap   map[string]bool
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

// [IfValid](...) is similar to [If](...), but merges the given [Validation] session
// only when the current [Validation] session is valid (i.e., no validation errors
// have been produced so far). When the current session is not valid, no operation
// is performed and the original instance is returned unchanged.
//
// See [Valid](...) for the predicate used by this method.
//
//	v.Is(v.String(username, "username").Not().Blank()).
//		IfValid(v.Is(v.String(role, "role").Equal("admin")))
func (validation *Validation) IfValid(_validation *Validation) *Validation {
	if validation.Valid() {
		return validation.merge("", _validation)
	}
	return validation
}

// [IfPathValid](...) is similar to [If](...), but merges the given [Validation] session
// only when the specified path is valid in the current [Validation] session.
// When the path is not valid, no operation is performed and the original instance
// is returned unchanged.
//
// See [PathValid](...) for how paths (including parent namespaces) are evaluated.
//
//	v.Is(v.String(username, "username").Not().Blank()).
//		IfPathValid("username", v.Is(v.String(role, "role").Equal("admin")))
func (validation *Validation) IfPathValid(path string, _validation *Validation) *Validation {
	if validation.PathValid(path) {
		return validation.merge("", _validation)
	}
	return validation
}

// [IfAllValid](...) is similar to [If](...), but merges the given [Validation] session
// only when all provided paths are valid in the current [Validation] session.
// When any of the paths is not valid, no operation is performed and the original
// instance is returned unchanged.
//
// When the slice is empty, it behaves like [AllValid](...) called with no arguments,
// which is equivalent to [Valid]().
//
//	v.Is(
//		v.String(email, "email").Not().Empty(),
//		v.String(password, "password").Not().Empty(),
//	).IfAllValid([]string{"email", "password"},
//		v.Is(v.String(role, "role").Equal("user")),
//	)
func (validation *Validation) IfAllValid(paths []string, _validation *Validation) *Validation {
	if validation.AllValid(paths...) {
		return validation.merge("", _validation)
	}
	return validation
}

// [IfAnyValid](...) is similar to [If](...), but merges the given [Validation] session
// only when at least one of the provided paths is valid in the current
// [Validation] session. When none of the paths are valid, no operation is performed
// and the original instance is returned unchanged.
//
// When the slice is empty, it behaves like [AnyValid](...) called with no arguments,
// which returns false.
//
//	v.Is(
//		v.String(email, "email").Email(),
//		v.String(phone, "phone").Not().Empty(),
//	).IfAnyValid([]string{"email", "phone"},
//		v.Is(v.String(preferred, "preferred").InSlice([]string{"email", "phone"})),
//	)
func (validation *Validation) IfAnyValid(paths []string, _validation *Validation) *Validation {
	if validation.AnyValid(paths...) {
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

// [WhenValid](...) is similar to [When](...), but executes the given function
// only when the current [Validation] session is valid (i.e., no validation errors
// have been produced so far). When the current session is not valid, no operation
// is performed and the original instance is returned unchanged.
//
// See [Valid](...) for the predicate used by this method.
//
//	v.Is(v.String(username, "username").Not().Blank()).
//		WhenValid(func(val *v.Validation) {
//			val.Is(v.String(role, "role").Equal("admin"))
//		})
func (validation *Validation) WhenValid(function func(val *Validation)) *Validation {
	if validation.Valid() {
		function(validation)
	}
	return validation
}

// [WhenPathValid](...) is similar to [When](...), but executes the given function
// only when the specified path is valid in the current [Validation] session.
// When the path is not valid, no operation is performed and the original
// instance is returned unchanged.
//
// See [PathValid](...) for how paths (including parent namespaces) are evaluated.
//
//	v.Is(v.String(username, "username").Not().Blank()).
//		WhenPathValid("username", func(val *v.Validation) {
//			val.Is(v.String(role, "role").Equal("admin"))
//		})
func (validation *Validation) WhenPathValid(path string, function func(val *Validation)) *Validation {
	if validation.PathValid(path) {
		function(validation)
	}
	return validation
}

// [WhenAllValid](...) is similar to [When](...), but executes the given function
// only when all provided paths are valid in the current [Validation] session.
// When any of the paths is not valid, no operation is performed and the original
// instance is returned unchanged.
//
// When the slice is empty, it behaves like [AllValid](...) called with no arguments,
// which is equivalent to [Valid]().
//
//	v.Is(
//		v.String(email, "email").Not().Empty(),
//		v.String(password, "password").Not().Empty(),
//	).WhenAllValid([]string{"email", "password"}, func(val *v.Validation) {
//		val.Is(v.String(role, "role").Equal("user"))
//	})
func (validation *Validation) WhenAllValid(paths []string, function func(val *Validation)) *Validation {
	if validation.AllValid(paths...) {
		function(validation)
	}
	return validation
}

// [WhenAnyValid](...) is similar to [When](...), but executes the given function
// only when at least one of the provided paths is valid in the current
// [Validation] session. When none of the paths are valid, no operation is
// performed and the original instance is returned unchanged.
//
// When the slice is empty, it behaves like [AnyValid](...) called with no arguments,
// which returns false.
//
//	v.Is(
//		v.String(email, "email").Email(),
//		v.String(phone, "phone").Not().Empty(),
//	).WhenAnyValid([]string{"email", "phone"}, func(val *v.Validation) {
//		val.Is(v.String(preferred, "preferred").InSlice([]string{"email", "phone"}))
//	})
func (validation *Validation) WhenAnyValid(paths []string, function func(val *Validation)) *Validation {
	if validation.AnyValid(paths...) {
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
	v.valid = false

	ev := v.getOrCreateValueError(name, nil)

	ev.errorMessages = append(ev.errorMessages, message)

	return v
}

func (v *Validation) mergeError(prefix string, err *Error) *Validation {

	if err != nil && len(err.errors) > 0 {
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

func (validation *Validation) invalidate(name *string, title *string, invalidFragments []*invalidFragment) {
	validation.valid = false

	var _name string
	if name == nil {
		_name = "value_" + strconv.Itoa(validation.currentIndex-1)
	} else {
		_name = *name
	}

	ev := validation.getOrCreateValueError(_name, title)

	for _, invalidFragment := range invalidFragments {
		isOrFragment := len(invalidFragment.fragments) > 1
		etOneOf := &errorTemplateOneOf{}
		for i, fragment := range invalidFragment.fragments {
			errorKey := fragment.errorKey
			if !fragment.boolOperation {
				errorKey = "not_" + errorKey
			}
			et := &errorTemplate{
				key:    errorKey,
				params: fragment.templateParams,
			}
			if len(fragment.template) > 0 {
				et.template = &fragment.template[0]
			}
			if i == 0 {
				if isOrFragment {
					etOneOf.errorTemplates = []*errorTemplate{}
				} else {
					etOneOf.errorTemplate = et
				}
			}
			if isOrFragment {
				etOneOf.errorTemplates = append(etOneOf.errorTemplates, et)
			}
		}
		ev.errorTemplates = append(ev.errorTemplates, etOneOf)
	}
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

// Deprecated: use PathValid(path) instead.
//
// IsValid reports whether the validator result for the given path is valid.
//
// A path can be a simple field name (e.g. "email") or a nested/indexed namespace
// (e.g. "person.addresses[0].line1").
//
// A path is considered valid if it does not appear in the internal invalidation
// map produced during validation. In nested or indexed namespaces, parent paths
// of an invalid field are also considered invalid.
func (validation *Validation) IsValid(path string) bool {
	if _, isNotValid := validation.invalidateMap[path]; isNotValid {
		return false
	}
	return true
}

// PathValid reports whether the validator result for the given field path is valid.
//
// A path can be a simple field name (e.g. "email") or a nested/indexed namespace
// (e.g. "person.addresses[0].line1").
//
// A path is considered valid if it does not appear in the internal invalidation
// map produced during validation. In nested or indexed namespaces, parent paths
// of an invalid field are also considered invalid.
//
// Note: PathValid does not validate again; it only queries the results of a
// previous Is(...) or Check(...) call.
//
// Example:
//
//	val := Is(
//		String("", "email").Not().Empty(),   // invalid
//		String("John", "name").Not().Blank(), // valid
//	)
//
//	_ = val.PathValid("email") // false
//	_ = val.PathValid("name")  // true
//
// Example (nested namespaces):
//
//	val := In("person",
//		InRow("addresses", 0,
//			Is(String("", "line1").Not().Blank()), // invalid
//		),
//	)
//
//	_ = val.PathValid("person.addresses[0].line1") // false
//	_ = val.PathValid("person.addresses[0]")       // false (parent path)
//	_ = val.PathValid("person.addresses")          // false (parent path)
//	_ = val.PathValid("person")                    // false (parent path)
//	_ = val.PathValid("person.addresses[1]")       // true  (unrelated path)
func (validation *Validation) PathValid(path string) bool {
	if _, isNotValid := validation.invalidateMap[path]; isNotValid {
		return false
	}
	return true
}

// AllValid reports whether all provided paths are valid.
//
// When one or more paths are provided, AllValid returns true only if every path
// is valid.
//
// When called with no arguments, AllValid returns the overall validation result,
// equivalent to v.Valid().
//
// Example:
//
//	val := Is(
//		String("", "email").Not().Empty(),        // invalid
//		String("+123", "phone").Not().Empty(),    // valid
//	)
//
//	_ = val.AllValid("phone")            // true
//	_ = val.AllValid("email")            // false
//	_ = val.AllValid("email", "phone")   // false
//	_ = val.AllValid()                   // false (same as val.Valid())
func (v *Validation) AllValid(paths ...string) bool {
	if len(paths) == 0 {
		return v.Valid()
	}
	for _, path := range paths {
		if _, invalid := v.invalidateMap[path]; invalid {
			return false
		}
	}
	return true
}

// AnyValid reports whether at least one of the provided paths is valid.
//
// AnyValid returns true as soon as it finds a valid path.
//
// When called with no arguments, AnyValid returns false. The caller must
// explicitly provide the set of paths to evaluate.
//
// Example:
//
//	val := Is(
//		String("", "email").Not().Empty(),        // invalid
//		String("+123", "phone").Not().Empty(),    // valid
//	)
//
//	_ = val.AnyValid("email")           // false
//	_ = val.AnyValid("email", "phone")  // true
//	_ = val.AnyValid()                  // false
func (v *Validation) AnyValid(paths ...string) bool {
	if len(paths) == 0 {
		return false
	}
	for _, path := range paths {
		if _, invalid := v.invalidateMap[path]; !invalid {
			return true
		}
	}
	return false
}

func (validation *Validation) getOrCreateValueError(name string, title *string) *valueError {
	if validation.errors == nil {
		validation.errors = map[string]*valueError{}
		validation.invalidateMap = map[string]bool{}
	}

	if _, ok := validation.errors[name]; !ok {
		validation.addInvalidationNamespaces(name)
		validation.errors[name] = &valueError{
			name:           &name,
			title:          title,
			errorTemplates: []*errorTemplateOneOf{},
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

// name examples:
//
//	"object.users[1].value"
//
// namespaces generated:
//
//	"object"
//	"object.users"
//	"object.users[1]"
//	"object.users[1].value"
func (validation *Validation) addInvalidationNamespaces(name string) {
	if name == "" {
		return
	}

	segStart := 0 // start index of current segment (after last '.')
	bracketAdded := false

	for i := 0; i < len(name); i++ {
		switch name[i] {
		case '[':
			// First '[' in this segment: add prefix without the index.
			// e.g. "object.users[1]" -> add "object.users".
			if !bracketAdded && i > segStart {
				bracketAdded = true
				validation.invalidateMap[name[:i]] = true
			}
		case '.':
			// End of segment: add prefix up to this dot.
			// e.g. "object.users[1].value" at '.' after "[1]" -> add "object.users[1]".
			if i > 0 {
				validation.invalidateMap[name[:i]] = true
			}
			segStart = i + 1
			bracketAdded = false
		}
	}

	// Always add the full path
	validation.invalidateMap[name] = true
}
