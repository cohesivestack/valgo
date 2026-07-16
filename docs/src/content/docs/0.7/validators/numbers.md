---
title: Numbers
description: Validate numeric values and pointers with generic and family-specific validators.
slug: 0.7/validators/numbers
---

## Number

`Number()` accepts signed integers, unsigned integers, `float32`, `float64`,
and custom types based on them.

```go
v.Is(v.Number(10).EqualTo(10))
v.Is(v.Number(11).GreaterThan(10))
v.Is(v.Number(11).Between(10, 12)) // inclusive
v.Is(v.Number(0).Zero())
v.Is(v.Number(20).InSlice([]int{10, 20, 30}))
```

Use `NumberP()` for pointers; it also provides `Nil()` and `ZeroOrNil()`.

## Signed integers

Use `Int()`, `Int8()`, `Int16()`, `Int32()`, `Int64()`, or `Rune()`, with
pointer forms such as `Int16P()`. These validators add `Positive()` and
`Negative()`.

```go
v.Is(v.Int(10).Positive())
v.Is(v.Int16(int16(10)).GreaterThan(int16(5)))
```

## Unsigned integers

Use `Uint()`, `Uint8()`, `Uint16()`, `Uint32()`, `Uint64()`, or `Byte()`, with
pointer forms such as `Uint64P()`.

```go
v.Is(v.Uint64(uint64(10)).LessOrEqualTo(uint64(10)))
```

## Floating-point numbers

Use `Float32()` or `Float64()` and their pointer forms. Float validators add
`Positive()`, `Negative()`, `NaN()`, `Infinite()`, and `Finite()`.

```go
v.Is(v.Float32(float32(1.5)).GreaterThan(float32(1.0)))
v.Is(v.Float64(3.14).Finite())
```
