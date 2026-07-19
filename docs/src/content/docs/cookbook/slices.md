---
title: Validate Go Slices and Indexed Errors
description: Validate Go slices, lists, and repeated payloads with Valgo InRow(), InCell(), and indexed validation errors.
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
