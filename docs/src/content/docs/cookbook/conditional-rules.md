---
title: Conditional Validation Rules in Go
description: Gate Valgo validation rules in Go based on business logic, previous validity, and conditional input requirements.
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
).WhenValid(func(val *v.Validation) {
  vatRate, err := lookupVATRate(country)
  if err != nil {
    val.AddErrorMessage("vat_rate", "Error computing VAT")
    return
  }
  val.Is(v.Number(vatRate, "vat_rate").GreaterThan(0))
})
```
