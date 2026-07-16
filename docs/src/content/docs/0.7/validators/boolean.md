---
title: Boolean
description: Validate bool values and pointers with truth, equality, membership, and nil rules.
slug: 0.7/validators/boolean
---

`Bool()` supports `True()`, `False()`, `EqualTo()`, `InSlice()`, and
`Passing()`.

```go
v.Is(v.Bool(true, "active").True())
v.Is(v.Bool(false, "active").Not().False())
```

`BoolP()` accepts `*bool` and adds `Nil()` and `FalseOrNil()`:

```go
var active *bool
v.Is(v.BoolP(active, "active").FalseOrNil())
```
