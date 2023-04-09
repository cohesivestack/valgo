package valgo

// FactoryOptions is a struct in Go that is used to pass options to a [Factory()]
type FactoryOptions struct {
	// A string field that represents the default locale code to use by the
	// factory if a specific locale code is not provided when a Validation is
	// created
	LocaleCodeDefault string
	// A map field that allows to modify the current or add new locales
	Locales map[string]*Locale
	// A function field that allows to set a custom JSON marshaler for [Error]
	MarshalJsonFunc func(e *Error) ([]byte, error)
}

// ValidationFactory is a struct provided by Valgo that enables the creation of
// Validation sessions with preset options. This allows for more flexibility and
// easier management of options when creating [Validation] sessions, as it avoids
// having to pass options each time a new [Validation] is created.
//
// The [Factory()] function is used to create a ValidationFactory instance, and
// it takes a [FactoryOptions] struct as a parameter. This allows customization
// of the default locale code, addition of new locales, and setting a custom
// JSON marshaler for errors.
//
// A ValidationFactory instance offers all the functions for creating
// Validations available at the package level ([Is()], [In()], [Check()], [New()]),
// which create new Validation sessions with the preset options defined in the
// factory.

type ValidationFactory struct {
	localeCodeDefault string
	locales           map[string]*Locale
	marshalJsonFunc   func(e *Error) ([]byte, error)
}

// This New function allows you to create, through a factory, a new Validation
// session without a Validator. This is useful for conditional validation or
// reusing validation logic.
//
// The function is similar to the [New()] function, but it uses a factory.
// For more information see the [New()] function.
func (_factory *ValidationFactory) New(options ...Options) *Validation {

	var _options *Options
	finalOptions := Options{
		localeCodeDefaultFromFactory: _factory.localeCodeDefault,
	}

	if _factory.locales != nil {
		finalOptions.localesFromFactory = _factory.locales
	}

	if len(options) > 0 {
		_options = &options[0]
	}

	if _options != nil && _options.LocaleCode != "" {
		finalOptions.LocaleCode = _options.LocaleCode
	}

	if _options != nil && _options.MarshalJsonFunc != nil {
		finalOptions.MarshalJsonFunc = _options.MarshalJsonFunc
	} else if _factory.marshalJsonFunc != nil {
		finalOptions.MarshalJsonFunc = _factory.marshalJsonFunc
	}

	return newValidation(finalOptions)
}

// The Is function allows you to pass, through a factory, a [Validator]
// with the value and the rules for validating it. At the same time, create a
// [Validation] session, which lets you add more Validators in order to verify
// more values.
//
// The function is similar to the [Is()] function, but it uses a factory.
// For more information see the [Is()] function.
func (_factory *ValidationFactory) Is(v Validator) *Validation {
	return _factory.New().Is(v)
}

// The In function executes, through a factory, one or more validators in a
// namespace, so the value names in the error result are prefixed with this
// namespace. This is useful for validating nested structures.
//
// The function is similar to the [In()] function, but it uses a factory.
// For more information see the [In()] function.
func (_factory *ValidationFactory) In(name string, v *Validation) *Validation {
	return _factory.New().In(name, v)
}

// The InRow function executes, through a factory, one or more validators in a
// namespace similar to the [In](...) function, but with indexed namespace. So,\
// the value names in the error result are prefixed with this indexed namespace.
// It is useful for validating nested lists in structures.
//
// The function is similar to the [InRow()] function, but it uses a factory.
// For more information see the [InRow()] function.
func (_factory *ValidationFactory) InRow(name string, index int, v *Validation) *Validation {
	return _factory.New().InRow(name, index, v)
}

// The Check function, through a factory, is similar to the Is function, however
// with Check the Rules of the [Validator] parameter are not short-circuited,
// which means that regardless of whether a previous rule was valid, all rules
// are checked.
//
// The function is similar to the [Check()] function, but it uses a factory.
// For more information see the [Check()] function.
func (_factory *ValidationFactory) Check(v Validator) *Validation {
	return _factory.New().Check(v)
}

// Create a new [Validation] session, through a factory, and add an error
// message to it without executing a field validator. By adding this error
// message, the [Validation] session will be marked as invalid.
func (_factory *ValidationFactory) AddErrorMessage(name string, message string) *Validation {
	return _factory.New().AddErrorMessage(name, message)
}
