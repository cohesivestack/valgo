package valgo

type Validator interface {
	Context() *ValidatorContext
}
