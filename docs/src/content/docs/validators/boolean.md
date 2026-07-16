---
title: Boolean
description: Validate bool values and pointers with truth, equality, membership, and nil rules.
---

`Bool()` supports `True()`, `False()`, `EqualTo()`, `InSlice()`, and
`Passing()`.

```go
v.Is(v.Bool(true, "active").True())
v.Is(v.Bool(false, "active").Not().False())
v.Is(v.Bool(true, "enabled").EqualTo(true))
```

## Pointer variant

`BoolP()` accepts `*bool`. Its callback receives the pointer, and it adds
`Nil()` and `FalseOrNil()`.

```go
var active *bool
v.Is(v.BoolP(active, "active").Nil())
v.Is(v.BoolP(active, "active").FalseOrNil())
```
