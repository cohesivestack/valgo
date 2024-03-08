package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorBoolPNot(t *testing.T) {

	bool1 := true

	v := Is(BoolP(&bool1).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPEqualToWhenIsValid(t *testing.T) {

	var v *Validation

	valTrue := true
	v = Is(BoolP(&valTrue).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	valFalse := false
	v = Is(BoolP(&valFalse).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPEqualToWhenIsInvalid(t *testing.T) {

	var v *Validation

	valTrue := true

	v = Is(BoolP(&valTrue).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.Errors()["value_0"].Messages()[0])

	valFalse := false
	v = Is(BoolP(&valFalse).EqualTo(true))
	assert.Equal(t,
		"Value 0 must be equal to \"true\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPTrueWhenIsValid(t *testing.T) {

	var v *Validation

	valTrue := true

	v = Is(BoolP(&valTrue).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(BoolP(&mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPTrueWhenIsInvalid(t *testing.T) {

	var v *Validation

	valFalse := false

	v = Is(BoolP(&valFalse).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(BoolP(&mybool1).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPFalseWhenIsValid(t *testing.T) {

	var v *Validation

	_valFalse := false
	valFalse := &_valFalse

	v = Is(BoolP(valFalse).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = false
	var mybool1 *MyBool = &_mybool1

	v = Is(BoolP(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPFalseWhenIsInvalid(t *testing.T) {

	var v *Validation

	_valTrue := true
	valTrue := &_valTrue

	v = Is(BoolP(valTrue).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = true
	var mybool1 *MyBool = &_mybool1

	v = Is(BoolP(mybool1).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolNilIsValid(t *testing.T) {

	var v *Validation

	var valTrue *bool

	v = Is(BoolP(valTrue).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 *MyBool

	v = Is(BoolP(mybool1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolNilIsInvalid(t *testing.T) {

	var v *Validation

	valTrue := true

	v = Is(BoolP(&valTrue).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(BoolP(&mybool1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPPassingWhenIsValid(t *testing.T) {

	var v *Validation

	valTrue := true

	v = Is(BoolP(&valTrue).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPPassingWhenIsInvalid(t *testing.T) {

	var v *Validation

	valFalse := false

	v = Is(BoolP(&valFalse).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPInSliceValid(t *testing.T) {

	var v *Validation

	boolValue := false

	v = Is(BoolP(&boolValue).InSlice([]bool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = false

	v = Is(BoolP(&myBool1).InSlice([]MyBool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPInSliceInvalid(t *testing.T) {

	var v *Validation

	_boolValue := true
	boolValue := &_boolValue

	v = Is(BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	boolValue = nil

	v = Is(BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = true

	v = Is(BoolP(&myBool1).InSlice([]MyBool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(BoolP(&_true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(BoolP(&_true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(BoolP(&_true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(BoolP(&_true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(BoolP(&_true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(BoolP(&_true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(BoolP(&_true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(BoolP(&_true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(BoolP(&_true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(BoolP(&_true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(BoolP(&_true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(BoolP(&_true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(BoolP(&_true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorBoolPOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(BoolP(&_true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(BoolP(&_true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(BoolP(&_true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(BoolP(&_true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(BoolP(&_true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(BoolP(&_true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(BoolP(&_true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(BoolP(&_true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(BoolP(&_true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(BoolP(&_true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(BoolP(&_true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(BoolP(&_true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(BoolP(&_true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
