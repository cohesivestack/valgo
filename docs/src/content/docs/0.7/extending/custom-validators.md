---
title: Custom Validators
description: Create your own validator types using ValidatorContext.
slug: 0.7/extending/custom-validators
---

A custom validator wraps a `*valgo.ValidatorContext`, implements `Context()`,
and exposes rule methods that add checks.

```go
type ValidatorSecretWord struct {
  context *valgo.ValidatorContext
}

func (validator *ValidatorSecretWord) Correct(template ...string) *ValidatorSecretWord {
  validator.context.Add(
    func() bool {
      value := validator.context.Value().(string)
      return value == "cohesive" || value == "stack"
    },
    "not_valid_secret",
    template...,
  )
  return validator
}

func (validator *ValidatorSecretWord) Context() *valgo.ValidatorContext {
  return validator.context
}

func SecretWord(value string, nameAndTitle ...string) *ValidatorSecretWord {
  return &ValidatorSecretWord{
    context: valgo.NewContext(value, nameAndTitle...),
  }
}
```

The custom error key must exist in the validation locale or the rule must
receive a template:

```go
val := valgo.New(valgo.Options{
  Locale: &valgo.Locale{
    "not_valid_secret": "{{title}} is invalid.",
  },
}).Is(SecretWord("loose", "secret").Correct())
```

`ValidatorContext.WithLocaleFallback()` was added after v0.7 and is not
available in this version.
