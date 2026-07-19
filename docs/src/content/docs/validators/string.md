---
title: String Validators for Go
description: Validate Go strings and *string pointers with Valgo, including blank checks, rune length, byte length, equality, and custom rules.
---

Use `v.String(value, name?, title?)` and `v.StringP(&value, ...)`.

Common rules:

```go
v.Is(v.String("Dennis").EqualTo("Dennis"))
v.Is(v.String("").Empty())
v.Is(v.String(" ").Blank())
v.Is(v.String("processing").InSlice([]string{"idle", "processing"}))
```

## Length: bytes vs characters

Byte-based:

```go
v.Is(v.String("123456").MaxBytes(6))
v.Is(v.String("123").MinBytes(3))
v.Is(v.String("1234").ByteLength(4))
```

Rune-based (characters):

```go
v.Is(v.String("虎視眈々").MaxLength(4))
v.Is(v.String("虎視眈々").Length(4))
v.Is(v.String("虎視眈々").LengthBetween(2, 4))
```

`OfByteLength`, `OfByteLengthBetween`, `OfLength`, and `OfLengthBetween` are
deprecated aliases kept for compatibility. Use `ByteLength`,
`ByteLengthBetween`, `Length`, and `LengthBetween` instead. The `Of*` aliases
will be removed in v1.0.

## Regex

```go
regex, _ := regexp.Compile("pre-.+")
v.Is(v.String("pre-approved").MatchingTo(regex))
```

## Pointer-specific rules

```go
var s *string
v.Is(v.StringP(s).Nil())
```
