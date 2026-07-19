---
slug: 0.8/validators/numbers
title: Numbers
description: Validate numeric values and pointers with generic and family-specific validators.
---

Valgo provides a generic validator for all supported numeric types and
family-specific constructors that preserve the exact input type.

## Number

`Number()` accepts any value whose underlying type is a signed integer,
unsigned integer, `float32`, or `float64`.

```go
v.Is(v.Number(10).EqualTo(10))
v.Is(v.Number(11).GreaterThan(10))
v.Is(v.Number(11).Between(10, 12)) // inclusive
v.Is(v.Number(0).Zero())
v.Is(v.Number(20).InSlice([]int{10, 20, 30}))
```

`Number()` provides equality, ordering, inclusive `Between()`, `Zero()`,
`InSlice()`, and `Passing()`. Use `NumberP()` for pointers; it additionally
provides `Nil()` and `ZeroOrNil()`.

```go
var amount *int
v.Is(v.NumberP(amount, "amount").Nil())
```

## Signed integers

Use the constructor matching the value's underlying type: `Int()`, `Int8()`,
`Int16()`, `Int32()`, `Int64()`, or `Rune()`. The pointer forms add a `P`
suffix, such as `Int16P()` and `RuneP()`.

```go
v.Is(v.Int(10).Positive())
v.Is(v.Int16(int16(10)).GreaterThan(int16(5)))
v.Is(v.Int64(int64(-1)).Negative())
```

Signed integer validators add `Positive()` and `Negative()` to the common
numeric rules.

## Unsigned integers

Use `Uint()`, `Uint8()`, `Uint16()`, `Uint32()`, `Uint64()`, or `Byte()`, with
matching pointer constructors such as `Uint64P()` and `ByteP()`.

```go
v.Is(v.Uint64(uint64(10)).LessOrEqualTo(uint64(10)))
v.Is(v.Byte(byte(1)).GreaterThan(byte(0)))
```

## Floating-point numbers

Use `Float32()` or `Float64()` and their pointer forms `Float32P()` and
`Float64P()`.

```go
v.Is(v.Float32(float32(1.5)).GreaterThan(float32(1.0)))
v.Is(v.Float64(3.14).Finite())
```

Float validators add `Positive()`, `Negative()`, `NaN()`, `Infinite()`, and
`Finite()` to the common numeric rules.
