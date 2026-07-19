---
title: Go Validation Errors and JSON Output
description: Convert Valgo validation sessions to Go errors, inspect messages, and customize structured JSON validation output.
---

## ToError() and ToValgoError()

- `ToError()` returns a standard `error`.
- `ToValgoError()` returns `*valgo.Error` for structured access.

```go
val := v.Is(v.String("", "name").Not().Blank())

if err := val.ToError(); err != nil {
  return err
}
```

```go
if errInfo := val.ToValgoError(); errInfo != nil {
  for field, valueErr := range errInfo.Errors() {
    _ = field
    _ = valueErr.Messages()
  }
}
```

## AddErrorMessage(name, message)

Attach an error without a validator (useful for "entity-level" errors).

```go
val := v.Is(
  v.String(a.City, "city").Not().Blank(),
  v.String(a.Street, "street").Not().Blank(),
)

if !val.Valid() {
  val.AddErrorMessage("address", "The address is wrong!")
}
```

## Merge(other)

Merge two sessions to reuse validation logic.

```go
val := v.Is(
  v.String(r.Name, "name").Not().Blank(),
  v.String(r.Status, "status").Not().Blank(),
)

val.Merge(validatePreStatus(r.Status))
```

## Custom JSON output (Factory)

If you need a different JSON shape, use a `Factory` with a custom marshal function. See `Localization & Factory`.
