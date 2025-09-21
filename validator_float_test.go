package valgo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorFloat64Not(t *testing.T) {
	v := Is(Float32(float32(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64EqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).EqualTo(myFloat642))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64EqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 1
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).EqualTo(myFloat642))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64GreaterThanValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(3)).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 3
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).GreaterThan(myFloat642))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64GreaterThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(2)).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).GreaterThan(myFloat642))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64GreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(3)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).GreaterOrEqualTo(myFloat642))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64GreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 3

	v = Is(Float64(MyFloat64(myFloat641)).GreaterOrEqualTo(myFloat642))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64LessThanValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 3

	v = Is(Float64(MyFloat64(myFloat641)).LessThan(myFloat642))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64LessThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(3)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).LessThan(myFloat642))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64LessOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(1)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).LessOrEqualTo(myFloat642))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64LessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(3)).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 3
	var myFloat642 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).LessOrEqualTo(myFloat642))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64BetweenValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(4)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(6)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 2
	var myFloat643 MyFloat64 = 6

	v = Is(Float64(MyFloat64(myFloat641)).Between(myFloat642, myFloat643))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64BetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(7)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2
	var myFloat642 MyFloat64 = 3
	var myFloat643 MyFloat64 = 6

	v = Is(Float64(MyFloat64(myFloat641)).Between(myFloat642, myFloat643))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64ZeroValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(0)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 0

	v = Is(Float64(float64(myFloat641)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorFloat64ZeroInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 1

	v = Is(Float64(float64(myFloat641)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64PositiveValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Positive())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PositiveInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(0)).Positive())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be positive",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(-1)).Positive())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be positive",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64NegativeValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(-1)).Negative())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64NegativeInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(0)).Negative())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be negative",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(1)).Negative())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be negative",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64PassingValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Passing(func(val float64) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 1

	v = Is(Float64(MyFloat64(myFloat641)).Passing(func(val MyFloat64) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64PassingInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Passing(func(val float64) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 1

	v = Is(Float64(MyFloat64(myFloat641)).Passing(func(val MyFloat64) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64InSliceValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(2)).InSlice([]float64{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 2

	v = Is(Float64(MyFloat64(myFloat641)).InSlice([]MyFloat64{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64InSliceInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(4)).InSlice([]float64{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyFloat64 float64
	var myFloat641 MyFloat64 = 4

	v = Is(Float64(MyFloat64(myFloat641)).InSlice([]MyFloat64{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64OrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Float64(float64(1)).Not().EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Float64(float64(1)).Not().EqualTo(float64(1)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)).EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)).EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorFloat64OrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Float64(float64(1)).Not().EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Float64(float64(1)).Not().EqualTo(float64(1)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Float64(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)).EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Float64(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)).EqualTo(float64(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(1)).Or().EqualTo(float64(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Float64(float64(1)).EqualTo(float64(1)).EqualTo(float64(0)).Or().EqualTo(float64(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64NaNValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(math.NaN())).NaN())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64NaNInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).NaN())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be NaN",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64InfiniteValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(math.Inf(1))).Infinite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(math.Inf(-1))).Infinite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64InfiniteInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Infinite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be infinite",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(0)).Infinite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be infinite",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(-1)).Infinite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be infinite",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorFloat64FiniteValid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(1)).Finite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(0)).Finite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Float64(float64(-1)).Finite())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorFloat64FiniteInvalid(t *testing.T) {
	var v *Validation

	v = Is(Float64(float64(math.NaN())).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(math.Inf(1))).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Float64(float64(math.Inf(-1))).Finite())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be finite",
		v.Errors()["value_0"].Messages()[0])
}
