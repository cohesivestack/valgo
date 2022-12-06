package valgo_test

import (
	"regexp"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorStringNot(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.Is(valgo.String("text1").Not().EqualTo("text2"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("text").EqualTo("text"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"
	var myString2 MyString = "text"

	v = valgo.Is(valgo.String(myString1).EqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("text1").EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"
	var myString2 MyString = "text2"

	v = valgo.Is(valgo.String(myString1).EqualTo(myString2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringGreaterThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("ab").GreaterThan("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).GreaterThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringGreaterThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").GreaterThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String("aa").GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).GreaterThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringGreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("ab").GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).GreaterOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringGreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = valgo.Is(valgo.String(myString1).GreaterOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringLessThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").LessThan("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = valgo.Is(valgo.String(myString1).LessThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringLessThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String("ab").LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).LessThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringLessOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").LessOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("aa").LessOrEqualTo("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).LessOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringLessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("ab").LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = valgo.Is(valgo.String(myString1).LessOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringBetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("ac").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("ae").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"
	var myString3 MyString = "ae"

	v = valgo.Is(valgo.String(myString1).Between(myString2, myString3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringBetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("aa").Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String("af").Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"
	var myString3 MyString = "ae"

	v = valgo.Is(valgo.String(myString1).Between(myString2, myString3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringEmptyValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("").Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString

	v = valgo.Is(valgo.String(myString1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringEmptyInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("a").Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String(" ").Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = valgo.Is(valgo.String(myString1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringBlankValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("").Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String(" ").Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v = valgo.Is(valgo.String(myString1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringBlankInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.String("a").Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = valgo.Is(valgo.String(myString1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPassingValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("text").Passing(func(val string) bool {
		return val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v = valgo.Is(valgo.String(myString1).Passing(func(val MyString) bool {
		return val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPassingInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("text1").Passing(func(val string) bool {
		return val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"

	v = valgo.Is(valgo.String(myString1).Passing(func(val MyString) bool {
		return val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("up").InSlice([]string{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = valgo.Is(valgo.String(myString1).InSlice([]MyString{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("up").InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = valgo.Is(valgo.String(myString1).InSlice([]MyString{"down", "indle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringMatchingToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	regex := regexp.MustCompile("pre-.+")

	v = valgo.Is(valgo.String("pre-approved").MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "pre-approved"

	v = valgo.Is(valgo.String(myString1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringMatchingToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	regex := regexp.MustCompile("pre-.+")

	v = valgo.Is(valgo.String("pre_approved").MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "pre_approved"

	v = valgo.Is(valgo.String(myString1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringMaxLengthValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("123456").MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("12345").MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = valgo.Is(valgo.String(myString1).MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringMaxLengthInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("1234567").MaxLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "1234567"

	v = valgo.Is(valgo.String(myString1).MaxLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringMinLengthValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("123456").MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("1234567").MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = valgo.Is(valgo.String(myString1).MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringMinLengthInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("12345").MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = valgo.Is(valgo.String(myString1).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringLengthValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("123456").OfLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = valgo.Is(valgo.String(myString1).OfLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringLengthInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("12345").OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String("1234567").OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = valgo.Is(valgo.String(myString1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringLengthBetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("123456").OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("12345678").OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.String("1234567890").OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = valgo.Is(valgo.String(myString1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringLengthBetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.String("12345").OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	v = valgo.Is(valgo.String("12345678901").OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = valgo.Is(valgo.String(myString1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])
}
