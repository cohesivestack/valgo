---
title: Conditional Flows
description: Use If(), When(), Do(), and the *Valid() helpers to build fluent conditional validation.
---

## If(condition, validation)

Merge another `Validation` only when the condition is true.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  If(isAdmin, v.Is(v.String(role, "role").EqualTo("admin")))
```

## Do(func(*Validation))

Run arbitrary logic in the chain.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  Do(func(val *v.Validation) {
    if isAdmin {
      val.Is(v.String(role, "role").EqualTo("admin"))
    }
  })
```

## When(condition, func(*Validation))

Like `Do`, but gated by a boolean.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  When(isAdmin, func(val *v.Validation) {
    val.Is(v.String(role, "role").EqualTo("admin"))
  })
```

## Validity-gated helpers

These evaluate validity from the current result and conditionally merge/execute:

- `IfPathValid`, `IfAllValid`, `IfAnyValid`, `IfValid`
- `WhenPathValid`, `WhenAllValid`, `WhenAnyValid`, `WhenValid`

Use `WhenValid` when you want to skip costly work if validation already failed.
