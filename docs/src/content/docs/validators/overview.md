---
title: Go Validators Overview
description: Understand how Valgo validators work in Go, including names, titles, Not(), rule evaluation, and validation chains.
---

A validator stores one value and a chain of rules. Pass validators to `Is()` or
`Check()` to evaluate them and record any failures in a `Validation` session.

## Name and title

Validator constructors accept the value followed by an optional error path and
title:

```go
v.String(value, "company_name", "Company name")
```

If the name is omitted, Valgo generates `value_0`, `value_1`, and so on within
the session. If a name is supplied without a title, Valgo humanizes the name
for messages: `company_name` becomes `Company name`.

```go
generated := v.Is(v.String("").Not().Empty())
named := v.Is(v.String("", "company_name").Not().Empty())
titled := v.Is(v.String("", "company_name", "Customer").Not().Empty())
```

`New()` creates an empty validation session and accepts only `Options`. Use
`Is()` or `Check()` to evaluate validators.

## Not

`Not()` inverts only the next rule and then resets.

```go
valid := v.Is(v.Number(1).Not().Zero()).Valid() // true
```

## Rule evaluation

Rules are implicitly joined with AND. `Is()` stops a validator chain after a
failure, while `Check()` continues collecting failures. Use `Or()` and
`OrElse()` for alternatives; their grouping and short-circuit behavior is
documented on the OR Operators page.
