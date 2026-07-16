---
title: Comparable
description: Validate comparable Go values with equality, membership, and custom predicates.
---

`Comparable()` accepts any type satisfying Go's `comparable` constraint,
including strings, numbers, pointers, and structs whose fields are all
comparable. It does not provide ordering rules.

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
v.Is(v.Comparable(point, "point").Passing(func(p Point) bool {
  return p.X >= 0 && p.Y >= 0
}))
```

For ordered numeric or string comparisons, use the corresponding numeric or
string validator.

## Pointer variant

`ComparableP()` validates a pointer to a comparable value. Its callback
receives the pointer, and it also provides `Nil()`.

```go
var point *Point
v.Is(v.ComparableP(point, "point").Nil())
```
