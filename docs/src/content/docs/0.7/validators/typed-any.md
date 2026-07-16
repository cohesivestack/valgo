---
title: Typed & Any
description: Validate arbitrary values with typed or dynamic predicates, equality, and nil checks.
slug: 0.7/validators/typed-any
---

`Typed()` preserves the value's compile-time type, so `Passing()` receives a
callback with that exact type. It also provides `Nil()`.

```go
type Status string
status := Status("running")

v.Is(v.Typed(status, "status").Passing(func(status Status) bool {
  return status == "running" || status == "paused"
}))
```

`Any()` provides `EqualTo()`, `Nil()`, and a dynamic `Passing()` callback:

```go
v.Is(v.Any(value, "payload").Passing(func(value any) bool {
  return value != nil
}))
```
