package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorBoolNot(t *testing.T) {
	t.Parallel()

	v := valgo.Is(valgo.Bool(true).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.Bool(false).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = valgo.Is(valgo.Bool(mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = false

	v = valgo.Is(valgo.Bool(mybool1).EqualTo(mybool2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolTrueWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = valgo.Is(valgo.Bool(mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolTrueWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(false).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = valgo.Is(valgo.Bool(mybool1).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolFalseWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(false).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = valgo.Is(valgo.Bool(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolFalseWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = valgo.Is(valgo.Bool(mybool1).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPassingWhenIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).Passing(func(val bool) bool {
		return val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = valgo.Is(valgo.Bool(mybool1).Passing(func(val MyBool) bool {
		return val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPassingWhenIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(false).Passing(func(val bool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = valgo.Is(valgo.Bool(mybool1).Passing(func(val MyBool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(false).InSlice([]bool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = false

	v = valgo.Is(valgo.Bool(myBool1).InSlice([]MyBool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Bool(true).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = true

	v = valgo.Is(valgo.Bool(myBool1).InSlice([]MyBool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}
