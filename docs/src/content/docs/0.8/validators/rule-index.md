---
slug: 0.8/validators/rule-index
title: Rule Index
description: A quick map of rules and constructors by validator family.
---

All validators provide `Not()`, `Or()`, and `OrElse()`. Most rule methods also
accept an optional custom message template as their last argument.

## String and StringP

- Equality and ordering: `EqualTo`, `GreaterThan`, `GreaterOrEqualTo`,
  `LessThan`, `LessOrEqualTo`, `Between`
- Presence: `Empty`, `Blank`; pointer forms also have `EmptyOrNil`,
  `BlankOrNil`, and `Nil`
- Length in bytes: `MaxBytes`, `MinBytes`, `OfByteLength`,
  `OfByteLengthBetween`
- Length in runes: `MaxLength`, `MinLength`, `OfLength`, `OfLengthBetween`
- Other: `InSlice`, `MatchingTo`, `Passing`

## Numeric families

- Common: `EqualTo`, `GreaterThan`, `GreaterOrEqualTo`, `LessThan`,
  `LessOrEqualTo`, `Between`, `Zero`, `InSlice`, `Passing`
- Signed integers: `Positive`, `Negative`
- Floats: `Positive`, `Negative`, `NaN`, `Infinite`, `Finite`
- Pointer variants: `Nil`, `ZeroOrNil`

## Boolean

`EqualTo`, `True`, `False`, `InSlice`, `Passing`; pointer forms also provide
`Nil` and `FalseOrNil`.

## Time

`EqualTo`, `After`, `AfterOrEqualTo`, `Before`, `BeforeOrEqualTo`, `Between`,
`Zero`, `InSlice`, `Passing`; the pointer form also provides `Nil` and
`NilOrZero`.

## Comparable

`EqualTo`, `InSlice`, and `Passing`; the pointer form also provides `Nil`.
Comparable validators do not provide ordering methods.

## Typed and Any

- `Typed`: `Passing`, `Nil`
- `Any`: `EqualTo`, `Passing`, `Nil`
