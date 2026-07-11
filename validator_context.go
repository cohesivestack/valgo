package valgo

type validatorFragment struct {
	errorKey       string
	template       []string
	templateParams map[string]any
	function       func() bool
	boolOperation  bool
	orOperation    orOperationType
	isValid        bool
}

type orOperationType uint8

const (
	orOperationTypeNone   = 0
	orOperationTypeOr     = 1
	orOperationTypeOrElse = 2
)

// The context keeps the state and provides the functions to control a
// custom validator.
type ValidatorContext struct {
	fragments      []*validatorFragment
	value          any
	name           *string
	title          *string
	fallbackLocale *Locale
	boolOperation  bool
	orOperation    orOperationType
}

// Create a new [ValidatorContext] to be used by a custom validator.
func NewContext(value any, nameAndTitle ...string) *ValidatorContext {

	context := &ValidatorContext{
		value:         value,
		fragments:     []*validatorFragment{},
		boolOperation: true,
		orOperation:   orOperationTypeNone,
	}

	sizeNameAndTitle := len(nameAndTitle)
	if sizeNameAndTitle > 0 {
		name := nameAndTitle[0]
		context.name = &name
		if sizeNameAndTitle > 1 {
			title := nameAndTitle[1]
			context.title = &title
		}
	}

	return context
}

// Invert the boolean value associated with the next validator function in
// a custom validator.
func (ctx *ValidatorContext) Not() *ValidatorContext {
	ctx.boolOperation = false
	return ctx
}

// Add Or operation to validation.
func (ctx *ValidatorContext) Or() *ValidatorContext {
	if len(ctx.fragments) > 0 {
		ctx.orOperation = orOperationTypeOr
	}
	return ctx
}

// Add OrElse operation to validation.
func (ctx *ValidatorContext) OrElse() *ValidatorContext {
	if len(ctx.fragments) > 0 {
		ctx.orOperation = orOperationTypeOrElse
	}
	return ctx
}

// WithLocaleFallback adds locale entries to the current validation session as
// fallbacks.
//
// Entries are merged only if the active locale does not already define the key
// (i.e., it will not override Valgo's built-in locale entries nor any user
// overrides provided via Options{Locale: ...}).
//
// The locales are applied when the validator is executed (Is/Check), not when
// this method is called.
func (ctx *ValidatorContext) WithLocaleFallback(locales ...*Locale) *ValidatorContext {
	if len(locales) == 0 {
		return ctx
	}
	if ctx.fallbackLocale == nil {
		l := Locale{}
		ctx.fallbackLocale = &l
	}
	for _, l := range locales {
		if l == nil {
			continue
		}
		// Only keep the first value for a key, so repeated calls are idempotent.
		for k, v := range *l {
			if _, exists := (*ctx.fallbackLocale)[k]; !exists {
				(*ctx.fallbackLocale)[k] = v
			}
		}
	}
	return ctx
}

// Add a function to a custom validator and pass a value used for the
// validator function to be displayed in the error message.
//
// Use [AddWithParams()] if the error message requires more input values.
func (ctx *ValidatorContext) AddWithValue(function func() bool, errorKey string, value any, template ...string) *ValidatorContext {
	return ctx.AddWithParams(
		function,
		errorKey,
		map[string]any{"title": ctx.title, "value": value}, template...)
}

// Add a function to a custom validator.
func (ctx *ValidatorContext) Add(function func() bool, errorKey string, template ...string) *ValidatorContext {
	return ctx.AddWithParams(
		function,
		errorKey,
		map[string]any{"title": ctx.title}, template...)
}

// Add a function to a custom validator and pass a map with values used for the
// validator function to be displayed in the error message.
func (ctx *ValidatorContext) AddWithParams(function func() bool, errorKey string, templateParams map[string]any, template ...string) *ValidatorContext {

	fragment := &validatorFragment{
		errorKey:       errorKey,
		templateParams: templateParams,
		function:       function,
		boolOperation:  ctx.boolOperation,
		orOperation:    ctx.orOperation,
		isValid:        true,
	}
	if len(template) > 0 {
		fragment.template = template
	}
	ctx.fragments = append(ctx.fragments, fragment)
	ctx.boolOperation = true
	ctx.orOperation = orOperationTypeNone

	return ctx
}

func (ctx *ValidatorContext) validateIs(validation *Validation) *Validation {
	return ctx.validate(validation, true)
}

func (ctx *ValidatorContext) validateCheck(validation *Validation) *Validation {
	return ctx.validate(validation, false)
}

type invalidFragment struct {
	fragments []*validatorFragment // When the fragment is part of an "or" operation, this is a list of fragments that are part of the "or" operation
}

func (ctx *ValidatorContext) validate(validation *Validation, shortCircuit bool) *Validation {
	// valid := true
	validation.currentIndex++

	// Apply fallback locales (if any) without mutating shared locale maps.
	if ctx.fallbackLocale != nil && validation._locale != nil {
		// Copy-on-write: only clone when at least one fallback key is missing.
		needClone := false
		for k := range *ctx.fallbackLocale {
			if _, exists := (*validation._locale)[k]; !exists {
				needClone = true
				break
			}
		}
		if needClone {
			localeCopy := make(Locale, len(*validation._locale)+len(*ctx.fallbackLocale))
			for k, v := range *validation._locale {
				localeCopy[k] = v
			}
			for k, v := range *ctx.fallbackLocale {
				if _, exists := localeCopy[k]; !exists {
					localeCopy[k] = v
				}
			}
			validation._locale = &localeCopy
		}
	}

	invalidFragments := []*invalidFragment{}

	// Iterating through each fragment in the context's fragment list
	for i, fragment := range ctx.fragments {

		// If the previous fragment is not valid, the current fragment is not in an "or" operation, and the short circuit flag is true,
		// we return the current state of the validation without evaluating the current fragment
		if i > 0 && !ctx.fragments[i-1].isValid && fragment.orOperation == orOperationTypeNone && shortCircuit {
			break
		}

		// If the current fragment is a part of an "or" operation and the previous fragment in the "or" operation
		// is valid, we mark the current fragment as valid and move to the next iteration
		if fragment.orOperation == orOperationTypeOr && ctx.fragments[i-1].isValid {
			continue
		}

		//
		if fragment.orOperation == orOperationTypeOrElse && ctx.fragments[i-1].isValid {
			break
		}

		// Evaluating the validation function of the current fragment and updating the valid flag
		// The valid flag will be true only if the fragment function returns a value matching the fragment's boolean operation
		// and the valid flag was true before this evaluation
		fragment.isValid = fragment.function() == fragment.boolOperation

		if !fragment.isValid {
			if fragment.orOperation != orOperationTypeNone {
				invalidFragments[len(invalidFragments)-1].fragments = append(invalidFragments[len(invalidFragments)-1].fragments, fragment)
			} else {
				invalidFragments = append(invalidFragments, &invalidFragment{
					fragments: []*validatorFragment{fragment},
				})
			}
			// If the current fragment is valid and is part of an "or" operation, we remove
			// the invalid fragment from the invalid fragments list
		} else if fragment.orOperation != orOperationTypeNone {
			invalidFragments = invalidFragments[:len(invalidFragments)-1]
		}
	}

	if len(invalidFragments) > 0 {
		validation.invalidate(ctx.name, ctx.title, invalidFragments)
	}

	return validation
}

// Return the value being validated in a custom validator.
func (ctx *ValidatorContext) Value() any {
	return ctx.value
}
