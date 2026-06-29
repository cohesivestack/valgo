---
title: Localization
description: Switch locales and override message keys.
slug: 0.7/cookbook/localization
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
