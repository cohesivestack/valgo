---
title: Querying Results
description: Inspect the last validation outcome with PathValid(), AllValid(),
  and AnyValid().
slug: 0.7/using-valgo/querying-results
---

These helpers do not re-run validation. They inspect the latest recorded errors.

## PathValid(path)

```go
val := v.Is(
  v.String("", "email").Not().Empty(),
  v.String("John", "name").Not().Blank(),
)

_ = val.PathValid("email") // false
_ = val.PathValid("name")  // true
```

For nested/indexed paths, parent namespaces are considered invalid when any child is invalid.

## AllValid(paths...)

```go
val := v.Is(
  v.String("john@example.com", "email").Not().Empty(),
  v.String("", "password").Not().Empty(),
)

_ = val.AllValid("email")             // true
_ = val.AllValid("email", "password") // false
_ = val.AllValid()                     // same as val.Valid()
```

## AnyValid(paths...)

```go
val := v.Is(
  v.String("", "email").Not().Empty(),
  v.String("+3569999999", "phone").Not().Empty(),
)

_ = val.AnyValid("email", "phone") // true
_ = val.AnyValid()                   // false (explicit set required)
```

## IsValid(name) (deprecated)

`IsValid("path")` is deprecated; use `PathValid("path")`.
