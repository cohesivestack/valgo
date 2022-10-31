package valgo

// Interface implemented by valgo Validators and custom Validators.
type Validator interface {
	Context() *ValidatorContext
}
