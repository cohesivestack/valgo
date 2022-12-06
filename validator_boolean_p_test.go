package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorBoolPNot(t *testing.T) {
	t.Parallel()

	bool1 := true

	v := valgo.Is(valgo.BoolP(&bool1).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPEqualToWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valTrue := true
	v = valgo.Is(valgo.BoolP(&valTrue).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	valFalse := false
	v = valgo.Is(valgo.BoolP(&valFalse).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = valgo.Is(valgo.BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPEqualToWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valTrue := true

	v = valgo.Is(valgo.BoolP(&valTrue).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.ErrorByKey("value_0").Messages()[0])

	valFalse := false
	v = valgo.Is(valgo.BoolP(&valFalse).EqualTo(true))
	assert.Equal(t,
		"Value 0 must be equal to \"true\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = valgo.Is(valgo.BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPTrueWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valTrue := true

	v = valgo.Is(valgo.BoolP(&valTrue).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = valgo.Is(valgo.BoolP(&mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPTrueWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valFalse := false

	v = valgo.Is(valgo.BoolP(&valFalse).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = valgo.Is(valgo.BoolP(&mybool1).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPFalseWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	_valFalse := false
	valFalse := &_valFalse

	v = valgo.Is(valgo.BoolP(valFalse).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = false
	mybool1 := &_mybool1

	v = valgo.Is(valgo.BoolP(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPFalseWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	_valTrue := true
	valTrue := &_valTrue

	v = valgo.Is(valgo.BoolP(valTrue).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = true
	mybool1 := &_mybool1

	v = valgo.Is(valgo.BoolP(mybool1).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolNilIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	var valTrue *bool

	v = valgo.Is(valgo.BoolP(valTrue).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 *MyBool

	v = valgo.Is(valgo.BoolP(mybool1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolNilIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valTrue := true

	v = valgo.Is(valgo.BoolP(&valTrue).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = valgo.Is(valgo.BoolP(&mybool1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPPassingWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valTrue := true

	v = valgo.Is(valgo.BoolP(&valTrue).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = valgo.Is(valgo.BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPPassingWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valFalse := false

	v = valgo.Is(valgo.BoolP(&valFalse).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = valgo.Is(valgo.BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	boolValue := false

	v = valgo.Is(valgo.BoolP(&boolValue).InSlice([]bool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = false

	v = valgo.Is(valgo.BoolP(&myBool1).InSlice([]MyBool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	_boolValue := true
	boolValue := &_boolValue

	v = valgo.Is(valgo.BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	boolValue = nil

	v = valgo.Is(valgo.BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = true

	v = valgo.Is(valgo.BoolP(&myBool1).InSlice([]MyBool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}
