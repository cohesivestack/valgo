---
title: Go Validation Without Struct Tags
description: Learn how to validate Go values without struct tags using Valgo function-based validation chains.
---

Valgo validates Go values with regular functions instead of struct tags. This is useful when validation rules depend on runtime context, when you want reusable rule chains, or when the value being validated is not a struct field.

Struct tags are compact, but they put validation rules in string literals. Valgo keeps validation in Go code, so your editor, compiler, and tests can work with rule constructors directly.

## Why avoid struct tags

Function-based validation is a good fit when validation should be:

- type-safe
- composable
- conditional
- reusable across structs, API handlers, and services
- localized with custom message templates

Valgo validators accept the value, a machine-readable field name, and then a chain of rules:

```go
val := v.Is(
  v.String(input.Email, "email").Not().Blank().Email(),
  v.String(input.Password, "password").Not().Blank().LengthBetween(8, 72),
)

if err := val.ToError(); err != nil {
  return err
}
```

## Validate values where they are used

Because rules are regular Go calls, you can keep validation close to the workflow that needs it. A public API request, an internal command, and a background job can validate the same field differently without requiring multiple struct definitions or tag variants.

For a first example, start with [Getting Started](/getting-started/). For reusable validation functions, see [Reusable Go Validation Functions](/cookbook/reusable-validations/).

## Validate nested input

Valgo uses namespaces to build structured error paths for nested structs and slices. That keeps validation errors useful for JSON APIs and form clients.

```go
val := v.Is(
  v.In("profile",
    v.String(input.Profile.Name, "name").Not().Blank(),
  ),
)
```

For deeper examples, see [Validate Nested Structs in Go](/cookbook/nested-structs/) and [Validate Go Slices and Indexed Errors](/cookbook/slices/).
