---
title: Optional Field Validation in Go
description: Use Valgo pointer validators to validate optional Go fields, nil values, and present-but-invalid input.
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
