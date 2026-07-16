---
title: Comparable
description: Validate comparable Go values with equality, membership, and custom predicates.
slug: 0.7/validators/comparable
---

`Comparable()` accepts any type satisfying Go's `comparable` constraint. It
provides `EqualTo()`, `InSlice()`, and `Passing()`; it does not provide ordering
rules.

```go
type Point struct {
  X int
  Y int
}

point := Point{X: 1, Y: 2}
v.Is(v.Comparable(point, "point").EqualTo(Point{X: 1, Y: 2}))
v.Is(v.Comparable(point, "point").InSlice([]Point{
  {X: 0, Y: 0},
  {X: 1, Y: 2},
}))
```

For ordered numeric or string comparisons, use the matching numeric or string
validator.

Pointer values use `ComparableP()`, which also provides `Nil()`:

```go
var point *Point
v.Is(v.ComparableP(point, "point").Nil())
```
