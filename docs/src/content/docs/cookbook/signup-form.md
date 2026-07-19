---
title: Go Sign-up Form Validation
description: Validate a Go sign-up form payload with Valgo and return structured field errors for API responses.
---

```go
val := v.Is(
  v.String(email, "email").Not().Blank(),
  v.String(password, "password").Not().Blank().MinBytes(8),
  v.Number(age, "age").GreaterOrEqualTo(18),
)

if err := val.ToError(); err != nil {
  // Marshal err to JSON and return it from your handler.
  return err
}
```

If you need all messages for a field, use `Check(...)`.
