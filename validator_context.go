package valgo

type validatorFragment struct {
	errorKey       string
	template       []string
	templateParams map[string]any
	function       func() bool
	boolOperation  bool
}

// The context keeps the state and provides the functions to control a
// custom validator.
type ValidatorContext struct {
	fragments     []*validatorFragment
	value         any
	name          *string
	title         *string
	boolOperation bool
}

// Create a new [ValidatorContext] to be used by a custom validator.
func NewContext(value any, nameAndTitle ...string) *ValidatorContext {

	context := &ValidatorContext{
		value:         value,
		fragments:     []*validatorFragment{},
		boolOperation: true,
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

// Reverse the logical value associated with the next validator function in
// a custom validator.
func (ctx *ValidatorContext) Not() *ValidatorContext {
	ctx.boolOperation = false
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
	}
	if len(template) > 0 {
		fragment.template = template
	}
	ctx.fragments = append(ctx.fragments, fragment)
	ctx.boolOperation = true

	return ctx
}

func (ctx *ValidatorContext) validateIs(validation *Validation) *Validation {
	return ctx.validate(validation, true)
}

func (ctx *ValidatorContext) validateCheck(validation *Validation) *Validation {
	return ctx.validate(validation, false)
}

func (ctx *ValidatorContext) validate(validation *Validation, shortCircuit bool) *Validation {
	validation.valid = true
	validation.currentIndex++

	for i, fragment := range ctx.fragments {
		if i > 0 && !validation.valid && shortCircuit {
			return validation
		}

		validation.valid = fragment.function() == fragment.boolOperation && validation.valid
		if !validation.valid {
			validation.invalidate(ctx.name, fragment)
		}
	}

	return validation
}

// Return the value being validated in a custom validator.
func (ctx *ValidatorContext) Value() any {
	return ctx.value
}
