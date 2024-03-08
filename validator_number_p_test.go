package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorNumberPNot(t *testing.T) {

	number1 := 2

	v := Is(NumberP(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPEqualToValid(t *testing.T) {
	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := 2
	number1 := &_number1

	v = Is(NumberP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(NumberP(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPGreaterThanValid(t *testing.T) {
	var v *Validation

	number1 := 3

	v = Is(NumberP(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPGreaterThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := 2
	number1 := &_number1

	v = Is(NumberP(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 2

	v = Is(NumberP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 3

	v = Is(NumberP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := 2
	number1 := &_number1

	v = Is(NumberP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(NumberP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPLessThanValid(t *testing.T) {
	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(NumberP(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPLessThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := 2
	number1 := &_number1

	v = Is(NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 3

	v = Is(NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPLessOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(NumberP(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := 3
	number1 := &_number1

	v = Is(NumberP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil
	v = Is(NumberP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(NumberP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPBetweenValid(t *testing.T) {
	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 4

	v = Is(NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 6

	v = Is(NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = Is(NumberP(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPBetweenInvalid(t *testing.T) {
	var v *Validation

	_number1 := 2
	number1 := &_number1

	v = Is(NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 7

	v = Is(NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = Is(NumberP(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPZeroValid(t *testing.T) {
	var v *Validation

	_number1 := 0
	number1 := &_number1

	v = Is(NumberP(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	_number2 := 0.0
	number2 := &_number2

	v = Is(NumberP(number2).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 0

	v = Is(NumberP(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPZeroInvalid(t *testing.T) {
	var v *Validation

	_number1 := 1
	number1 := &_number1

	v = Is(NumberP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(NumberP(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPZeroOrNilValid(t *testing.T) {
	var v *Validation

	_number1 := 0
	number1 := &_number1

	v = Is(NumberP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = Is(NumberP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 0

	v = Is(NumberP(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPZeroOrNilInvalid(t *testing.T) {
	var v *Validation

	number1 := 1

	v = Is(NumberP(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(NumberP(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPPassingValid(t *testing.T) {

	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = Is(NumberP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPPassingInvalid(t *testing.T) {

	var v *Validation

	number1 := 1

	v = Is(NumberP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(NumberP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPInSliceValid(t *testing.T) {

	var v *Validation

	number1 := 2

	v = Is(NumberP(&number1).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = Is(NumberP(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberPInSliceInvalid(t *testing.T) {

	var v *Validation

	_number1 := 1
	number1 := &_number1

	v = Is(NumberP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(NumberP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(NumberP(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberNilIsValid(t *testing.T) {

	var v *Validation

	var valNumber *int

	v = Is(NumberP(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 *MyNumber

	v = Is(NumberP(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberNilIsInvalid(t *testing.T) {

	var v *Validation

	valNumber := 1

	v = Is(NumberP(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(NumberP(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false
	var one = 1

	// Testing Or operation with two valid conditions
	v = Is(NumberP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(NumberP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(NumberP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(NumberP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(NumberP(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(NumberP(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(NumberP(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(NumberP(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(NumberP(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(NumberP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(NumberP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(NumberP(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(NumberP(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorNumberPOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	var one = 1

	// Testing Or operation with two valid conditions
	v = Check(NumberP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(NumberP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(NumberP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(NumberP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(NumberP(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(NumberP(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(NumberP(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(NumberP(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(NumberP(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(NumberP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(NumberP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(NumberP(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(NumberP(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
