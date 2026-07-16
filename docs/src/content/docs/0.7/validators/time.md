---
title: Time
description: Validate time.Time values and pointers with equality, ordering, ranges, and nil checks.
slug: 0.7/validators/time
---

Time ordering uses `After()`, `AfterOrEqualTo()`, `Before()`, and
`BeforeOrEqualTo()`.

```go
v.Is(v.Time(start, "start").Before(end))
v.Is(v.Time(end, "end").AfterOrEqualTo(start))
v.Is(v.Time(now, "created_at").Between(start, end)) // inclusive
```

Other rules are `EqualTo()`, `Zero()`, `InSlice()`, and `Passing()`.

`TimeP()` accepts `*time.Time` and adds `Nil()` and `NilOrZero()`:

```go
var expiresAt *time.Time
v.Is(v.TimeP(expiresAt, "expires_at").Nil())
```
