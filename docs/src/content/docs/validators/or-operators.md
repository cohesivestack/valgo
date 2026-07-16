---
title: OR Operators (Or / OrElse)
description: Combine rules in a validator chain with OR logic and short-circuiting.
---

Valgo provides `Or()` and `OrElse()` for expressing alternatives inside a
validator chain. Both methods take no arguments and apply to the rules that
surround them in the same chain.

## Or

`Or()` groups adjacent rules into a logical OR expression. The group passes
when any rule in it passes and fails only when every rule fails.

For example, accept a contact value that looks like either an email address or
a phone number:

```go
emailRe := regexp.MustCompile(`^.+@.+\..+$`)
phoneRe := regexp.MustCompile(`^\+?[0-9]{7,}$`)

val := v.Is(
  v.String(contact, "contact").
    MatchingTo(emailRe).
    Or().
    MatchingTo(phoneRe),
)
```

You can join more than two alternatives:

```go
val := v.Is(
  v.Int(status, "status").
    EqualTo(1).
    Or().EqualTo(2).
    Or().EqualTo(3),
)
```

### Precedence

An OR group is evaluated before the implicit AND that continues the chain:

```go
v.String(value).
  MinLength(5).
  Or().EqualTo("test").
  Not().Blank()
```

This is evaluated as:

```text
(MinLength(5) OR EqualTo("test")) AND Not().Blank()
```

Rules that fail before an OR group begins are not rescued by a later `Or()`.
For example, `A.B.Or().C` is evaluated as `A AND (B OR C)`.

### Error messages

If every rule in an OR group fails, Valgo returns one message containing all
the alternatives. The separator is localized for the validation session.

For example:

```go
val := v.Is(
  v.Int(1, "status", "Status").
    EqualTo(2).
    Or().EqualTo(3),
)
```

Produces:

```text
Status must be equal to "2" or Status must be equal to "3"
```

## OrElse

`OrElse()` introduces an OR boundary with short-circuiting. If the rule or OR
group to its left passes, the entire remainder of the validator chain is
skipped. If the left side fails, validation continues on the right normally.

This is useful for optional values: accept an empty value immediately;
otherwise, require every rule that follows it to pass.

```go
val := v.Is(
  v.String(value, "value").
    Empty().
    OrElse().
    MinLength(5).
    EqualTo("test"),
)
```

This behaves as:

```text
Empty() OR (MinLength(5) AND EqualTo("test"))
```

When `Empty()` passes, neither `MinLength(5)` nor `EqualTo("test")` is
evaluated. When it fails, both rules on the right are evaluated using the
normal `Is()` or `Check()` behavior.

`OrElse()` participates in an OR group like `Or()` and uses the same localized
joined error format when all alternatives fail. It does not accept a custom
message. To customize an error, pass a template to the relevant validation
rule.
