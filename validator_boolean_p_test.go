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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
	ResetMessages()

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
