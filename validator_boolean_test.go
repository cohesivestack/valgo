package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorBoolNot(t *testing.T) {

	v := Is(Bool(true).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsValid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Bool(false).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(Bool(mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsInvalid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = false

	v = Is(Bool(mybool1).EqualTo(mybool2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolTrueWhenIsValid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(Bool(mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolTrueWhenIsInvalid(t *testing.T) {

	var v *Validation

	v = Is(Bool(false).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolFalseWhenIsValid(t *testing.T) {

	var v *Validation

	v = Is(Bool(false).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolFalseWhenIsInvalid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(Bool(mybool1).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPassingWhenIsValid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).Passing(func(val bool) bool {
		return val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(Bool(mybool1).Passing(func(val MyBool) bool {
		return val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPassingWhenIsInvalid(t *testing.T) {

	var v *Validation

	v = Is(Bool(false).Passing(func(val bool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).Passing(func(val MyBool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolInSliceValid(t *testing.T) {

	var v *Validation

	v = Is(Bool(false).InSlice([]bool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = false

	v = Is(Bool(myBool1).InSlice([]MyBool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolInSliceInvalid(t *testing.T) {

	var v *Validation

	v = Is(Bool(true).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = true

	v = Is(Bool(myBool1).InSlice([]MyBool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Bool(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Bool(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Bool(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Bool(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Bool(true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Bool(true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Bool(true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Bool(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Bool(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Bool(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Bool(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Bool(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Bool(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorBoolOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Bool(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Bool(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Bool(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Bool(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Bool(true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Bool(true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Bool(true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Bool(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Bool(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Bool(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Bool(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Bool(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Bool(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
