---
title: Typed & Any
description: Validate with Typed validators or fall back to Any for dynamic values.
slug: 0.7/validators/typed-any
---

Use typed validators when you want explicit rule sets for a domain type.

Use `Any` when your value type is dynamic and you want to run custom predicates.

```go
v.Is(v.Any(value, "payload").Passing(func(v any) bool {
  return v != nil
}))
```
