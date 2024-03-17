package valgo

type validatorFragment struct {
	errorKey       string
	template       []string
	templateParams map[string]any
	function       func() bool
	boolOperation  bool
	orOperation    bool
	isValid        bool
}

// The context keeps the state and provides the functions to control a
// custom validator.
type ValidatorContext struct {
	fragments     []*validatorFragment
	value         any
	name          *string
	title         *string
	boolOperation bool
	orOperation   bool
}

// Create a new [ValidatorContext] to be used by a custom validator.
func NewContext(value any, nameAndTitle ...string) *ValidatorContext {

	context := &ValidatorContext{
		value:         value,
		fragments:     []*validatorFragment{},
		boolOperation: true,
		orOperation:   false,
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
		ctx.orOperation = true
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
	ctx.orOperation = false

	return ctx
}

func (ctx *ValidatorContext) validateIs(validation *Validation) *Validation {
	return ctx.validate(validation, true)
}

func (ctx *ValidatorContext) validateCheck(validation *Validation) *Validation {
	return ctx.validate(validation, false)
}

func (ctx *ValidatorContext) validate(validation *Validation, shortCircuit bool) *Validation {
	// valid := true
	validation.currentIndex++

	// Iterating through each fragment in the context's fragment list
	for i, fragment := range ctx.fragments {

		// If the previous fragment is not valid, the current fragment is not in an "or" operation, and the short circuit flag is true,
		// we return the current state of the validation without evaluating the current fragment
		if i > 0 && !ctx.fragments[i-1].isValid && !fragment.orOperation && shortCircuit {
			break
		}

		// If the current fragment is a part of an "or" operation and the previous fragment in the "or" operation
		// is valid, we mark the current fragment as valid and move to the next iteration
		if fragment.orOperation && ctx.fragments[i-1].isValid {
			continue
		}

		// Evaluating the validation function of the current fragment and updating the valid flag
		// The valid flag will be true only if the fragment function returns a value matching the fragment's boolean operation
		// and the valid flag was true before this evaluation
		fragment.isValid = fragment.function() == fragment.boolOperation

		// If the current fragment is valid and is part of an "or" operation, we backtrack to mark all preceding
		// fragments in the "or" operation chain as valid
		if fragment.isValid && fragment.orOperation {
			for j := i - 1; j >= 0; j-- {
				ctx.fragments[j].isValid = true
				// Breaking the loop when we reach the start of the "or" operation chain
				if !ctx.fragments[j].orOperation {
					break
				}
			}
		}

		// Setting the validation state of the current fragment
		// valid = fragment.isValid && valid
	}

	for _, fragment := range ctx.fragments {
		if !fragment.isValid {
			validation.invalidate(ctx.name, ctx.title, fragment)
		}
	}

	return validation
}

// Return the value being validated in a custom validator.
func (ctx *ValidatorContext) Value() any {
	return ctx.value
}
