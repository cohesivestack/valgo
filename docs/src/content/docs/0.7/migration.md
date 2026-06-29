---
title: Migration Notes
description: Breaking changes and behavioral details across Valgo v0 releases.
slug: 0.7/migration
---

Valgo is pre-v1.0, so breaking changes can happen.

## v0.6.0 - String length counts runes (characters)

String length validators now measure **runes** instead of bytes:

* A multi-byte UTF-8 character counts as 1.
* Use explicit byte-length validators if you need `len(s)` semantics.

Byte-based validators:

* `MaxBytes`
* `MinBytes`
* `OfByteLength`
* `OfByteLengthBetween`

## v0.7.0 - Numeric validators switched to generics

Numeric validators are now generic per family:

* `ValidatorInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64]`
* `ValidatorUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64]`
* `ValidatorFloat[T ~float32 | ~float64]`

Most call sites remain the same (constructors like `v.Int16(...)` still exist). You mainly need to update declared types.
