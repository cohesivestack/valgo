package valgo

type DataType int

const (
	DataTypeString DataType = iota + 1
	DataTypeInteger
	DataTypeInt64
	DataTypeBoolean
)

type Validator interface {
	IsString(value string, nameAndTitle ...string) *StringValidator
	CheckString(value string, nameAndTitle ...string) *StringValidator
	IsInt64(value int64, nameAndTitle ...string) *Int64Validator
	CheckInt64(value int64, nameAndTitle ...string) *Int64Validator
	Is(value interface{}, nameAndTitle ...string) *GenericValidator
	Check(value interface{}, nameAndTitle ...string) *GenericValidator
	Valid() bool
	Error() error
	Errors() map[string]*valueError
	AddErrorMessage(name string, message string) Validator
	IsValid(name string) bool
}
