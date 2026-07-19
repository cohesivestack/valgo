---
title: Localized Validation Errors in Go
description: Switch Valgo locales and override validation message keys to return localized Go validation errors.
---

## Switch locale

```go
val := v.New(v.Options{LocaleCode: "es"}).
  Is(v.String(" ", "nombre").Not().Blank())
```

## Add/override a key

```go
val := v.New(v.Options{
  Locale: &v.Locale{
    v.ErrorKeyNotBlank: "{{title}} should not be blank",
  },
}).Is(v.String(" ", "name").Not().Blank())
```
