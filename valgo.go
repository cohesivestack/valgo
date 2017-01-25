package valgo

import (
	"fmt"
	"strconv"
	"strings"
)

type Validator struct {
	value  interface{}
	errors *Errors
	valid  bool
}

func (this *Validator) Valid() bool {
	return this.valid
}

func (this *Validator) Errors() *Errors {
	if this.errors.Messages != nil && len(this.errors.Messages) > 0 {
		return this.errors
	} else {
		return nil
	}
}

func (this *Validator) ValidAndErrors() (bool, *Errors) {
	return this.valid, this.errors
}

func (this *Validator) Empty() *Validator {
	value := this.ensureString()
	if len(value) > 0 {
		this.valid = false
		this.errors.Add("Is not empty")
	}
	return this
}

func (this *Validator) NotEmpty() *Validator {
	value := this.ensureString()
	if len(value) == 0 {
		this.valid = false
		this.errors.Add("Is empty")
	}
	return this
}

func (this *Validator) Blank() *Validator {
	value := strings.Trim(this.ensureString(), " ")

	if len(value) > 0 {
		this.valid = false
		this.errors.Add("Is not blank")
	}
	return this
}

func (this *Validator) NotBlank() *Validator {
	value := strings.Trim(this.ensureString(), " ")

	if len(value) == 0 {
		this.valid = false
		this.errors.Add("Is blank")
	}
	return this
}

func Is(value interface{}) *Validator {
	validator := &Validator{
		value:  value,
		errors: &Errors{Field: "field"},
		valid:  true,
	}
	return validator
}

func (this *Validator) Called(name string) *Validator {
	this.errors.Field = name
	return this
}

func (this *Validator) ensureString() string {
	switch v := this.value.(type) {
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(this.value.(uint64), 10)
	case int8, int16, int32, int64:
		return strconv.FormatInt(this.value.(int64), 10)
	case float32, float64:
		return strconv.FormatFloat(this.value.(float64), 'f', -1, 64)
	case string:
		return this.value.(string)
	default:
		fmt.Printf("unexpected type %T", v)
		return ""
	}
}
