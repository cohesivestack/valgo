---
title: Localized Validation Errors in Go
description: Build localized Go validation errors with Valgo locales, custom message templates, and reusable validation factories.
---

Valgo includes localization support for validation messages. You can switch locales, override message templates, and reuse localization options through a validation factory.

Localized errors matter when validation messages are shown directly to users. A backend can keep field names and rule names stable while returning messages in the user’s preferred language.

## Use localized messages

Valgo validation messages are generated from rule templates. You can configure those templates globally for a validation session or reuse them with a factory.

```go
factory := v.NewFactory(
  v.WithLocale("es"),
)

val := factory.Is(
  v.String(input.Name, "name").Not().Blank(),
)
```

## Keep output predictable

The validation result can still use stable field paths while the message text changes by locale. That is useful for APIs where clients depend on field keys like `email`, `profile.name`, or `items[0]`.

For details, see [Localized Go Validation Messages](/using-valgo/localization/) and the [Localization cookbook](/cookbook/localization/).
