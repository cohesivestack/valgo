---
title: Comparable
description: Validate custom comparable types (ordering/equality) and pointers.
---

Comparable validators let you validate any ordered comparable types supported by Valgo.

```go
type Score int
v.Is(v.Comparable(Score(10), "score").GreaterThan(Score(0)))
```

Pointer variant:

```go
var s *Score
v.Is(v.ComparableP(s, "score").Nil())
```
