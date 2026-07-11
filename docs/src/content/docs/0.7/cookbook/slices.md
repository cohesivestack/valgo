---
title: Slices & Indexed Errors
description: Validate slices of structs and scalar lists with InRow() and InCell().
slug: 0.7/cookbook/slices
---

## Slice of structs

```go
val := v.New()
for i, a := range p.Addresses {
  val.InRow("addresses", i, v.Is(
    v.String(a.Name, "name").Not().Blank(),
    v.String(a.Street, "street").Not().Blank(),
  ))
}
```

## Slice of strings

```go
val := v.New()
for i, tag := range tags {
  val.InCell("tags", i, v.Is(
    v.String(tag, "name").Not().Blank(),
  ))
}
```
