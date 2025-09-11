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
  val := v.Is(
    v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20),
    v.Number(17, "age").GreaterThan(18),
  )

  if !val.Valid() {
    out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

Valgo is used in production by [Statsignal](https://statsignal.dev), but we want community feedback before releasing version 1.

## 🚨 Breaking Change in v0.6.0 — String Length Validation

Starting from **v0.6.0**, all string length validators now measure length in **characters (runes)** instead of bytes.
This means that multi-byte UTF-8 characters (such as Japanese/Chinese/Korean characters, accented letters, and other international characters) are now counted as one character each, making the validators more intuitive for international (i18n) applications.

### What Changed

* `MaxLength`, `MinLength`, `OfLength`, and `OfLengthBetween` now use `utf8.RuneCountInString`.

### Migration

If your code relied on **byte length** (using `len` semantics), use the new explicit byte-based validators:

* `MaxBytes`
* `MinBytes`
* `OfByteLength`
* `OfByteLengthBetween`

### Example

```go
s := "你好" // 2 characters (runes), 6 bytes
japanese := "こんにちは" // 5 characters (runes), 15 bytes

// New default: counts characters (runes)
v.String(s, "field").MaxLength(2)        // ✅ passes
v.String(japanese, "field").MaxLength(5) // ✅ passes

// Byte-based: counts bytes
v.String(s, "field").MaxBytes(6)         // ✅ passes
v.String(japanese, "field").MaxBytes(15) // ✅ passes
```

# Table of content

- [Valgo](#valgo)
- [Quick example](#quick-example)
- [v0.x.x and backward compatibility](#v0xx-and-backward-compatibility)
- [Table of content](#table-of-content)
- [Getting started](#getting-started)
- [Using Valgo](#using-valgo)
  - [`Validation` session](#validation-session)
  - [`Is(...)` function](#is-function)
  - [`Validation.Valid()` function](#validationvalid-function)
  - [`Validation.IsValid(...)` function](#validationisvalid-function)
  - [`In(...)` function](#in-function)
  - [`InRow(...)` function](#inrow-function)
  - [`InCell(...)` function](#incell-function)
  - [`Check(...)` function](#check-function)
  - [`If(...)` function](#if-function)
  - [`Do(...)` function](#do-function)
  - [`When(...)` function](#when-function)
  - [`AddErrorMessage(...)` function](#adderrormessage-function)
  - [Merging two `Validation` sessions with `Validation.Merge( ... )`](#merging-two-validation-sessions-with-validationmerge--)
  - [`New()` function](#new-function)
  - [Error handling functions](#error-handling-functions)
  - [Custom error message template](#custom-error-message-template)
  - [Localizing a validation session with New(...options) function](#localizing-a-validation-session-with-newoptions-function)
  - [Managing common options with Factory](#managing-common-options-with-factory)
  - [Custom errors JSON output with Factory](#custom-errors-json-output-with-factory)
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
  - [Time validator](#time-validator)
  - [Time pointer validator](#time-pointer-validator)
  - [Any validator](#any-validator)
  - [Custom type validators](#custom-type-validators)
- [Or Operator in Validators](#or-operator-in-validators)
  - [Overview](#overview)
  - [Usage](#usage)
  - [Key Points](#key-points)
  - [Examples](#examples)
- [Extending Valgo with custom validators](#extending-valgo-with-custom-validators)
- [List of rules by validator type](#list-of-rules-by-validator-type)
- [Github Code Contribution Guide](#github-code-contribution-guide)
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
  * `InCell(...)`,
  * `Check(...)`,
  * `AddErrorMessage(...)`

`Is(...)` is likely to be the most frequently used function in your validations. When `Is(...)` is called, the function creates a validation and receives a validator at the same time. In the next section, you will learn more about the `Is(...)` function.

## `Is(...)` function

The `Is(...)` function allows you to pass one or multiple `Validator`s, each with their respective values and rules for validation. This creates a `Validation` session, which can be used to validate multiple values.


In the following example, we pass multiple `Validator`s for the `full_name`, `age`, and `status` values to the `Is(...)` function:

```go
val := v.Is(
  v.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20),
  v.Number(17, "age").GreaterThan(18),
  v.String("singl", "status").InSlice([]string{"married", "single"})
)

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

## `Validation.Valid()` function

A `Validation` session provides this function, which returns either `true` if all their validators are valid or `false` if any one of them is invalid.

In the following example, even though the Validator for `age` is valid, the `Validator` for `status` is invalid, making the entire `Validator` session invalid.

```go
val := v.Is(
  v.Number(21, "age").GreaterThan(18),
  v.String("singl", "status").InSlice([]string{"married", "single"}),
)

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

This function allows checking if a specific value in a `Validation` session is valid or not. This is very useful for conditional logic.

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
  In("address", v.Is(
    String(p.Address.Name, "name").Not().Blank(),
    String(p.Address.Street, "street").Not().Blank(),
  ))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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
  val.InRow("addresses", i, v.Is(
    v.String(a.Name, "name").Not().Blank(),
    v.String(a.Street, "street").Not().Blank(),
  ))
}

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

## `InCell(...)` function

The `InCell(...)` function executes one or more validators in an indexed namespace where the target is a scalar value (e.g., entries of a primitive slice). The value names in the error result are prefixed with this indexed namespace. It is useful for validating lists of primitive values.

In the following example, we validate a list of tag names. The error results can distinguish the errors for each list entry.

```go
tags := []string{"", ""}

val := v.New()
for i, tag := range tags {
  val.InCell("tags", i, v.Is(
    v.String(tag, "name").Not().Blank(),
  ))
}

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
}
```
output:
```json
{
  "tags[0]": [
    "Name can't be blank"
  ],
  "tags[1]": [
    "Name can't be blank"
  ]
}
```

## `Check(...)` function

The `Check(...)` function, similar to the `Is(...)` function, however with `Check(...)` the Rules of the Validator parameter are not short-circuited, which means that regardless of whether a previous rule was valid, all rules are checked.

This example shows two rules that fail due to the empty value in the `full_name` `Validator`, and since the `Validator` is not short-circuited, both error messages are added to the error result.

```go
val := v.Check(v.String("", "full_name").Not().Blank().OfLengthBetween(4, 20))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

## `If(...)` function

The `If(...)` function is similar to `Merge(...)`, but merges the `Validation` session only when the condition is true, and returns the same `Validation` instance. When the condition is false, no operation is performed and the original instance is returned unchanged.

This function allows you to write validation code in a more fluent and compact way, especially useful for conditional merging of validation sessions without the need for separate if statements or complex branching logic.

```go

// Only merge admin validation if user is admin
val := v.
  Is(v.String(username, "username").Not().Blank()).
  If(isAdmin, v.Is(v.String(role, "role").EqualTo("admin")))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
}
```

## `Do(...)` function

The `Do(...)` function executes the given function with the current `Validation` instance and returns the same instance. This allows you to extend a validation chain with additional or conditional rules in a concise way.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  Do(func(val *v.Validation) {
    if isAdmin {
      val.Is(v.String(role, "role").EqualTo("admin"))
    }
  })

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
}
```

## `When(...)` function

The `When(...)` function is similar to `Do(...)`, but executes the given function only when the condition is true, and returns the same `Validation` instance. When the condition is false, no operation is performed and the original instance is returned unchanged.

This function provides a more concise way to add conditional validation logic compared to using `Do(...)` with an internal if statement.

```go
val := v.
  Is(v.String(username, "username").Not().Blank()).
  When(isAdmin, func(val *v.Validation) {
    val.Is(v.String(role, "role").EqualTo("admin"))
  })

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
}
```

## `AddErrorMessage(...)` function

The `AddErrorMessage` function allows to add an error message to a Validation session without executing a field validator. This function takes in two arguments: `name`, which is the name of the field for which the error message is being added, and `message`, which is the error message being added to the session.

When an error message is added using this function, the Validation session is marked as invalid, indicating that at least one validation error has occurred.

One use case for the `AddErrorMessage` function is to add a general error message for the validation of an entity structure. As shown in the example below, if you have an entity structure for an `address` and need to validate multiple fields within it, such as the `city` and `street`, you can use `AddErrorMessage` to include a general error message for the entire `address` in case any of the fields fail validation.

```go
type Address struct {
  City string
  Street string
}

a := Address{"", "1600 Amphitheatre Pkwy"}

val := v.Is(
  v.String(a.city, "city").Not().Blank(),
  v.String(a.Street, "street").Not().Blank(),
)

if !val.Valid() {
  v.AddErrorMessage("address", "The address is wrong!")

  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
}

```
output:
```json
{
  "address": [
    "The address is wrong!"
  ],
  "city": [
    "City can't be blank"
  ]
}
```

 It's worth noting that there may be other use cases for this function as well.

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

val := v.Is(
  v.String(r.Name, "name").Not().Blank(),
  v.String(r.Status, "status").Not().Blank(),
)

val.Merge(validatePreStatus(r.Status))

if !val.Valid() {
  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

## `New()` function

This function allows you to create a new `Validation` session without a `Validator`. This is useful for conditional validation or reusing validation logic.

The function accepts an optional parameter of type [Options] struct, which allows you to specify options such as the specific locale code and locale to use, and a custom JSON marshaler for errors.

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

As we mentioned above, you can pass the Options type to the New() function, in order to specify additional options when creating a new Validation session, such as the specific locale code and locale to use, and a custom JSON marshaler for errors. More information about the Options parameter in the following sections.

## Error handling functions

Valgo provides three functions for handling validation errors:

### `ToError()` function

The `ToError()` function returns the same value as `ToValgoError()` but as a standard Go `error` interface. This function is ideal for idiomatic error handling and integration with Go's native error system.

```go
val := v.Is(v.String("", "name").Not().Blank())

if err := val.ToError(); err != nil {
    log.Printf("Validation failed: %v", err)
    return err
}
```

### `ToValgoError()` function

The `ToValgoError()` function returns the same value as `ToError()` but as a concrete `*valgo.Error` type instead of the standard `error` interface. It's essentially a shortcut to `ToError().(*valgo.Error)`, providing access to rich, structured error details.



```go
val := v.Is(v.String("", "name").Not().Blank())

if errInfo := val.ToValgoError(); errInfo != nil {
    for field, valueError := range errInfo.Errors() {
        fmt.Printf("Field '%s': %v\n", field, valueError.Messages())
    }
}
```

### `Error()` function (Deprecated)

The `Error()` function is deprecated in favor of `ToError()` or `ToValgoError()`. The `Error()` function name conflicts with Go's error interface implementation convention, where `Error()` typically implements the error interface for a type. This function will be removed when Valgo reaches version 1.

```go
// DEPRECATED: Use ToError() or ToValgoError() instead
val := v.Is(v.String("", "name").Not().Blank())
if !val.Valid() {
    out, _ := json.MarshalIndent(val.Error(), "", "  ")
    fmt.Println(string(out))
}
```

### When to use each function

- **Use `ToError()`** for standard error handling and integration with Go's error system
- **Use `ToValgoError()`** when you need detailed validation information, per-field messages, or custom error processing
- **Avoid `Error()`** as it's deprecated and may be removed when Valgo reaches version 1

### Custom JSON marshaling

All three functions support custom JSON marshaling functions:

```go
customFunc := func(e *Error) ([]byte, error) {
    return []byte(`{"custom": "error"}`), nil
}

// Using custom marshaling with any of the functions
err := val.ToError(customFunc)
errInfo := val.ToValgoError(customFunc)
```

## Custom error message template

Customizing the default Valgo error messages is possible through the `New()` function as it's explained in the [Localizing a validation session with New](#localizing-a-validation-session-with-newoptions-function) section, however the Valgo validators allow to customize the template of a specific template validator rule. Below is an example illustrating this with the String empty validator rule.

```go
val := v.Is(v.String("", "address_field", "Address").Not().Empty("{{title}} must not be empty. Please provide the value in the input {{name}}."))

out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

## Localizing a validation session with New(...options) function

Valgo has localized error messages. The error messages are currently available in English (default), Spanish, German and Hungarian. However, it is possible to set error messages for any locale by passing the Options parameter to the `New()` function. Using this parameter, you can also customize the existing Valgo locale messages.

There are two options for localization: `localeCode` and `locale`. Below, we list the different ways to customize localization with these two parameters.

* Changing the validation session's locale
  In the following example, we are setting the Spanish locale:

  ```go
  // Creating the new validation session with other locale
  val := v.New(v.Options{ LocaleCode: "es" })

  // Testing the output
  val := val.Check(v.String(" ", "nombre").Not().Blank())

  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
  ```

  output:
  ```json
  {
    "name": [
      "Nombre no puede estar en blanco"
    ]
  }
  ```

  If the specified locale does not exist, Valgo's default English locale will be used. If you wish to change the default locale, you should use a Factory function, which is explained in the [Factory section](#managing-common-options-with-factory).

* Changing the locale entries
  In the example below, we are changing the entry for the "Not Blank" error. Since we are not specifying the `localeCode`, we are using and replacing the default English locale. However, you can also specify another localeCode if necessary.

  ```go
  // Creating a new validation session and changing a locale entry
  val := v.New(v.Options{
    Locale: &Locale{
		  ErrorKeyNotBlank: "{{title}} should not be blank",
	  }
  })

  // Testing the output
  val := val.Check(v.String(" ", "name").Not().Blank())

  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
  ```

  output:
  ```json
  {
    "name": [
      "Name should not be blank"
    ]
  }
  ```

* Adding a new locale
  As mentioned previously, Valgo currently only has the English, Spanish, German and Hungarian locales, but we hope to have more soon. However, you can add your own custom locale. Below is an example using the Estonian language:

  ```go
  // Creating a new validation session and adding a new locale with two entries
  val := v.New(v.Options{
    LocaleCode: "ee",
    Locale: &Locale{
		  ErrorKeyNotBlank: "{{title}} ei tohi olla tühi",
      ErrorKeyNotFalse: "{{title}} ei tohi olla vale",
	  }
  })

  // Testing the output
  val := val.Is(
    v.String(" ", "name").Not().Blank(),
    v.Bool(false, "active").Not().False(),
  )

  out, _ := json.MarshalIndent(val.ToError(), "", "  ")
  fmt.Println(string(out))
  ```

  output:
  ```json
  {
    "name": [
      "Name ei tohi olla tühi"
    ],
    "active": [
      "Active ei tohi olla vale"
    ]
  }
  ```

  For entries not specified in the custom locale, the default Valgo locale (English) will be used. If you wish to change the default locale, you can use the `Factory` function, which is further explained in the [Factory section](#managing-common-options-with-factory).

We welcome pull requests for adding new locale messages, but please ensure that the translations are of high quality.

## Managing common options with Factory

Valgo provides the `Factory()` function which allows you to create a valgo factory. With a valgo factory, you can create `Validation` sessions with preset options, avoiding having to pass options each time when a Validation is created. This allows more flexibility and easier management of options when creating `Validation` sessions.

The `Factory` function takes a parameter of type `FactoryOptions` struct, which allows you to modify the default locale code, add new locales, and set a custom JSON marshaler for errors. The `ValidationFactory` instance created by this function has all the functions to create Validations available in the package level (`Is()`, `In()`, `Check()`, `New()`) which creates a new Validation session with the preset options in the factory.

In the following example, we create a `Factory` with the default locale code set to Spanish, a new locale added for Estonian. This factory instance enables us to create validation sessions.

```go
factory := v.Factory(v.FactoryOptions{
  LocaleCodeDefault: "es",
  Locales: map[string]*Locale{
      "ee": {
          v.ErrorKeyNotBlank: "{{title}} ei tohi olla tühi",
          v.ErrorKeyNotFalse: "{{title}} ei tohi olla vale",
      },
  }
})

// Error will contain the spanish error "Nombre no puede estar en blanco"
v1 := factory.Is(String(" ", "nombre").NotBlank())

// Error will contain the spanish error "Nime ei tohi olla tühi"
v2 := factory.New(Options{LocaleCode: "ee"}).Is(String(" ", "nime").Not().Blank())
```

## Custom errors JSON output with Factory

It is possible to use the `MarshalJsonFunc` parameter of `Factory` for customizing the JSON output for errors.

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
factory := v.Factory(v.FactoryOptions{
  MarshalJsonFunc: customMarshalJson
})

// Now validate something to check if the output JSON contains the errors root key

val := factory.Is(v.String("", "name").Not().Empty())

out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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

# Validators

Validators establish the rules for validating a value. The validators are passed to a `Validation` session.

For each primitive Golang value type, Valgo provides a `Validator`. A `Validator` has different functions that set its value's validation rules.

Although Valgo has multiple types of validators, it can be extended with custom validators. Check the section [Extending Valgo with Custom Validators](#extending-valgo-with-custom-validators) for more information.

## Validator value's name and title.

Validators only require the value to be validated, so, for example, the following code validates a string value by checking if it is empty.

```go
val := v.New(v.String("").Empty())
```
`val.ToError()` output:
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
`val.ToError()` output:
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
`val.ToError()` output:
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
v.Is(v.String("processing").InSlice([]string{"idle", "processing", "ready"}))

// Byte-length based
v.Is(v.String("123456").MaxBytes(6))
v.Is(v.String("123").MinBytes(3))
v.Is(v.String("1234").OfByteLength(4))
v.Is(v.String("12345").OfByteLengthBetween(4,6)) // Inclusive

// Rune-length based (unicode code points, tends to matter for languages that use non-Latin alphabet)
// 虎視眈々 is 4 runes/characters, but len(x) = 12 bytes
v.Is(v.String("虎視眈々").MaxLength(4))
v.Is(v.String("虎視眈々").MinLength(4))
v.Is(v.String("虎視眈々").OfLength(4))
v.Is(v.String("虎視眈々").OfLengthBetween(2,4)) // Inclusive

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
x := "processing";     v.Is(v.StringP(&x).InSlice([]string{"idle", "processing", "ready"}))

// Byte-length based
x := "123456";         v.Is(v.StringP(&x).MaxBytes(6))
x := "123";            v.Is(v.StringP(&x).MinBytes(3))
x := "1234";           v.Is(v.StringP(&x).OfByteLength(4))
x := "12345";          v.Is(v.StringP(&x).OfByteLengthBetween(4,6)) // Inclusive

// Rune-length based (counts characters instead of bytes)
// 虎視眈々 is 4 runes/characters, but len(x) = 12 bytes
x := "虎視眈々";      v.Is(v.StringP(&x).MaxLength(4))
x := "虎視眈々";      v.Is(v.StringP(&x).MinLength(4))
x := "虎視眈々";      v.Is(v.StringP(&x).OfLength(4))
x := "虎視眈々";      v.Is(v.StringP(&x).OfLengthBetween(2,4)) // Inclusive

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
Byte(v byte)       ByteP(v *byte)
Rune(v rune)       RuneP(v *rune)
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

## Time validator

The `ValidatorTime` provides functions for setting validation rules for a `time.Time` type value, or a custom type based on a `time.Time`.

Below is a valid example for every Time validator rule.

```go
import "time"

v.Is(v.Time(time.Now()).EqualTo(time.Now()))
v.Is(v.Time(time.Now()).After(time.Now().Add(-time.Hour)))
v.Is(v.Time(time.Now()).AfterOrEqualTo(time.Now().Add(-time.Hour)))
v.Is(v.Time(time.Now()).Before(time.Now().Add(time.Hour)))
v.Is(v.Time(time.Now()).BeforeOrEqualTo(time.Now().Add(time.Hour)))
v.Is(v.Time(time.Now()).Between(time.Now().Add(-time.Hour), time.Now().Add(2*time.Hour))) // Inclusive
v.Is(v.Time(time.Time{}).Zero())
v.Is(v.Time(time.Now()).Passing(func(val time.Time) bool { return val.Before(time.Now().Add(2*time.Hour)) }))
v.Is(v.Time(time.Now()).InSlice([]time.Time{time.Now(), time.Now().Add(time.Hour)}))
```

## Time pointer validator

The `ValidatorTimeP` provides functions for setting validation rules for a `time.Time` type pointer, or a custom type based on a `time.Time` pointer.

Below is a valid example for every Time pointer validator rule.

```go
import "time"

x := time.Now(); v.Is(v.TimeP(&x).EqualTo(time.Now()))
x = time.Now(); v.Is(v.TimeP(&x).After(time.Now().Add(-time.Hour)))
x = time.Now(); v.Is(v.TimeP(&x).AfterOrEqualTo(time.Now().Add(-time.Hour)))
x = time.Now(); v.Is(v.TimeP(&x).Before(time.Now().Add(time.Hour)))
x = time.Now(); v.Is(v.TimeP(&x).BeforeOrEqualTo(time.Now().Add(time.Hour)))
x = time.Now(); v.Is(v.TimeP(&x).Between(time.Now().Add(-time.Hour), time.Now().Add(2*time.Hour))) // Inclusive
x = time.Time{}; v.Is(v.TimeP(&x).Zero())
x = time.Now(); v.Is(v.TimeP(&x).Passing(func(val *time.Time) bool { return val.Before(time.Now().Add(2*time.Hour)) }))
x = time.Now(); v.Is(v.TimeP(&x).InSlice([]time.Time{time.Now(), time.Now().Add(time.Hour)}))
var x *time.Time; v.Is(v.TimeP(x).Nil())
x = new(time.Time); v.Is(v.TimeP(x).NilOrZero())
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

All golang validators allow to pass a custom type based on its value type. Below some valid examples.

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

# Or Operator in Validators

The `Or` operator function enables developers to combine validator rules using a logical OR chain. This addition allows for more nuanced validator scenarios, where a value may satisfy one of multiple conditions to be considered valid.

## Overview

In Valgo, validator rules are typically chained together using an implicit AND logic. This means that for a value to be deemed valid, it must satisfy all specified conditions. The `Or` operator provides an alternative by allowing conditions to be linked with OR logic. In such cases, a value is considered valid if it meets at least one of the chained conditions.

The `Or` operator follows a simple left-to-right boolean priority, akin to the Go language's approach to evaluating boolean expressions. Valgo does not have an equivalent to parentheses in API functions, in order to keep the syntax simple and readable. We believe that complex boolean logic becomes harder to read with a Fluent API interface, so for those cases, it is preferred to use imperative Go programming language constructs.

## Usage

To utilize the `Or` operator, simply insert `.Or().` between two conditions within your validator chain. Here's a basic example:

```go
v := Is(Bool(true).True().Or().False())
```

In this case, the validator passes because the boolean value `true` satisfies the first condition before the `Or()` operator.

## Key Points

- **Implicit AND Logic**: By default, when validators are chained without specifying the `Or()` operator, they are combined using an AND logic. Each condition must be met for the validation to pass.
- **No Short-circuiting for `Check`**: Unlike the `Is` function, which evaluates conditions lazily and may short-circuit (stop evaluating once the overall outcome is determined), the `Check` function ensures that all conditions are evaluated, regardless of their order and the use of `Or`.

## Examples

Below are examples demonstrating different scenarios using the `Or` operator, including combinations with the `Not` operator and multiple `Or` conditions in sequence. These examples illustrate how you can tailor complex validation logic to suit your needs.

```go
// Validation with two valid OR conditions
v = Is(Bool(true).True().Or().True())
assert.True(t, v.Valid())

// Validation with a valid OR condition followed by an invalid AND condition
v = Is(Bool(true).False().Or().True().False())
assert.False(t, v.Valid())

// Validation combining NOT and OR operators
v = Is(Bool(true).Not().False().Or().False())
assert.True(t, v.Valid())
```

These examples are intended to provide a clear understanding of how to effectively use the `Or` operator in your validations. By leveraging this functionality, you can create more flexible and powerful validation rules, enhancing the robustness and usability of your applications.

# Extending Valgo with custom validators

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

out, _ := json.MarshalIndent(val.ToError(), "", "  ")
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
  - `MaxBytes`
  - `MinBytes`
  - `OfByteLength`
  - `OfByteLengthBetween`
  - `MaxLength`
  - `MinLength`
  - `OfLength`
  - `OfLengthBetween`

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
  - `MaxBytes`
  - `MinBytes`
  - `OfByteLength`
  - `OfByteLengthBetween`
  - `MaxLength`
  - `MinLength`
  - `OfLength`
  - `OfLengthBetween`
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

- `Number` and `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Float32`, `Float64`, `Byte`, `Rune` - for number pointer
  - `EqualTo`
  - `GreaterThan`
  - `GreaterOrEqualTo`
  - `LessThan`
  - `LessOrEqualTo`
  - `Between`
  - `Zero`
  - `InSlice`
  - `Passing`
  
- `NumberP` and `IntP`, `Int8P`, `Int16P`, `Int32P`, `Int64P`, `UintP`, `Uint8P`, `Uint16P`, `Uint32P`, `Uint64P`, `Float32P`, `Float64P`, `ByteP`, `RuneP` - for number pointer
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

- `Time` validator
  - `EqualTo`
  - `After`
  - `AfterOrEqualTo`
  - `Before`
  - `BeforeOrEqualTo`
  - `Between`
  - `Zero`
  - `Passing`
  - `InSlice`
 
- `TimeP` validator - for `time.Time` pointer
  - `EqualTo`
  - `After`
  - `AfterOrEqualTo`
  - `Before`
  - `BeforeOrEqualTo`
  - `Between`
  - `Zero`
  - `Passing`
  - `InSlice`
  - `Nil`
  - `NilOrZero`

- `Any` validator
  - `EqualTo`
  - `Passing`
  - `Nil`

# Github Code Contribution Guide

We welcome contributions to our project! To make the process smooth and efficient, please follow these guidelines when submitting code:

* **Discuss changes with the community**: We encourage contributors to discuss their proposed changes or improvements with the [community](https://github.com/cohesivestack/valgo/discussions/categories/ideas) before starting to code. This ensures that the changes align with the focus and purpose of the project, and that other contributors are aware of the work being done.

* **Make commits small and cohesive**: It is important to keep your commits focused on a single task or change. This makes it easier to review and understand your changes.

* **Check code formatting with go fmt**: Before submitting your code, please ensure that it is properly formatted using the go fmt command.

* **Make tests to cover your changes**: Please include tests that cover the changes you have made. This ensures that your code is functional and reduces the likelihood of bugs.

* **Update golang docs and README to cover your changes**: If you have made changes that affect documentation or the README file, please update them accordingly.

* **Keep a respectful language with a collaborative tune**: We value a positive and collaborative community. Please use respectful language when communicating with other contributors or maintainers.

# License

Copyright © 2025 Carlos Forero

Valgo is released under the [MIT License](LICENSE)
