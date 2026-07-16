---
title: Localization & Factory
description: Localize messages, override templates, and reuse options with a
  ValidationFactory.
slug: 0.7/using-valgo/localization
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
