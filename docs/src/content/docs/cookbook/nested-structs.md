---
title: Nested Structs
description: Validate nested models with namespaces using In().
---

```go
val := v.In("person",
  v.Is(
    v.String(p.Name, "name").Not().Blank(),
  ),
).In("address",
  v.Is(
    v.String(p.Address.Line1, "line1").Not().Blank(),
    v.String(p.Address.City, "city").Not().Blank(),
  ),
)
```

This produces paths like `person.name`, `address.city`.
