---
title: Numbers
description: Validate Number, Int/Uint/Float families and their pointer variants.
slug: 0.7/validators/numbers
---

Valgo provides numeric validators for:

* `Number` (generic number family)
* `Int`, `Uint`, `Float` (family-specific)

## Number

```go
v.Is(v.Number(10).EqualTo(10))
v.Is(v.Number(11).GreaterThan(10))
v.Is(v.Number(11).Between(10, 12))
v.Is(v.Number(0).Zero())
v.Is(v.Number(20).InSlice([]int{10, 20, 30}))
```

## Int / Uint / Float

```go
v.Is(v.Int(int16(10)).GreaterThan(int16(5)))
v.Is(v.Uint(uint64(10)).LessOrEqualTo(uint64(10)))
v.Is(v.Float(float32(1.5)).GreaterThan(float32(1.0)))
```

## Pointer variants

```go
var x *int
v.Is(v.NumberP(x).Nil())
```
