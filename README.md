# Valgo

Valgo is a type-safe, expressive, and extensible validator library for Golang. Valgo is built with generics, so Go 1.18 or higher is required.


Valgo differs from other Golang validation libraries in that the rules are written in code and not as struct tags. This allows greater flexibility and freedom when it comes to where and how data is validated.

Moreover, Valgo supports customizing and localizing validation messages, making it a useful alternative for API interface development.

## Syntax example

Here is a quick example of the syntax:

```go
result := v.
  Is(v.String(" ").NotBlank().LengthBetween(1,2)).
  Is(v.Number(10).GreaterThan(12)).
  Valid()
```

## Getting Started

Install in your project:

```bash
go get github.com/cohesivestack/valgo
```

Import in your code:
```go
import v github.com/cohesivestack/valgo
```

## Using Valgo

### `Is(...)` function and validation session

The `Is(...)` function allows you to pass the value and the rules for validating it. The function returns a validation session which allows for the validation of more values. As shown in the following example, we are validating three different values within the same validation session.

```go
validator := v.
  Is(v.String(" ").Not().Blank().LengthBetween(1,2))
```



### `Valid()` function

An validation session provides a function `Valid()` that returns the validation result.

```go
validator := v.
  Is(v.String(" ").Not().Blank().LengthBetween(1,2))
```

Valid() returns the result, but it does not end the validation session, so you can keep adding values.

```go
v.Is(String(" ").
// Return true
```

### `IsValid()` function

TODO

### `In(...)` function

TODO

### Concat a validator `Merge( ... )`

The `Merge( ... )` allows to merge two different validation sessions.

### Name and title of the value

TODO

### Custom message template

TODO

### Localizing the validator messages

TODO

### The `Passing( ... )` validators

TODO

### Customizing the output JSON

TODO

## Validators

TODO

### String Validators

- [x] `EqualTo`
- [x] `GreaterThan`
- [x] `GreaterOrEqualTo`
- [x] `LessThan`
- [x] `LessOrEqualTo`
- [x] `Between`
- [x] `Empty`
- [x] `Blank`
- [x] `Passing`
- [x] `InSlice`
- [x] `MatchingTo`
- [x] `MaxLength`
- [x] `MinLength`
- [x] `Length`
- [x] `LengthBetween`
- [x] `BlankOrNil` (for pointer validator - `StringP`)
- [x] `EmptyOrNil` (for pointer validator - `StringP`)
- [x] `Nil` (for pointer validator - `StringP`)

### Boolean Validators

- [x] `EqualTo`
- [x] `Passing`
- [x] `True`
- [x] `False`
- [x] `InSlice`
- [x] `FalseOrNil` (for pointer validator - `BoolP`)`
- [x] `Nil` (for pointer validator - `BoolP`)`

### Number Validators

- [x] `EqualTo`
- [x] `GreaterThan`
- [x] `GreaterOrEqualTo`
- [x] `LessThan`
- [x] `LessOrEqualTo`
- [x] `Between`
- [x] `Zero`
- [x] `InSlice`
- [x] `Passing`
- [x] `ZeroOrNil` (for pointer value - `NumberP(...)`)
- [x] `Nil` (for pointer value - `NumberP(...)`)

### Any Validators

- [x] `EqualTo`
- [x] `Passing`
- [x] `Nil`

## Custom Validators

Although all validator types in Golang has a `Passing( ... )`, to validate a value with a custom function. Valgo allow you to create your own validator implementing the X interface.

Suppose we want to create a validation for a specific struct


## License

Copyright © 2022 Carlos Forero

Valgo is released under the [MIT License](LICENSE)