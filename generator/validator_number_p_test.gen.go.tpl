// Code generated by Valgo; DO NOT EDIT.
package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

{{ range . -}}
func TestValidator{{ .Name }}PNot(t *testing.T) {

	number1 := {{ .Type }}(2)

	v := Is({{ .Name }}P(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PEqualToValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is({{ .Name }}P(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PGreaterThanValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(3)

	v = Is({{ .Name }}P(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PGreaterThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 2

	v = Is({{ .Name }}P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(3)

	v = Is({{ .Name }}P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is({{ .Name }}P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PLessThanValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is({{ .Name }}P(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PLessThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = {{ .Type }}(3)

	v = Is({{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PLessOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is({{ .Name }}P(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(3)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil
	v = Is({{ .Name }}P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PBetweenValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(4)

	v = Is({{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(6)

	v = Is({{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = Is({{ .Name }}P(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PBetweenInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = {{ .Type }}(7)

	v = Is({{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = Is({{ .Name }}P(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PZeroValid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(0)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 0

	v = Is({{ .Name }}P(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PZeroInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(1)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = Is({{ .Name }}P(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PZeroOrNilValid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(0)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = Is({{ .Name }}P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 0

	v = Is({{ .Name }}P(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PZeroOrNilInvalid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(1)

	v = Is({{ .Name }}P(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = Is({{ .Name }}P(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PPassingValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).Passing(func(val *{{ .Type }}) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PPassingInvalid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(1)

	v = Is({{ .Name }}P(&number1).Passing(func(val *{{ .Type }}) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = Is({{ .Name }}P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PInSliceValid(t *testing.T) {
	var v *Validation

	number1 := {{ .Type }}(2)

	v = Is({{ .Name }}P(&number1).InSlice([]{{ .Type }}{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2

	v = Is({{ .Name }}P(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PInSliceInvalid(t *testing.T) {
	var v *Validation

	_number1 := {{ .Type }}(1)
	number1 := &_number1

	v = Is({{ .Name }}P(number1).InSlice([]{{ .Type }}{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is({{ .Name }}P(number1).InSlice([]{{ .Type }}{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = Is({{ .Name }}P(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PNilIsValid(t *testing.T) {
	var v *Validation

	var valNumber *{{ .Type }}

	v = Is({{ .Name }}P(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 *MyNumber

	v = Is({{ .Name }}P(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PNilIsInvalid(t *testing.T) {
	var v *Validation

	valNumber := {{ .Type }}(1)

	v = Is({{ .Name }}P(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = Is({{ .Name }}P(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}POrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false
	var one = {{ .Type }}(1)

	// Testing Or operation with two valid conditions
	v = Is({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is({{ .Name }}P(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is({{ .Name }}P(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is({{ .Name }}P(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is({{ .Name }}P(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is({{ .Name }}P(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidator{{ .Name }}POrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	var one = {{ .Type }}(1)

	// Testing Or operation with two valid conditions
	v = Check({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check({{ .Name }}P(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check({{ .Name }}P(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check({{ .Name }}P(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check({{ .Name }}P(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check({{ .Name }}P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check({{ .Name }}P(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check({{ .Name }}P(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

{{ end }}