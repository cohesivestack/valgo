---
title: Conditional Validation Rules in Go
description: Write conditional Go validation rules with Valgo when validation depends on business logic, previous results, or runtime state.
---

Real validation often depends on context. A field may be required only for a specific account type, an address may be required only when shipping physical goods, or one validation branch may depend on a previous result.

Valgo keeps those conditions in Go code instead of encoding them into tag strings.

## Gate rules with Go logic

Use ordinary Go conditions around validation calls when the condition belongs to the surrounding workflow:

```go
checks := []v.Validator{
  v.String(input.Email, "email").Not().Blank().Email(),
}

if input.RequiresCompanyName {
  checks = append(checks, v.String(input.CompanyName, "company_name").Not().Blank())
}

val := v.Is(checks...)
```

## Use validation chain flow

Valgo also supports conditional validation flows inside a validation session. This helps when a follow-up validation should run only after a previous branch is valid.

For the API surface, see [Conditional Validation in Go](/using-valgo/conditional-flows/) and [Conditional Validation Rules in Go](/cookbook/conditional-rules/).
