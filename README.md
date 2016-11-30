# How to use



```go
var email = "elonmusk@tesla.com"

valid := valgo.Is(name).NotBlank().Email().Valid()

var sku = "MAN-WH-001"

errors := valgo.Is(sku).Match("[A-Z]{3}-[A-Z]{2}-[0-9]{3}").Errors()

var name = "Elon Musk"

valid, errors := valgo.Is(sku).Match("[A-Z]{3}-[A-Z]{2}-[0-9]{3}").ValidAndErrors()
```





# Validation types

## Empty

```go
var email = ""
valid := valgo.Is(email).Empty().Valid()
fmt.Println(valid) // true

var email = " "
valid := valgo.Is(email).Empty().Valid()
fmt.Println(valid) // false
```

## NotEmpty

```go
var email = ""
valid := valgo.Is(email).NotEmpty().Valid()
fmt.Println(valid) // false

var email = " "
valid := valgo.Is(email).NotEmpty().Valid()
fmt.Println(valid) // true
```

## Blank

```go
var email = " "
valid := valgo.Is(email).Blank().Valid()
fmt.Println(valid) // true

var email = ""
valid := valgo.Is(email).Blank().Valid()
fmt.Println(valid) // true
```

## NotBlank

```go
var email = " "
valid := valgo.Is(email).NotBlank().Valid()
fmt.Println(valid) // false

var email = ""
valid := valgo.Is(email).NotBlank().Valid()
fmt.Println(valid) // false
```

## Match

```go
var sku = "MAN-WH-001"
valid := valgo.Is(sku).Match("[A-Z]{3}-[A-Z]{2}-[0-9]{3}").Valid()
fmt.Println(valid) // true
```

## NotMatch

```go
var sku = "MAN-WH-001"
valid := valgo.Is(sku).NotMatch("[A-Z]{3}-[A-Z]{2}-[0-9]{3}").Valid()
fmt.Println(valid) // false
```

