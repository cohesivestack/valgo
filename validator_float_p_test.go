package valgo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorFloatPConstructor(t *testing.T) {
	numberFloat32 := float32(1)

	v := Is(Float32P(&numberFloat32).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	numberFloat64 := float64(1)

	v = Is(Float64P(&numberFloat64).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloatPNot(t *testing.T) {

	number1 := float64(2)

	v := Is(Float64P(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloatPEqualToValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(2)
	number1 := &_number1

	v = Is(Float64P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(Float64P(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPGreaterThanValid(t *testing.T) {
	var v *Validation

	number1 := float64(3)

	v = Is(Float64P(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloatPGreaterThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(2)
	number1 := &_number1

	v = Is(Float64P(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 2

	v = Is(Float64P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = float64(3)

	v = Is(Float64P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(2)
	number1 := &_number1

	v = Is(Float64P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(Float64P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPLessThanValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = Is(Float64P(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPLessThanInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(2)
	number1 := &_number1

	v = Is(Float64P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = float64(3)

	v = Is(Float64P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPLessOrEqualToValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64P(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(3)
	number1 := &_number1

	v = Is(Float64P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil
	v = Is(Float64P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = Is(Float64P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPBetweenValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = float64(4)

	v = Is(Float64P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = float64(6)

	v = Is(Float64P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = Is(Float64P(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPBetweenInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(2)
	number1 := &_number1

	v = Is(Float64P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = float64(7)

	v = Is(Float64P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = Is(Float64P(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPZeroValid(t *testing.T) {
	var v *Validation

	_number1 := float64(0)
	number1 := &_number1

	v = Is(Float64P(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 0

	v = Is(Float64P(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPZeroInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(1)
	number1 := &_number1

	v = Is(Float64P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 1

	v = Is(Float64P(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPZeroOrNilValid(t *testing.T) {
	var v *Validation

	_number1 := float64(0)
	number1 := &_number1

	v = Is(Float64P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = Is(Float64P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 0

	v = Is(Float64P(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPZeroOrNilInvalid(t *testing.T) {
	var v *Validation

	number1 := float64(1)

	v = Is(Float64P(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 1

	v = Is(Float64P(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPPassingValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).Passing(func(val *float64) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2

	v = Is(Float64P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloatPPassingInvalid(t *testing.T) {
	var v *Validation

	number1 := float64(1)

	v = Is(Float64P(&number1).Passing(func(val *float64) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 1

	v = Is(Float64P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPInSliceValid(t *testing.T) {
	var v *Validation

	number1 := float64(2)

	v = Is(Float64P(&number1).InSlice([]float64{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 2

	v = Is(Float64P(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloatPInSliceInvalid(t *testing.T) {
	var v *Validation

	_number1 := float64(1)
	number1 := &_number1

	v = Is(Float64P(number1).InSlice([]float64{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = Is(Float64P(number1).InSlice([]float64{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 1

	v = Is(Float64P(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPNilIsValid(t *testing.T) {
	var v *Validation

	var valNumber *float64

	v = Is(Float64P(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber float64
	var myNumber1 *MyNumber

	v = Is(Float64P(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloatPNilIsInvalid(t *testing.T) {
	var v *Validation

	valNumber := float64(1)

	v = Is(Float64P(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber float64
	var myNumber1 MyNumber = 1

	v = Is(Float64P(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloatPOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false
	var one = float64(1)

	// Testing Or operation with two valid conditions
	v = Is(Float64P(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Float64P(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Float64P(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Float64P(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Float64P(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Float64P(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Float64P(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Float64P(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Float64P(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Float64P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Float64P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Float64P(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Float64P(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorFloatPOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	var one = float64(1)

	// Testing Or operation with two valid conditions
	v = Check(Float64P(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Float64P(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Float64P(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Float64P(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Float64P(&one).EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Float64P(&one).Not().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Float64P(&one).Not().EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Float64P(&one).EqualTo(1).Or().EqualTo(0).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Float64P(&one).EqualTo(0).Or().EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Float64P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Float64P(&one).EqualTo(0).Or().EqualTo(1).EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Float64P(&one).EqualTo(1).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Float64P(&one).EqualTo(1).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PNaNValid(t *testing.T) {
	var v *Validation

	number1 := float64(math.NaN())
	v = Is(Float64P(&number1).NaN())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PNaNInvalid(t *testing.T) {
	var v *Validation

	number1 := float64(1)
	v = Is(Float64P(&number1).NaN())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be NaN",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64PInfiniteValid(t *testing.T) {
	var v *Validation

	number1 := float64(math.Inf(1))
	v = Is(Float64P(&number1).Infinite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number2 := float64(math.Inf(-1))
	v = Is(Float64P(&number2).Infinite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PInfiniteInvalid(t *testing.T) {
	var v *Validation

	number1 := float64(1)
	v = Is(Float64P(&number1).Infinite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be infinite",
		v.Errors()["value_0"].Messages()[0])

	number2 := float64(-1)
	v = Is(Float64P(&number2).Infinite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be infinite",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64PFiniteValid(t *testing.T) {
	var v *Validation

	number1 := float64(1)
	v = Is(Float64P(&number1).Finite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PFiniteInvalid(t *testing.T) {
	var v *Validation

	number1 := float64(math.NaN())
	v = Is(Float64P(&number1).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])

	number2 := float64(math.Inf(1))
	v = Is(Float64P(&number2).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])

	number3 := float64(math.Inf(-1))
	v = Is(Float64P(&number3).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])
}
