---
title: OR Operators (Or / OrElse)
description: Combine alternative rule branches and control error messages.
slug: 0.7/validators/or-operators
---

Valgo supports OR-style logic inside a validator chain.

## Overview

* `Or(...)` evaluates an alternative branch.
* `OrElse(...)` is like `Or`, but lets you override the resulting message when all branches fail.

## Example

Require either an email-like string or a phone-like string:

```go
emailRe := regexp.MustCompile(`^.+@.+\..+$`)
phoneRe := regexp.MustCompile(`^\+?[0-9]{7,}$`)

val := v.Is(
  v.String(contact, "contact").
    MatchingTo(emailRe).
    Or(v.String(contact, "contact").MatchingTo(phoneRe)),
)
```

If you want a single consolidated error message:

```go
val := v.Is(
  v.String(contact, "contact").
    MatchingTo(emailRe).
    OrElse(
      v.String(contact, "contact").MatchingTo(phoneRe),
      "Contact must be a valid email or phone number",
    ),
)
```
