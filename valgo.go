package main

import (
  "fmt"
  "strconv"
)

type Validator struct {
  value interface{}
  errors []string
  valid bool
}

func (this *Validator) Valid() bool {
  return this.valid
}

func (this *Validator) Errors() []string {
  return this.errors
}

func (this *Validator) ValidAndErrors() (bool, []string) {
  return this.valid, this.errors
}

func (this *Validator) Empty() *Validator {
  value := this.ensureString()
  if (len(value) > 0) {
    this.valid = false
    this.errors = append(this.errors, "Is not empty")
  }
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

func Is(value interface{}) *Validator {
  validator := &Validator{
    value: value,
    errors: make([]string, 0, 0),
    valid: true,
  }
  return validator
}
