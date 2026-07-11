---
title: Custom Messages
description: Override a rule template for a single validator call.
slug: 0.7/cookbook/custom-messages
---

Most rules accept an optional template argument.

```go
val := v.Is(
  v.String("", "address_field", "Address").
    Not().
    Empty("{{title}} must not be empty. Please provide the value in the input {{name}}."),
)
```
