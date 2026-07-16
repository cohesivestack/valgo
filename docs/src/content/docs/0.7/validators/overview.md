---
title: Overview
description: How validators work, including names, titles, Not(), and rule evaluation.
slug: 0.7/validators/overview
---

A validator stores one value and a chain of rules. Pass validators to `Is()` or
`Check()` to evaluate them and record failures in a `Validation` session.

## Name and title

Constructors accept a value followed by an optional error path and title:

```go
generated := v.Is(v.String("").Not().Empty())
named := v.Is(v.String("", "company_name").Not().Empty())
titled := v.Is(v.String("", "company_name", "Customer").Not().Empty())
```

Without a name, Valgo generates `value_0`, `value_1`, and so on. Without a
title, it humanizes the name for messages. `New()` creates an empty session and
accepts only `Options`; it does not accept validators.

## Not

`Not()` inverts only the next rule and then resets.

```go
valid := v.Is(v.Number(1).Not().Zero()).Valid() // true
```

Rules are implicitly joined with AND. `Is()` short-circuits after a failure,
whereas `Check()` continues evaluating rules.
