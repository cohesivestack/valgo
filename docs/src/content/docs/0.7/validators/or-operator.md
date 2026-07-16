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

The v0.7 implementation follows the usual AND-before-OR precedence. For
example:

```text
A.B.Or().C  == (A AND B) OR C
A.Or().B.C  == A OR (B AND C)
```

Consequently, a successful rule after `Or()` can rescue an earlier failed AND
condition. This differs from the OR-group behavior introduced after v0.7.

`OrElse()` is not available in v0.7.
