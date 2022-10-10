package valgo

type ValidatorFragment struct {
	errorKey       string
	template       []string
	templateParams map[string]any
	function       func() bool
	boolOperation  bool
}

type ValidatorContext struct {
	fragments     []*ValidatorFragment
	value         any
	name          *string
	title         *string
	boolOperation bool
}

func NewContext(value any, nameAndTitle ...string) *ValidatorContext {

	context := &ValidatorContext{
		value:         value,
		fragments:     []*ValidatorFragment{},
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

func (ctx *ValidatorContext) Not() *ValidatorContext {
	ctx.boolOperation = false
	return ctx
}

func (ctx *ValidatorContext) AddWithValue(function func() bool, errorKey string, value any, template ...string) *ValidatorContext {
	return ctx.AddWithParams(
		function,
		errorKey,
		map[string]any{"title": ctx.title, "value": value}, template...)
}

func (ctx *ValidatorContext) Add(function func() bool, errorKey string, template ...string) *ValidatorContext {
	return ctx.AddWithParams(
		function,
		errorKey,
		map[string]any{"title": ctx.title}, template...)
}

func (ctx *ValidatorContext) AddWithParams(function func() bool, errorKey string, templateParams map[string]any, template ...string) *ValidatorContext {

	fragment := &ValidatorFragment{
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

func (ctx *ValidatorContext) Value() any {
	return ctx.value
}
