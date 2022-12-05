# Valgo

Valgo is a type-safe, expressive, and extensible validator library for Golang. Valgo is built with generics, so Go 1.18 or higher is required.

Valgo differs from other Golang validation libraries in that the rules are written in functions and not in struct tags. This allows greater flexibility and freedom when it comes to where and how data is validated.

Additionally, Valgo supports customizing and localizing validation messages.

# Quick example

Here is a quick example:
```go
package main

import v "github.com/cohesivestack/valgo"

func main() {
  val := v.
    Is(v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
    Is(v.Number(17, "age").GreaterThan(18))

  if !val.Valid() {
    out, _ := json.MarshalIndent(val.Error(), "", "  ")
    fmt.Println(string(out))
  }
}
```
output:
```json
{
  "age": [
    "Age must be greater than \"18\""
  ],
  "full_name": [
    "Full name must have a length between \"4\" and \"20\""
  ]
}
```

# v0.x.x and backward compatibility

Valgo is in its early stages, so backward compatibility won't be guaranteed until v1.

Valgo is used in production by [Statsignal](https://statsignal.dev), but we want community feedback before setting up version 1.

# Table of content

- [Valgo](#valgo)
- [Quick example](#quick-example)
- [v0.x.x and backward compatibility](#v0xx-and-backward-compatibility)
- [Table of content](#table-of-content)
- [Getting started](#getting-started)
- [Using Valgo](#using-valgo)
  - [`Validation` session](#validation-session)
  - [`Is(...)` function](#is-function)
  - [`New()` function](#new-function)
  - [`Validation.Valid()` function](#validationvalid-function)
  - [`Validation.IsValid(...)` function](#validationisvalid-function)
  - [`In(...)` function](#in-function)
  - [`InRow(...)` function](#inrow-function)
  - [`Check(...)` function](#check-function)
  - [Merging two `Validation` sessions with `Validation.Merge( ... )`](#merging-two-validation-sessions-with-validationmerge--)
- [Validators](#validators)
  - [Validator value's name and title.](#validator-values-name-and-title)
  - [`Not()` validator function](#not-validator-function)
  - [String validator](#string-validator)
  - [String pointer validator](#string-pointer-validator)
  - [Number validator](#number-validator)
  - [Number pointer validator](#number-pointer-validator)
  - [Number specific type validators](#number-specific-type-validators)
  - [Bool validator](#bool-validator)
  - [Boolean pointer validator](#boolean-pointer-validator)
  - [Any validator](#any-validator)
  - [Custom type validators](#custom-type-validators)
- [Customizing](#customizing)
  - [Custom errors JSON output](#custom-errors-json-output)
  - [Custom error message template](#custom-error-message-template)
  - [Extending Valgo with custom validators](#extending-valgo-with-custom-validators)
- [Localizing validator messages](#localizing-validator-messages)
- [List of rules by validator type](#list-of-rules-by-validator-type)
- [License](#license)

# Getting started

Install in your project:

```bash
go get github.com/cohesivestack/valgo
```

Import in your code:
```go
import v github.com/cohesivestack/valgo
```
**Note**: You can use any other aliases instead of `v` or just reference the
package `valgo` directly.

# Using Valgo

## `Validation` session

The `Validation` session in Valgo is the main structure for validating one or more values. It is called 'Validation' in code.

A validation session will contain one or more Validators, where each `Validator` will have the responsibility to validate a value with one or more rules.

There are multiple functions to create a `Validation` session, depending on the requirements:

  * `New()`,
  * `Is(...)`,
  * `In(...)`,
  * `InRow(...)`,
  * `Check(...)`,
  * `AddErrorMessage(...)`

`Is(...)` is likely to be the most frequently used function in your validations. When `Is(...)` is called, the function creates a validation and receives a validator at the same time. In the next section, you will learn more about the `Is(...)` function.

## `Is(...)` function

The `Is(...)` function allows you to pass a `Validator` with the value and the rules for validating it. At the same time, create a `Validation` session, which lets you add more Validators in order to verify more values.

As shown in the following example, we are passing to the function `Is(...)` the `Validator` for the `full_name` value. The function returns a `Validation` session that allows us to add more Validators to validate more values; in the example case the values `age` and `status`:

```go
val := v.
  Is(v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
  Is(v.Number(17, "age").GreaterThan(18)).
  Is(v.String("singl", "status").InSlice([]string{"married", "single"}))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}
```
output:
```json
{
  "age": [
    "Age must be greater than \"18\""
  ],
  "full_name": [
    "Full name must have a length between \"4\" and \"20\""
  ],
  "status": [
    "Status is not valid"
  ]
}
```

## `New()` function

This function allows you to create a new `Validation` session without a `Validator`. This is useful for conditional validation or reusing validation logic.

The following example conditionally adds a `Validator` rule for the `month_day` value.

```go
month := 5
monthDay := 11

val := v.New()

if month == 6 {
  val.Is(v.Number(monthDay, "month_day").LessOrEqualTo(10))
}

if val.Valid() {
  fmt.Println("The validation passes")
}
```
output:
```bash
The validation passes
```


## `Validation.Valid()` function

A `Validation` session provide this function, which returns either `true` if all their validators are valid or `false` if any one of them is invalid.

In the following example, even though the Validator for `age` is valid, the `Validator` for `status` is invalid, making the entire `Validator` session invalid.

```go
val := v.Is(v.Number(21, "age").GreaterThan(18)).
  Is(v.String("singl", "status").InSlice([]string{"married", "single"}))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}
```
output:
```json
{
  "status": [
    "Status is not valid"
  ]
}
```

## `Validation.IsValid(...)` function

This functions allows to check if an specific value in a `Validation` session is valid or not. This is very useful for conditional logic.

The following example prints an error message if the `age` value is invalid.

```go
val := v.Is(v.Number(16, "age").GreaterThan(18)).
  Is(v.String("single", "status").InSlice([]string{"married", "single"}))

if !val.IsValid("age") {
  fmt.Println("Warning: someone underage is trying to sign up")
}
```
output:
```
Warning: someone underage is trying to sign up
```

## `In(...)` function

The `In(...)` function executes one or more validators in a namespace, so the value names in the error result are prefixed with this namespace. This is useful for validating nested structures.

In the following example we are validating the `Person` and the nested `Address` structure. We can distinguish the errors of the nested `Address` structure in the error results.

```go
type Address struct {
  Name string
  Street string
}

type Person struct {
  Name string
  Address Address
}

p := Person{"Bob", Address{"", "1600 Amphitheatre Pkwy"}}

val := v.
  Is(v.String(p.Name, "name").OfLengthBetween(4, 20)).
  In("address",
    Is(String(p.Address.Name, "name").Not().Blank()).
    Is(String(p.Address.Street, "street").Not().Blank()))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}

```
output:
```json
{
  "address.name": [
    "Name can't be blank"
  ],
  "name": [
    "Name must have a length between \"4\" and \"20\""
  ]
}
```

## `InRow(...)` function

The `InRow(...)` function executes one or more validators in a namespace similar to the `In(...)` function, but with indexed namespace. So, the value names in the error result are prefixed with this indexed namespace. It is useful for validating nested lists in structures.

In the following example we validate the `Person` and the nested list `Addresses`. The error results can distinguish the errors of the nested list `Addresses`.

```go
type Address struct {
  Name   string
  Street string
}

type Person struct {
  Name      string
  Addresses []Address
}

p := Person{
  "Bob",
  []Address{
    {"", "1600 Amphitheatre Pkwy"},
    {"Home", ""},
  },
}

val := v.Is(String(p.Name, "name").OfLengthBetween(4, 20))

for i, a := range p.Addresses {
  val.InRow("addresses", i,
    v.Is(v.String(a.Name, "name").Not().Blank()).
    v.Is(v.String(a.Street, "street").Not().Blank()))
}

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}
```
output:
```json
{
  "addresses[0].name": [
    "Name can't be blank"
  ],
  "addresses[1].street": [
    "Street can't be blank"
  ],
  "name": [
    "Name must have a length between \"4\" and \"20\""
  ]
}
```

## `Check(...)` function

The `Check(...)` function, similar to the `Is(...)` function, however with `Check(...)` the Rules of the Validator parameter are not short-circuited, which means that regardless of whether a previous rule was valid, all rules are checked.

This example shows two rules that fail due to the empty value in the `full_name` `Validator`, and since the `Validator` is not short-circuited, both error messages are added to the error result.

```go
val := v.Check(v.String("", "full_name").Not().Blank().OfLengthBetween(4, 20))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}
```
output:
```json
{
  "full_name": [
    "Full name can't be blank",
	  "Full name must have a length between \"4\" and \"20\""
  ]
}
```

## Merging two `Validation` sessions with `Validation.Merge( ... )`

Using `Merge(...)` you can merge two `Validation` sessions. When two validations are merged, errors with the same value name will be merged. It is useful for reusing validation logic.

The following example merges the `Validation` session returned by the `validatePreStatus` function. Since both `Validation` sessions validate a value with the name `status`, the error returned will return two error messages, and without duplicate the `Not().Blank()` error message rule.

```go

type Record struct {
  Name   string
  Status string
}

validatePreStatus := func(status string) *Validation {
  regex, _ := regexp.Compile("pre-.+")

  return v.
    Is(v.String(status, "status").Not().Blank().MatchingTo(regex))
}

r := Record{"Classified", ""}

val := v.
  Is(v.String(r.Name, "name").Not().Blank()).
  Is(v.String(r.Status, "status").Not().Blank())

val.Merge(validatePreStatus(r.Status))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.Error(), "", "  ")
  fmt.Println(string(out))
}

```
output:
```json
{
  "status": [
	  "Status can't be blank",
	  "Status must match to \"pre-.+\""
  ]
}
```

# Validators

Validators establish the rules for validating a value. The validators are passed to a `Validation` session.

For each primitive Golang value type, Valgo provides a `Validator`. A `Validator` has different functions that set its value's validation rules.

Although Valgo has multiple types of validators, it can be extended with custom validators. Check the section [Extending Valgo with Custom Validators](#extending-valgo-with-custom-validators) for more information.

## Validator value's name and title.

Validators only require the value to be validated, so, for example, the following code validates a string value by checking if it is empty.

```go
val := v.New(v.String("").Empty())
```
`val.Error()` output:
```json
{
  "value_0": [
	  "Value 0 can't be empty",
  ]
}
```
In the example above, since we didn't specify a name for the value, Valgo generates a `value_0` name and consequently the `Value 0` title in the error message.

However, Validators allow you, optionally, to specify the value's name and title, as shown below:

Validator with value's name:
```go
val := v.New(v.String("", "company_name").Not().Empty())
```
`val.Error()` output:
```json
{
  "company_name": [
	  "Company name can't be empty",
  ]
}
```

Validator with value's name and title:
```go
val := v.New(v.String("", "company_name", "Customer").Not().Empty())
```
`val.Error()` output:
```json
{
  "company_name": [
	  "Customer can't be empty",
  ]
}
```

## `Not()` validator function

Valgo validators have a `Not()` function to invert the boolean value associated with the next validator rule function.

In the following example, the call to `Valid()` will return false because `Not()` inverts the boolean value associated with the `Zero()` function.

```go
valid := Is(v.Number(0).Not().Zero()).Valid()

fmt.Println(valid)
```
output:
```
false
```

## String validator

The `ValidatorString` provides functions for setting validation rules for a `string` type value, or a custom type based on a `string`.

Below is a valid example for every String validator rule.

```go
v.Is(v.String("Dennis Ritchie").EqualTo("Dennis Ritchie"))
v.Is(v.String("Steve Jobs").GreaterThan("Bill Gates"))
v.Is(v.String("Steve Jobs").GreaterOrEqualTo("Elon Musk"))
v.Is(v.String("C#").LessThan("Go"))
v.Is(v.String("Go").LessOrEqualTo("Golang"))
v.Is(v.String("Rust").Between("Go", "Typescript")) // Inclusive
v.Is(v.String("").Empty())
v.Is(v.String(" ").Blank())
v.Is(v.String("Dart").Passing(func(val string) bool { return val == "Dart" }))
v.Is(v.String("processing").InSlice([]string{"idle", "processing", "ready"})
v.Is(v.String("123456").MaxLength(6))
v.Is(v.String("123").MinLength(3))
v.Is(v.String("1234").MinLength(4))
v.Is(v.String("12345").LengthBetween(4,6)) // Inclusive
regex, _ := regexp.Compile("pre-.+"); v.Is(String("pre-approved").MatchingTo(regex))
```

## String pointer validator

The `ValidatorStringP` provides functions for setting validation rules for a `string` type pointer, or a custom type based on a `string` pointer.

Below is a valid example for every String pointer validator rule.

```go
x := "Dennis Ritchie"; v.Is(v.StringP(&x).EqualTo("Dennis Ritchie"))
x := "Steve Jobs";     v.Is(v.StringP(&x).GreaterThan("Bill Gates"))
x := "Steve Jobs";     v.Is(v.StringP(&x).GreaterOrEqualTo("Elon Musk"))
x := "C#";             v.Is(v.StringP(&x).LessThan("Go"))
x := "Go";             v.Is(v.StringP(&x).LessOrEqualTo("Golang"))
x := "Rust";           v.Is(v.StringP(&x).Between("Go", "Typescript")) // Inclusive
x := "";               v.Is(v.StringP(&x).Empty())
x := " ";              v.Is(v.StringP(&x).Blank())
x := "Dart";           v.Is(v.StringP(&x).Passing(func(val *string) bool { return *val == "Dart" }))
x := "processing";     v.Is(v.StringP(&x).InSlice([]string{"idle", "processing", "ready"})
x := "123456";         v.Is(v.StringP(&x).MaxLength(6))
x := "123";            v.Is(v.StringP(&x).MinLength(3))
x := "1234";           v.Is(v.StringP(&x).MinLength(4))
x := "12345";          v.Is(v.StringP(&x).LengthBetween(4,6)) // Inclusive
x := "pre-approved"; regex, _ := regexp.Compile("pre-.+"); v.Is(StringP(&x).MatchingTo(regex))
x := "";               v.Is(v.StringP(&x).EmptyOrNil())
x := " ";              v.Is(v.StringP(&x).BlankOrNil())
var x *string;         v.Is(v.StringP(x).Nil())
```

## Number validator

The Number validator provides functions for setting validation rules for a `TypeNumber` value, or a custom type based on a `TypeNumber`.

`TypeNumber` is a generic interface defined by Valgo that generalizes any standard Golang type. Below is Valgo's definition of `TypeNumber`:


```go
type TypeNumber interface {
  ~int |
  ~int8 |
  ~int16 |
  ~int32 |
  ~int64 |
  ~uint |
  ~uint8 |
  ~uint16 |
  ~uint32 |
  ~uint64 |
  ~float32 |
  ~float64
}
```

Below is a valid example for every Number validator rule.

```go
v.Is(v.Number(10).EqualTo(10))
v.Is(v.Number(11).GreaterThan(10))
v.Is(v.Number(10).GreaterOrEqualTo(10))
v.Is(v.Number(10).LessThan(11))
v.Is(v.Number(10).LessOrEqualTo(10))
v.Is(v.Number(11).Between(10, 12)) // Inclusive
v.Is(v.Number(0).Zero())
v.Is(v.Number(10).Passing(func(val int) bool { return val == 10 }))
v.Is(v.Number(20).InSlice([]int{10, 20, 30}))
```

## Number pointer validator

The Number pointer validator provides functions for setting validation rules for a `TypeNumber` pointer, or a custom type based on a `TypeNumber` pointer.

As it's explained in [Number validator](#number-validator), the `TypeNumber` is a generic interface defined by Valgo that generalizes any standard Golang type.

Below is a valid example for every Number pointer validator rule.

```go
x := 10;    v.Is(v.NumberP(&x).EqualTo(10))
x := 11;    v.Is(v.NumberP(&x).GreaterThan(10))
x := 10;    v.Is(v.NumberP(&x).GreaterOrEqualTo(10))
x := 10;    v.Is(v.NumberP(&x).LessThan(11))
x := 10;    v.Is(v.NumberP(&x).LessOrEqualTo(10))
x := 11;    v.Is(v.NumberP(&x).Between(10, 12)) // Inclusive
x := 0;     v.Is(v.NumberP(&x).Zero())
x := 10;    v.Is(v.NumberP(&x).Passing(func(val *int) bool { return *val == 10 }))
x := 20;    v.Is(v.NumberP(&x).InSlice([]int{10, 20, 30}))
x := 0;     v.Is(v.NumberP(&x).ZeroOrNil())
var x *int; v.Is(v.NumberP(x).Nil())
```

## Number specific type validators

While the validator `Number` works with all golang number types, Valgo also has a validator for each type. You can use them if you prefer or need a stronger safe type code.

Following is a list of functions for every specific number type validator, along with their equivalent pointer validators.

```go
Int(v int)         IntP(v *int)
Int8(v int8)       Int8P(v *int8)
Int16(v int16)     Int16P(v *int16)
Int32(v int32)     Int32P(v *int32)
Int64(v int64)     Int64P(v *int64)
Uint(v uint)       UintP(v *uint)
Uint8(v uint8)     Uint8P(v *uint8)
Uint16(v uint16)   Uint16P(v *uint16)
Uint32(v uint32)   Uint32P(v *uint32)
Uint64(v uint64)   Uint64P(v *uint64)
Float32(v float32) Float32P(v *float32)
Float64(v float64) Float64P(v *float64)
```

These validators have the same rule functions as the `Number` validator.

Similar to the `Number` validator, custom types can be passed based on the specific number type.

## Bool validator

The Bool validator provides functions for setting validation rules for a `bool` type value, or a custom type based on a `bool`.

Below is a valid example for every Bool validator rule.

```go
v.Is(v.Bool(true).EqualTo(true))
v.Is(v.Bool(true).True())
v.Is(v.Bool(false).False())
v.Is(v.Bool(true).Passing(func(val bool) bool { return val == true }))
v.Is(v.Bool(true).InSlice([]string{true, false}))
```

## Boolean pointer validator

The Bool pointer validator provides functions for setting validation rules for a `bool` pointer, or a custom type based on a `bool` pointer.

Below is a valid example for every Bool pointer validator rule.

```go
x := true;   v.Is(v.BoolP(&x).EqualTo(true))
x := true;   v.Is(v.BoolP(&x).True())
x := false;  v.Is(v.BoolP(&x).False())
x := true;   v.Is(v.BoolP(&x).Passing(func(val *bool) bool { return val == true }))
x := true;   v.Is(v.BoolP(&x).InSlice([]string{true, false}))
x := false;  v.Is(v.BoolP(&x).FalseOrNil())
var x *bool; v.Is(v.BoolP(x).Nil())
```

## Any validator

With the Any validator, you can set validation rules for any value or pointer.

Below is a valid example of every Any validator rule.

```go
v.Is(v.Any("react").EqualTo("react"))
v.Is(v.Any("svelte").Passing(func(val *bool) bool { return val == "svelte" }))
var x *bool; v.Is(v.Any(x).Nil())
```

For the `EqualTo(v any)` rule function, the parameter type must match the type used by the `Any()` function, otherwise it will be invalid. In the following example, since the value passed to the `Any(...)` function is `int`, and `EqualTo(...)` compares it with int64, the validation is invalid.

```go
valid := v.Is(v.Any(10).EqualTo(int64(10))).Valid()
fmt.Println(valid)
```
output
```
false
```

If a pointer is used, the same pointer must be passed to `EqualTo(v any)` as it is passed to `Any(v any)`, in order to get a valid validation. The following example illustrates it.

```go
// Valid since the same pointers are compared
numberA := 10
v.Is(v.Any(&numberA).EqualTo(&numberA)).Valid()

// Invalid since different pointers are compared
numberB := 10
v.Is(v.Any(&numberA).EqualTo(&numberB)).Valid()
```

## Custom type validators

All golang validators allow to pass a custom type based on its value type. Bellow some valid examples.

```go
type Status string
var status Status = "up"
val := v.Is(v.String(status).InSlice([]Status{"down", "up", "paused"}))

type Level int
var level Level = 1
val = v.Is(v.Int(level).LessThan(Level(2)))

type Stage int64
var stage Stage = 2
val := v.Is(v.NumberP(&stage).GreaterThan(Stage(1)))
```

# Customizing

## Custom errors JSON output

It is possible to customize the JSON output for errors using the function `SetMarshalJSON(err Error)`. The parameter in this function receives a `valgo.Error` structure, which provides all information regarding output validation errors. Below is an example of a customized function.

```go
customMarshalJson := func(e *Error) ([]byte, error) {

  errors := map[string]interface{}{}

  for k, v := range e.errors {
    errors[k] = v.Messages()
  }

  // Add a root key level called errors, which is not set by default in the Valgo implementation.
  return json.Marshal(map[string]map[string]interface{}{"errors": errors})
}

// Set the custom Marshal JSON function
v.SetMarshalJSON(customMarshalJson)

// Now validate something to check if the output JSON contains the errors root key

val := v.Is(v.String("", "name").Not().Empty())

out, _ := json.MarshalIndent(val.Error(), "", "  ")
fmt.Println(string(out))
```
output:
```json
{
  "errors": {
    "name": [
      "Name can't be empty"
    ]
  }
}
```

## Custom error message template

Customizing the default Valgo error messages is possible through the Valgo localization functions as it's explained in the [Localizing validator messages](#localizing-validator-messages) section, however the Valgo validators allow to customize the template of a specific template validator rule. Below is an example illustrating this with the String empty validator rule.

```go
val := v.Is(v.String("", "address_field", "Address").Not().Empty("{{title}} must not be empty. Please provide the value in the input {{name}}."))

out, _ := json.MarshalIndent(val.Error(), "", "  ")
fmt.Println(string(out))
```
output:
```json
{
  "address": [
	  "Address must not be empty. Please provide the value in the input address_field."
  ]
}
```

## Extending Valgo with custom validators

While all validators in Golang provide a `Passing(...)` function, which allows you to use a custom validator function, Valgo also allows you to create your own validator.

With this functionality Valgo can be extended with Validator libraries, which we encourage the community to do.

For example, let's say we want to create a validation for the following `ID` struct, where a user must provide at least one property.

The struct to validate:
```go
// Type to validate
type ID struct {
  Phone string
  Email string
}
```
the custom validator code:
```go
// The custom validator type
type ValidatorID struct {
	context *valgo.ValidatorContext
}

// The custom validator implementation of `valgo.Validator`
func (validator *ValidatorID) Context() *valgo.ValidatorContext {
	return validator.context
}

// Here is the function that passes the value to the custom validator
func IDValue(value ID, nameAndTitle ...string) *ValidatorID {
	return &ValidatorID{context: valgo.NewContext(value, nameAndTitle...)}
}

// The Empty rule implementation
func (validator *ValidatorID) Empty(template ...string) *ValidatorID {
	validator.context.Add(
		func() bool {
			return len(strings.Trim(validator.context.Value().(ID).Phone, " ")) == 0 &&
				len(strings.Trim(validator.context.Value().(ID).Email, " ")) == 0
		},
		v.ErrorKeyEmpty, template...)

	return validator
}

// It would be possible to create a rule NotEmpty() instead of Empty(), but if you add a Not() function then your validator will be more flexible.

func (validator *ValidatorID) Not() *ValidatorID {
	validator.context.Not()

	return validator
}

```
using our validator:
```go
val := v.Is(IDValue(ID{}, "id").Not().Empty())

out, _ := json.MarshalIndent(val.Error(), "", "  ")
fmt.Println(string(out))
```
output:
```json
{
  "identification": [
	  "Id can't be empty"
  ]
}
```

# Localizing validator messages

Valgo has localized error messages. The error messages are only available in English (default) and Spanish at the moment. It is possible, however, to set the error messages for any locale using the function `SetDefaultLocale(code string)`. Using this function, you can also customize the existing Valgo locale messages.

In the following example we are adding localized messages for the Estonian language. Essentially, we are copying the Valgo English messages and customizing the Not Blank error message template.

```go
messages, _ := v.GetLocaleMessages("en")

messages[ErrorKeyNotBlank] = "{{title}} ei tohi olla tühi"
v.SetLocaleMessages("ee", messages)

localized, err := v.Localized("ee")

// Testing the output
val := localized.New().Check(v.String(" ", "name").Not().Blank())

out, _ := json.MarshalIndent(val.Error(), "", "  ")
fmt.Println(string(out))

```
output:
```json
{
  "name": [
	  "Name ei tohi olla tühi"
  ]
}
```

For setting the default locale used by Valgo, you can just use the function `SetDefaultLocale(code string)` as shown in the following example.

```go
v.SetDefaultLocale("es")
```

PRs are welcome if you want to add locale messages for your language, but please make sure that your translations are high quality.

# List of rules by validator type

- `String` validator
  - `EqualTo`
  - `GreaterThan`
  - `GreaterOrEqualTo`
  - `LessThan`
  - `LessOrEqualTo`
  - `Between`
  - `Empty`
  - `Blank`
  - `Passing`
  - `InSlice`
  - `MatchingTo`
  - `MaxLength`
  - `MinLength`
  - `Length`
  - `LengthBetween`

- `StringP` validator - for string pointer
  - `EqualTo`
  - `GreaterThan`
  - `GreaterOrEqualTo`
  - `LessThan`
  - `LessOrEqualTo`
  - `Between`
  - `Empty`
  - `Blank`
  - `Passing`
  - `InSlice`
  - `MatchingTo`
  - `MaxLength`
  - `MinLength`
  - `Length`
  - `LengthBetween`
  - `BlankOrNil`
  - `EmptyOrNil`
  - `Nil`

- `Bool` validator
  - `EqualTo`
  - `Passing`
  - `True`
  - `False`
  - `InSlice`

- `BoolP` validator - for boolean pointer
  - `EqualTo`
  - `Passing`
  - `True`
  - `False`
  - `InSlice`
  - `FalseOrNil`
  - `Nil`

- `Number` and `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Float32`, `Float64` - for number pointer
  - `EqualTo`
  - `GreaterThan`
  - `GreaterOrEqualTo`
  - `LessThan`
  - `LessOrEqualTo`
  - `Between`
  - `Zero`
  - `InSlice`
  - `Passing`
  
- `NumberP` and `IntP`, `Int8P`, `Int16P`, `Int32P`, `Int64P`, `UintP`, `Uint8P`, `Uint16P`, `Uint32P`, `Uint64P`, `Float32P`, `Float64P` - for number pointer
  - `EqualTo`
  - `GreaterThan`
  - `GreaterOrEqualTo`
  - `LessThan`
  - `LessOrEqualTo`
  - `Between`
  - `Zero`
  - `InSlice`
  - `Passing`
  - `ZeroOrNil`
  - `Nil`

- `Any` validator
  - `EqualTo`
  - `Passing`
  - `Nil`

# License

Copyright © 2022 Carlos Forero

Valgo is released under the [MIT License](LICENSE)