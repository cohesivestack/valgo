---
slug: 0.8/validators/time
title: Time
description: Validate time.Time values and pointers with equality, ordering, ranges, and nil checks.
---

Use `Time()` for `time.Time` values. Time ordering uses `After()`,
`AfterOrEqualTo()`, `Before()`, and `BeforeOrEqualTo()`.

```go
v.Is(v.Time(start, "start").Before(end))
v.Is(v.Time(end, "end").AfterOrEqualTo(start))
v.Is(v.Time(now, "created_at").Between(start, end)) // inclusive
```

Other rules are `EqualTo()`, `Zero()`, `InSlice()`, and `Passing()`.

## Pointer variant

`TimeP()` accepts `*time.Time` and adds `Nil()` and `NilOrZero()`.

```go
var expiresAt *time.Time
v.Is(v.TimeP(expiresAt, "expires_at").Nil())
```
