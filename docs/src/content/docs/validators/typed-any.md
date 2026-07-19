---
title: Typed and Any Validators for Go
description: Validate arbitrary Go values with Valgo typed validators, dynamic predicates, equality checks, nil checks, and custom rules.
---

## Typed

`Typed()` preserves the compile-time type of any value. `Passing()` therefore
receives a callback with that exact type. It also provides `Nil()`.

```go
type Status string

status := Status("running")
val := v.Is(v.Typed(status, "status").Passing(func(status Status) bool {
  return status == "running" || status == "paused"
}))
```

## Any

`Any()` stores a dynamic value. It provides `EqualTo()`, `Nil()`, and a
`Passing()` callback whose argument type is `any`.

```go
val := v.Is(v.Any(value, "payload").Passing(func(value any) bool {
  return value != nil
}))
```

Use `Comparable()` instead when the type satisfies `comparable` and you need
type-safe equality or slice membership.
