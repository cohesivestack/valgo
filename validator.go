package valgo

import (
	"strconv"
)

type DataType int

const (
	DataTypeString DataType = iota + 1
	DataTypeInteger
	DataTypeBoolean
)

type Validator struct {
	currentValue    interface{}
	currentTitle    *string
	currentName     *string
	currentValid    bool
	currentIndex    int
	currentDataType DataType
	currentNegative bool

	_locale *locale
	valid   bool
	errors  map[string]*valueError
}

func (v *Validator) IsString(value string, nameAndTitle ...string) *Validator {
	v.currentDataType = DataTypeString
	return v.Is(value, nameAndTitle...)
}

func (v *Validator) Is(value interface{}, nameAndTitle ...string) *Validator {
	v.currentValue = value
	v.currentIndex += 1
	v.currentValid = true

	sizeNameAndTitle := len(nameAndTitle)
	if sizeNameAndTitle > 0 {
		v.currentName = &nameAndTitle[0]
		if sizeNameAndTitle > 1 {
			v.currentTitle = &nameAndTitle[1]
		}
	}
	return v
}

func (v *Validator) Valid() bool {
	return v.valid
}

func (v *Validator) Errors() map[string]*valueError {
	return v.errors
}

func (v *Validator) Not() *Validator {
	v.currentNegative = true

	return v
}

func (v *Validator) assert(value bool) bool {
	return v.currentNegative != value
}

func (v *Validator) resetNegative() {
	v.currentNegative = false
}

func (v *Validator) invalidate(errorKey string, values map[string]interface{}, template ...string) {
	if v.errors == nil {
		v.errors = map[string]*valueError{}
	}

	v.currentValid = false
	v.valid = false

	var name string
	if v.currentName == nil {
		name = concatString("value_", strconv.Itoa(v.currentIndex-1))
	} else {
		name = *v.currentName
	}

	if _, ok := v.errors[name]; !ok {
		v.errors[name] = &valueError{
			name:          &name,
			errorMessages: map[string]*errorMessage{},
			validator:     v,
		}
	}

	ev := v.errors[name]
	ev.dirty = true

	if v.currentNegative {
		errorKey = concatString("not_", errorKey)
	}

	if _, ok := ev.errorMessages[errorKey]; !ok {
		ev.errorMessages[errorKey] = &errorMessage{
			key: errorKey,
		}
	}

	em := ev.errorMessages[errorKey]
	if len(template) > 0 {
		em.template = &template[0]
	}
	em.values = values
}
