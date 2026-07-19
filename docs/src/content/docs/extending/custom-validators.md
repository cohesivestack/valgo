---
title: Custom Go Validators
description: Create custom Valgo validator types in Go with ValidatorContext and reusable validation rules.
---

Valgo is designed to be extended. A custom validator typically wraps a `*valgo.ValidatorContext` and exposes rule methods that add checks.

The pattern used in `custom/custom_validator_test.go` looks like this:

```go
var secretWordLocale = &valgo.Locale{
  "not_valid_secret": "{{title}} is invalid.",
}

type ValidatorSecretWord struct {
  context *valgo.ValidatorContext
}

func (v *ValidatorSecretWord) Correct(template ...string) *ValidatorSecretWord {
  v.context.Add(
    func() bool {
      s := v.context.Value().(string)
      return s == "cohesive" || s == "stack"
    },
    "not_valid_secret",
    template...,
  )
  return v
}

func (v *ValidatorSecretWord) Context() *valgo.ValidatorContext {
  return v.context
}

func SecretWord(value string, nameAndTitle ...string) *ValidatorSecretWord {
  context := valgo.NewContext(value, nameAndTitle...)
  context.WithLocaleFallback(secretWordLocale)

  return &ValidatorSecretWord{context: context}
}
```

`WithLocaleFallback(...)` lets a custom validator provide default messages for
its own error keys. Fallback entries are used only when the active validation
locale does not define the key, so consumers can still override
`not_valid_secret` with `Options{Locale: ...}`.
