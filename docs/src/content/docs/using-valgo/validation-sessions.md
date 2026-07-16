---
title: Validation Sessions
description: The Validation session is the core type for validating one or many values.
---

## The Validation session

A `Validation` session holds one or more validators. You typically build one with:

- `Is(...)`
- `Check(...)`
- `New()`

## New()

`New()` creates an empty session you can populate conditionally. Its optional
argument is an `Options` value; it does not accept validators.

```go
val := v.New()

if month == 6 {
  val.Is(v.Number(day, "month_day").LessOrEqualTo(10))
}

if err := val.ToError(); err != nil {
  return err
}
```

## Valid()

`Valid()` returns `true` when the session has no validation errors and `false`
when at least one error was recorded.

```go
val := v.Is(v.String("", "name").Not().Blank())
_ = val.Valid() // false
```

## Is() chaining

`Validation.Is(...)` appends validators to an existing session:

```go
val := v.Is(v.String("john", "username").Not().Blank())
val.Is(v.String("single", "status").InSlice([]string{"single", "married"}))
```
