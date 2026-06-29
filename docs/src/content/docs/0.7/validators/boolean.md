---
title: Boolean
description: Validate bool and *bool values.
slug: 0.7/validators/boolean
---

```go
v.Is(v.Bool(true, "active").True())
v.Is(v.Bool(false, "active").Not().False())
```

Pointer variant:

```go
var b *bool
v.Is(v.BoolP(b, "active").Nil())
```
