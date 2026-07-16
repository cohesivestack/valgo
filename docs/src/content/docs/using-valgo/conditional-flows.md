---
title: Conditional Flows
description: Conditionally merge validations or execute callbacks in a Validation chain.
---

All conditional methods return the same `*Validation`, so they can be used in a
fluent chain.

## If(condition, validation)

`If()` merges the supplied validation result only when the condition is true.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  If(isAdmin, v.Is(v.String(role, "role").EqualTo("admin")))
```

Go evaluates function arguments before calling `If()`. The supplied validation
is therefore built in either case; only the merge is conditional. Use `When()`
when creating the validation is expensive or has side effects.

## Do(function)

`Do()` always calls the function with the current validation session.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  Do(func(val *v.Validation) {
    if isAdmin {
      val.Is(v.String(role, "role").EqualTo("admin"))
    }
  })
```

## When(condition, function)

`When()` calls the function only when the condition is true.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  When(isAdmin, func(val *v.Validation) {
    val.Is(v.String(role, "role").EqualTo("admin"))
  })
```

## Validity-gated merges

These methods inspect results already recorded in the current session and
merge another validation when the predicate passes:

```go
val.IfValid(other)
val.IfPathValid("email", other)
val.IfAllValid([]string{"email", "password"}, other)
val.IfAnyValid([]string{"email", "phone"}, other)
```

Like `If()`, these methods receive an already-created validation. In
particular, `IfAllValid` and `IfAnyValid` take a `[]string`, not variadic path
arguments.

An empty path slice makes `IfAllValid` use the overall `Valid()` result. An
empty path slice makes `IfAnyValid` a no-op.

## Validity-gated callbacks

The corresponding `When` methods execute a callback only when the predicate
passes:

```go
val.WhenValid(func(val *v.Validation) { /* ... */ })
val.WhenPathValid("email", func(val *v.Validation) { /* ... */ })
val.WhenAllValid([]string{"email", "password"}, func(val *v.Validation) { /* ... */ })
val.WhenAnyValid([]string{"email", "phone"}, func(val *v.Validation) { /* ... */ })
```

Use these callback forms to defer expensive work. Their empty-slice behavior
matches the corresponding `IfAllValid` and `IfAnyValid` methods.
