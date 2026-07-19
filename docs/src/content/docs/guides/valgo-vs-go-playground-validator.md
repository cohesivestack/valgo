---
title: Valgo vs go-playground/validator
description: Compare Valgo and go-playground/validator for Go validation, including struct tags, type safety, conditional rules, and error output.
---

Valgo and `go-playground/validator` both help Go applications validate input, but they use different design styles.

`go-playground/validator` is widely used and centers validation rules around struct tags. Valgo centers validation around typed Go functions and rule chains.

## Main difference

With tag-based validation, rules usually live in struct field tags:

```go
type Signup struct {
  Email string `validate:"required,email"`
}
```

With Valgo, rules live in Go code:

```go
val := v.Is(
  v.String(input.Email, "email").Not().Blank().Email(),
)
```

That difference affects how you compose rules, test validation logic, and express validation that depends on context.

## When Valgo fits well

Valgo is a good fit when you want:

- validation without struct tags
- type-safe rule constructors
- validation of arbitrary values, not only struct fields
- conditional flows written in Go
- structured error paths for nested objects and slices
- localized validation messages

## When tag-based validation fits well

Tag-based validation can be a good fit when validation rules are simple, static, and tightly coupled to DTO struct definitions. It keeps simple required-field rules compact.

Valgo is more explicit. That costs a few more lines, but it makes validation logic easier to compose when rules vary by workflow.

## Related Valgo docs

- [Go Validation Without Struct Tags](/guides/go-validation-without-struct-tags/)
- [Conditional Validation in Go](/using-valgo/conditional-flows/)
- [Go Validation Errors and JSON Output](/using-valgo/errors/)
- [Localized Go Validation Messages](/using-valgo/localization/)
