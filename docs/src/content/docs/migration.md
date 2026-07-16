---
title: Migration Notes
description: Breaking changes and behavioral details across Valgo v0 releases.
---

Valgo is pre-v1.0, so breaking changes can happen.

## v0.6.0 - String length counts runes (characters)

String length validators now measure **runes** instead of bytes:

- A multi-byte UTF-8 character counts as 1.
- Use explicit byte-length validators if you need `len(s)` semantics.

Byte-based validators:

- `MaxBytes`
- `MinBytes`
- `OfByteLength`
- `OfByteLengthBetween`

## v0.7.0 - Numeric validators switched to generics

Numeric validators are now generic per family:

- `ValidatorInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64]`
- `ValidatorUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64]`
- `ValidatorFloat[T ~float32 | ~float64]`

Most call sites remain the same (constructors like `v.Int16(...)` still exist). You mainly need to update declared types.

## v0.8.0 - Go version and validation flow

- Valgo v0.8 is tested with Go 1.23 and later. Using one of these versions is
  recommended.
- When all `Or()` alternatives fail, v0.8 produces one localized message that
  joins the alternatives. OR grouping and precedence remain consistent with
  v0.7.
- `OrElse()` adds a short-circuiting alternative that skips the rest of the
  validator chain when its left side succeeds.
- `PathValid()`, `AllValid()`, and `AnyValid()` query recorded validation
  results. `IsValid()` remains available but is deprecated in favor of
  `PathValid()`.
- The `If*Valid` and `When*Valid` methods conditionally merge validations or
  execute callbacks based on those recorded results.
- Custom validators can supply missing message entries through
  `ValidatorContext.WithLocaleFallback()`.
