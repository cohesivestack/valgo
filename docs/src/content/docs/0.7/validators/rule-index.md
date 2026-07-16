---
title: Rule Index
description: A quick map of rules by validator family in Valgo v0.7.
slug: 0.7/validators/rule-index
---

All validators provide `Not()` and `Or()`. `OrElse()` is not available in
v0.7. Most rule methods accept an optional custom message template.

## String and StringP

- Ordering/equality: `EqualTo`, `GreaterThan`, `GreaterOrEqualTo`, `LessThan`,
  `LessOrEqualTo`, `Between`
- Presence: `Empty`, `Blank`; pointer-only `EmptyOrNil`, `BlankOrNil`, `Nil`
- Byte length: `MaxBytes`, `MinBytes`, `OfByteLength`, `OfByteLengthBetween`
- Rune length: `MaxLength`, `MinLength`, `OfLength`, `OfLengthBetween`
- Other: `InSlice`, `MatchingTo`, `Passing`

## Numeric families

- Common: `EqualTo`, ordering rules, inclusive `Between`, `Zero`, `InSlice`,
  `Passing`
- Signed integers: `Positive`, `Negative`
- Floats: `Positive`, `Negative`, `NaN`, `Infinite`, `Finite`
- Pointer variants: `Nil`, `ZeroOrNil`

## Other families

- Boolean: `EqualTo`, `True`, `False`, `InSlice`, `Passing`; pointer-only
  `Nil`, `FalseOrNil`
- Time: `EqualTo`, `After`, `AfterOrEqualTo`, `Before`, `BeforeOrEqualTo`,
  `Between`, `Zero`, `InSlice`, `Passing`; pointer-only `Nil`, `NilOrZero`
- Comparable: `EqualTo`, `InSlice`, `Passing`; pointer-only `Nil`
- Typed: `Passing`, `Nil`
- Any: `EqualTo`, `Passing`, `Nil`
