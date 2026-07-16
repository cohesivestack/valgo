---
title: Conditional Flows
description: Use If(), When(), and Do() to build conditional validation flows.
slug: 0.7/using-valgo/conditional-flows
---

## If(condition, validation)

`If()` merges the supplied validation result only when the condition is true.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  If(isAdmin, v.Is(v.String(role, "role").EqualTo("admin")))
```

Go evaluates the supplied validation before calling `If()`. Only its merge is
conditional. Use `When()` to defer validation work.

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

The `If*Valid` and `When*Valid` helper families were added after v0.7 and are
not available in this version.
