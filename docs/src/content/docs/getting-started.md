---
title: Getting Started
description: Install Valgo and validate values with Is(), Check(), and New().
---

## Introduction

Valgo is a type-safe, expressive, and extensible validator for Go with built-in i18n support.

Unlike validation libraries that rely on struct tags, Valgo defines validation rules as functions. This gives you greater flexibility to validate any value, compose rules programmatically, and decide where validation belongs within your application.

Valgo can be customized to fit your application's needs, from overriding validation messages to localizing them for different languages and contexts.

## Install

```bash
go get github.com/cohesivestack/valgo
```

## Agent skill

This repository includes a Valgo Agent Skill installable with [`npx skills`](https://github.com/vercel-labs/skills):

```bash
npx skills add cohesivestack/valgo --skill valgo
```

## Your first validation

`Is(...)` creates a validation session and short-circuits per field: once a rule fails for a value, later rules for that value are skipped.

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

- `Is(...)`: short-circuits rules per value (usually what you want for UX and performance).
- `Check(...)`: evaluates every rule even if earlier rules fail (useful when you want all messages at once).

```go
val := v.Check(
  v.String("", "full_name").Not().Blank().OfLengthBetween(4, 20),
)

_ = val.Valid() // false, with 2 messages for full_name
```

## Nested models

Use namespaces to build structured paths:

- `In("ns", ...)` for nested structs
- `InRow("list", i, ...)` for slices of structs
- `InCell("list", i, ...)` for slices of scalar values

See `Using Valgo -> Namespaces`.
