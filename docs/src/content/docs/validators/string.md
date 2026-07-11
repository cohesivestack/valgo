---
title: String
description: Validate strings and *string pointers, including rune-length and byte-length rules.
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
v.Is(v.String("1234").OfByteLength(4))
```

Rune-based (characters):

```go
v.Is(v.String("虎視眈々").MaxLength(4))
v.Is(v.String("虎視眈々").OfLengthBetween(2, 4))
```

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
