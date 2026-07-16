---
title: Conditional Rules
description: Gate validations based on business logic or on previous validity.
slug: 0.7/cookbook/conditional-rules
---

## Only validate when a flag is true

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  If(isAdmin, v.Is(v.String(role, "role").EqualTo("admin")))
```

## Only continue when valid so far

```go
val := v.Is(
  v.String(email, "email").Not().Blank(),
  v.String(country, "country").Not().Blank(),
)

val.When(val.Valid(), func(val *v.Validation) {
  vatRate, err := lookupVATRate(country)
  if err != nil {
    val.AddErrorMessage("vat_rate", "Error computing VAT")
    return
  }
  val.Is(v.Number(vatRate, "vat_rate").GreaterThan(0))
})
```
