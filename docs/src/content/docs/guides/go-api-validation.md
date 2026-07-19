---
title: Go API Validation
description: Validate Go API request payloads with Valgo and return structured field errors for JSON clients.
---

API validation usually needs more than a boolean result. Clients need field paths, readable messages, and predictable JSON output. Valgo builds those results from validation sessions.

## Validate a request payload

Use `Is(...)` when you want each validator chain to stop after the first failed rule:

```go
val := v.Is(
  v.String(input.Email, "email").Not().Blank().Email(),
  v.String(input.Password, "password").Not().Blank().LengthBetween(8, 72),
)
```

Use `Check(...)` when you want to collect multiple messages from the same chain:

```go
val := v.Check(
  v.String(input.Password, "password").Not().Blank().LengthBetween(8, 72),
)
```

## Return structured errors

Valgo can convert a validation session to an error, inspect messages, or produce structured output for JSON responses.

```go
if err := val.ToError(); err != nil {
  return err
}
```

See [Go Validation Errors and JSON Output](/using-valgo/errors/) for output options and [Go Sign-up Form Validation](/cookbook/signup-form/) for a complete request-style example.

## Nested API payloads

For nested JSON payloads, use namespaces to create stable paths:

```go
val := v.Is(
  v.In("address",
    v.String(input.Address.City, "city").Not().Blank(),
  ),
)
```

For arrays or repeated fields, see [Validate Go Slices and Indexed Errors](/cookbook/slices/).
