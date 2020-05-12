package valgo

type DataType int

const (
	DataTypeString DataType = iota + 1
	DataTypeInteger
	DataTypeBoolean
)

type Validator interface {
	IsString(value string, nameAndTitle ...string) *StringValidator
	CheckString(value string, nameAndTitle ...string) *StringValidator
	Is(value interface{}, nameAndTitle ...string) *GenericValidator
	Check(value interface{}, nameAndTitle ...string) *GenericValidator
	Valid() bool
	Error() error
	Errors() map[string]*valueError
	AddErrorMessage(name string, message string) Validator
}
