---
title: Namespaces
description: Group validations with In(), InRow(), and InCell() to build structured error paths.
---

Namespaces prefix field names in error output.

## In()

Use `In("address", ...)` for nested structs.

```go
val := v.
  Is(v.String(p.Name, "name").OfLengthBetween(4, 20)).
  In("address", v.Is(
    v.String(p.Address.Name, "name").Not().Blank(),
    v.String(p.Address.Street, "street").Not().Blank(),
  ))
```

Errors use dot notation (e.g. `address.name`).

## InRow()

Use `InRow("addresses", i, ...)` for a slice of structs:

```go
val := v.Is(v.String(p.Name, "name").OfLengthBetween(4, 20))
for i, a := range p.Addresses {
  val.InRow("addresses", i, v.Is(
    v.String(a.Name, "name").Not().Blank(),
    v.String(a.Street, "street").Not().Blank(),
  ))
}
```

Errors include indexes: `addresses[0].name`.

## InCell()

Use `InCell("tags", i, ...)` for a slice of scalar values:

```go
val := v.New()
for i, tag := range tags {
  val.InCell("tags", i, v.Is(
    v.String(tag, "name").Not().Blank(),
  ))
}
```

Errors target the indexed cell: `tags[0]`.
