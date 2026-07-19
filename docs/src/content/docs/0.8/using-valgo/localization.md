---
slug: 0.8/using-valgo/localization
title: Localization & Factory
description: Localize messages, override templates, and reuse options with a ValidationFactory.
---

Valgo ships with localized messages (English default; also Spanish, German, Hungarian) and lets you provide your own locale entries.

## Set a locale by code

```go
val := v.New(v.Options{LocaleCode: "es"}).
  Check(v.String(" ", "nombre").Not().Blank())
```

## Override locale entries

```go
val := v.New(v.Options{
  Locale: &v.Locale{
    v.ErrorKeyNotBlank: "{{title}} should not be blank",
  },
}).Check(v.String(" ", "name").Not().Blank())
```

## Add a new locale

```go
val := v.New(v.Options{
  LocaleCode: "ee",
  Locale: &v.Locale{
    v.ErrorKeyNotBlank: "{{title}} ei tohi olla tuhi",
  },
}).Is(v.String(" ", "name").Not().Blank())
```

## Custom validator fallback messages

Custom validators can provide default messages for their own error keys with
`ValidatorContext.WithLocaleFallback(...)`. Fallback entries are merged only
when the active validation locale does not already define the key, so
`Options{Locale: ...}` overrides still take precedence.

## Factory()

Use a factory to set defaults (locale, custom marshaling) once.

```go
factory := v.Factory(v.FactoryOptions{
  LocaleCodeDefault: "es",
  Locales: map[string]*v.Locale{
    "ee": {
      v.ErrorKeyNotBlank: "{{title}} ei tohi olla tuhi",
    },
  },
})

val := factory.Is(v.String(" ", "nombre").Not().Blank())
```
