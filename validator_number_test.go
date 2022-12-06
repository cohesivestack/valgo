package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorNumberNot(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.Is(valgo.Number(1).Not().EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(1).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberGreaterThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(3).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberGreaterThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.Number(2).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberGreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.Number(3).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberGreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.Number(myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberLessThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.Number(myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberLessThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.Number(3).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberLessOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.Number(1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberLessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(3).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberBetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.Number(4).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.Number(6).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.Number(myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberBetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.Number(7).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.Number(myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberZeroValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(0).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber

	v = valgo.Is(valgo.Number(myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberZeroInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.Number(1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.Number(0.1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.Number(myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPassingValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Number(1).Passing(func(val int) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.Number(myNumber1).Passing(func(val MyNumber) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPassingInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Number(1).Passing(func(val int) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.Number(myNumber1).Passing(func(val MyNumber) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Number(2).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = valgo.Is(valgo.Number(myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.Number(4).InSlice([]int{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 4

	v = valgo.Is(valgo.Number(myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}
