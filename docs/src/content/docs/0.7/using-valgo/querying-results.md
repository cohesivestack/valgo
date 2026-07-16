---
title: Querying Results
description: Inspect recorded validation results with IsValid().
slug: 0.7/using-valgo/querying-results
---

`IsValid(path)` does not run validation. It checks whether the path was marked
invalid in the current `Validation` session.

```go
val := v.Is(
  v.String("", "email").Not().Empty(),
  v.String("John", "name").Not().Blank(),
)

_ = val.IsValid("email") // false
_ = val.IsValid("name")  // true
```

Invalid nested and indexed paths also invalidate their parent namespaces. For
example, an error at `person.addresses[0].line1` makes all of these return
`false`:

```go
val.IsValid("person")
val.IsValid("person.addresses")
val.IsValid("person.addresses[0]")
val.IsValid("person.addresses[0].line1")
```

A path that was never marked invalid returns `true`, including an unknown or
never-validated path.

`PathValid()`, `AllValid()`, and `AnyValid()` were added after v0.7 and are not
available in this version.
