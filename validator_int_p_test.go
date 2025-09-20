package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorIntPConstructor(t *testing.T) {

	numberInt := int(1)

	v := Is(IntP(&numberInt).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	numberInt8 := int8(1)

	v = Is(Int8P(&numberInt8).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	numberInt16 := int16(1)

	v = Is(Int16P(&numberInt16).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	numberInt32 := int32(1)

	v = Is(Int32P(&numberInt32).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	numberInt64 := int64(1)

	v = Is(Int64P(&numberInt64).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPNot(t *testing.T) {

	number1 := int(2)

	v := Is(IntP(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPEqualToValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(2)
	number1 := &_number1

	v = Is(IntP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(IntP(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPGreaterThanValid(t *testing.T) {
	var v *Validation

	number1 := int(3)

	v = Is(IntP(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPGreaterThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(2)
	number1 := &_number1

	v = Is(IntP(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 2

	v = Is(IntP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = int(3)

	v = Is(IntP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(2)
	number1 := &_number1

	v = Is(IntP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(IntP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPLessThanValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(IntP(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPLessThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(2)
	number1 := &_number1

	v = Is(IntP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = int(3)

	v = Is(IntP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPLessOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(IntP(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(3)
	number1 := &_number1

	v = Is(IntP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil
	v = Is(IntP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(IntP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPBetweenValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = int(4)

	v = Is(IntP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = int(6)

	v = Is(IntP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = Is(IntP(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPBetweenInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(2)
	number1 := &_number1

	v = Is(IntP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = int(7)

	v = Is(IntP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = Is(IntP(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPZeroValid(t *testing.T) {
	var v *Validation

	_number1 := int(0)
	number1 := &_number1

	v = Is(IntP(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 0

	v = Is(IntP(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPZeroInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(1)
	number1 := &_number1

	v = Is(IntP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(IntP(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPZeroOrNilValid(t *testing.T) {
	var v *Validation

	_number1 := int(0)
	number1 := &_number1

	v = Is(IntP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = Is(IntP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 0

	v = Is(IntP(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPZeroOrNilInvalid(t *testing.T) {
	var v *Validation

	number1 := int(1)

	v = Is(IntP(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(IntP(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPPositiveValid(t *testing.T) {
	var v *Validation

	number1 := int(1)
	v = Is(IntP(&number1).Positive())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPPositiveInvalid(t *testing.T) {
	var v *Validation

	number1 := int(0)
	v = Is(IntP(&number1).Positive())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be positive",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPNegativeValid(t *testing.T) {
	var v *Validation

	number1 := int(-1)
	v = Is(IntP(&number1).Negative())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPNegativeInvalid(t *testing.T) {
	var v *Validation

	number1 := int(0)
	v = Is(IntP(&number1).Negative())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be negative",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPPassingValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = Is(IntP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPPassingInvalid(t *testing.T) {
	var v *Validation

	number1 := int(1)

	v = Is(IntP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(IntP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPInSliceValid(t *testing.T) {
	var v *Validation

	number1 := int(2)

	v = Is(IntP(&number1).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = Is(IntP(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPInSliceInvalid(t *testing.T) {
	var v *Validation

	_number1 := int(1)
	number1 := &_number1

	v = Is(IntP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(IntP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(IntP(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPNilIsValid(t *testing.T) {
	var v *Validation

	var valNumber *int

	v = Is(IntP(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 *MyNumber

	v = Is(IntP(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPNilIsInvalid(t *testing.T) {
	var v *Validation

	valNumber := int(1)

	v = Is(IntP(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = Is(IntP(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false
	var one = int(1)

	// Testing Or operation with two valid conditions
	v = Is(IntP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(IntP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(IntP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(IntP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(IntP(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(IntP(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(IntP(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(IntP(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(IntP(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(IntP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(IntP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(IntP(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(IntP(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorIntPOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	var one = int(1)

	// Testing Or operation with two valid conditions
	v = Check(IntP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(IntP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(IntP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(IntP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(IntP(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(IntP(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(IntP(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(IntP(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(IntP(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(IntP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(IntP(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(IntP(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(IntP(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
