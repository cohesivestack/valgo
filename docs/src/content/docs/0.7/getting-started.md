---
title: Getting Started
description: Install Valgo and validate values with Is(), Check(), and New().
slug: 0.7/getting-started
---

## Introduction

Valgo is a type-safe, expressive, and extensible validation library for Go with built-in i18n support.

Unlike validation libraries that rely on struct tags, Valgo defines validation rules as functions. This gives you greater flexibility to validate any value, compose rules programmatically, and decide where validation belongs within your application.

Valgo can be customized to fit your application's needs, from overriding validation messages to localizing them for different languages and contexts.

## Install

```bash
go get github.com/cohesivestack/valgo
```

Go 1.19 or later is required by Valgo v0.7.

## Your first validation

`Is(...)` creates a validation session. Within each validator chain, evaluation
normally stops after the first failed rule. Rules connected by `Or()` follow
the v0.7 OR behavior described under `Validators -> OR Operator`.

```go
import (
  "encoding/json"
  "fmt"
  v "github.com/cohesivestack/valgo"
)

val := v.Is(
  v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20),
  v.Number(17, "age").GreaterThan(18),
)

if err := val.ToError(); err != nil {
  out, _ := json.MarshalIndent(err, "", "  ")
  fmt.Println(string(out))
}
```

## When to use Is vs Check

* `Is(...)`: normally stops a validator chain after its first failed rule.
* `Check(...)`: continues after failures to collect multiple messages. An
  alternative after `Or()` can still be skipped when its left side succeeds.

```go
val := v.Check(
  v.String("", "full_name").Not().Blank().OfLengthBetween(4, 20),
)

_ = val.Valid() // false, with 2 messages for full_name
```

## Nested models

Use namespaces to build structured paths:

* `In("ns", ...)` for nested structs
* `InRow("list", i, ...)` for slices of structs
* `InCell("list", i, ...)` for slices of scalar values

See `Using Valgo -> Namespaces`.
