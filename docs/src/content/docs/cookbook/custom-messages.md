---
title: Custom Go Validation Messages
description: Override Valgo validation message templates in Go for single validator calls and user-facing API errors.
---

Most rules accept an optional template argument.

```go
val := v.Is(
  v.String("", "address_field", "Address").
    Not().
    Empty("{{title}} must not be empty. Please provide the value in the input {{name}}."),
)
```
