---
title: Overview
description: How validators work, naming/title, and Not().
---

Validators hold a value and a list of rules to evaluate.

## Name and title

If you omit a name, Valgo generates one (e.g. `value_0`).

```go
val := v.New(v.String("").Empty())
```

Prefer passing a stable field name (and optionally a human-friendly title):

```go
val := v.New(v.String("", "company_name").Not().Empty())
val := v.New(v.String("", "company_name", "Customer").Not().Empty())
```

## Not()

`Not()` inverts the next rule.

```go
valid := v.Is(v.Number(0).Not().Zero()).Valid()
```
