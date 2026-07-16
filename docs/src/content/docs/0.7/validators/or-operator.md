---
title: OR Operator
description: Combine validation rules with Or() in Valgo v0.7.
slug: 0.7/validators/or-operator
---

`Or()` takes no arguments. It joins the rule before it and the rule after it as
alternatives in the same validator chain.

```go
val := v.Is(
  v.String(contact, "contact").
    MatchingTo(emailRe).
    Or().
    MatchingTo(phoneRe),
)
```

Adjacent alternatives form an OR group. The group is evaluated before the
implicit AND rules that precede or follow it:

```text
A.Or().B.C  == (A OR B) AND C
A.B.Or().C  == A AND (B OR C)
```

A failure before an OR group begins cannot be rescued by a later `Or()`. When
the left side of `Or()` succeeds, its alternative is skipped, but any implicit
AND rules after the group are still evaluated.

`OrElse()` is not available in v0.7.
