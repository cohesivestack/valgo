package valgo

// Validator Interface implemented by valgo Validators and custom Validators.
type Validator interface {
	Context() *ValidatorContext
}
