---
title: Optional Fields (Pointers)
description: Use pointer validators to express optional fields and nil checks.
slug: 0.7/cookbook/optional-fields
---

If a field is optional, represent it as a pointer and validate with `*P` validators.

```go
var middleName *string

val := v.Is(
  v.StringP(middleName, "middle_name").BlankOrNil(),
)
```

For optional numbers:

```go
var discount *int
val := v.Is(v.NumberP(discount, "discount").ZeroOrNil())
```
