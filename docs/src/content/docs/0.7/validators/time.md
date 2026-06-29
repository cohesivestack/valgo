---
title: Time
description: Validate time.Time and *time.Time (ordering, ranges, nil checks).
slug: 0.7/validators/time
---

Use `v.Time(t, ...)` and `v.TimeP(&t, ...)`.

```go
v.Is(v.Time(start, "start").LessThan(end))
v.Is(v.Time(end, "end").GreaterOrEqualTo(start))
```

```go
var t *time.Time
v.Is(v.TimeP(t, "expires_at").Nil())
```
