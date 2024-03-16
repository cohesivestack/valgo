package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorNumberNot(t *testing.T) {

	v := Is(Number(1).Not().EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(1).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberGreaterThanValid(t *testing.T) {
	var v *Validation

	v = Is(Number(3).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberGreaterThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Number(2).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Number(3).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(Number(myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberLessThanValid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(Number(myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberLessThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Number(3).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberLessOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Number(1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(3).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(Number(myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberBetweenValid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Number(4).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Number(6).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = Is(Number(myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberBetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(2).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Number(7).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = Is(Number(myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberZeroValid(t *testing.T) {
	var v *Validation

	v = Is(Number(0).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 0

	v = Is(Number(myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorNumberZeroInvalid(t *testing.T) {
	var v *Validation

	v = Is(Number(1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Number(0.1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(Number(myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberPassingValid(t *testing.T) {

	var v *Validation

	v = Is(Number(1).Passing(func(val int) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(Number(myNumber1).Passing(func(val MyNumber) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPassingInvalid(t *testing.T) {

	var v *Validation

	v = Is(Number(1).Passing(func(val int) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(Number(myNumber1).Passing(func(val MyNumber) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberInSliceValid(t *testing.T) {

	var v *Validation

	v = Is(Number(2).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = Is(Number(myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberInSliceInvalid(t *testing.T) {

	var v *Validation

	v = Is(Number(4).InSlice([]int{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 4

	v = Is(Number(myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorNumberOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Number(1).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Number(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Number(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Number(1).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Number(1).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Number(1).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Number(1).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Number(1).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Number(1).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Number(1).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Number(1).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Number(1).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Number(1).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorNumberOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Number(1).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Number(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Number(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Number(1).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Number(1).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Number(1).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Number(1).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Number(1).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Number(1).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Number(1).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Number(1).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Number(1).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Number(1).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
