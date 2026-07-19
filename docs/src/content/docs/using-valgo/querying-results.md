---
title: Query Go Validation Results
description: Inspect Valgo validation results in Go with PathValid(), AllValid(), AnyValid(), and structured field checks.
---

These helpers do not run validation. They query the invalid paths recorded in
the current `Validation` session.

A path is considered valid when it has not been recorded as invalid. As a
result, an unknown path or a path that was never validated returns `true` from
`PathValid()` and can make `AnyValid()` return `true`.

## PathValid(path)

```go
val := v.Is(
  v.String("", "email").Not().Empty(),
  v.String("John", "name").Not().Blank(),
)

_ = val.PathValid("email") // false
_ = val.PathValid("name")  // true
```

For nested/indexed paths, parent namespaces are considered invalid when any child is invalid.

## AllValid(paths...)

```go
val := v.Is(
  v.String("john@example.com", "email").Not().Empty(),
  v.String("", "password").Not().Empty(),
)

_ = val.AllValid("email")             // true
_ = val.AllValid("email", "password") // false
_ = val.AllValid()                     // same as val.Valid()
```

## AnyValid(paths...)

```go
val := v.Is(
  v.String("", "email").Not().Empty(),
  v.String("+3569999999", "phone").Not().Empty(),
)

_ = val.AnyValid("email", "phone") // true
_ = val.AnyValid()                  // false (explicit set required)
```

## IsValid(name) (deprecated)

`IsValid("path")` is deprecated; use `PathValid("path")`.
