package valgo

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorStringNot(t *testing.T) {

	v := Is(String("text1").Not().EqualTo("text2"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(String("text").EqualTo("text"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"
	var myString2 MyString = "text"

	v = Is(String(myString1).EqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("text1").EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"
	var myString2 MyString = "text2"

	v = Is(String(myString1).EqualTo(myString2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringGreaterThanValid(t *testing.T) {
	var v *Validation

	v = Is(String("ab").GreaterThan("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = Is(String(myString1).GreaterThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringGreaterThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").GreaterThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("aa").GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(String(myString1).GreaterThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("ab").GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(String(myString1).GreaterOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = Is(String(myString1).GreaterOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringLessThanValid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").LessThan("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = Is(String(myString1).LessThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringLessThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("ab").LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(String(myString1).LessThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringLessOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").LessOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("aa").LessOrEqualTo("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(String(myString1).LessOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("ab").LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = Is(String(myString1).LessOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringBetweenValid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("ac").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("ae").Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"
	var myString3 MyString = "ae"

	v = Is(String(myString1).Between(myString2, myString3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringBetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("aa").Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("af").Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"
	var myString3 MyString = "ae"

	v = Is(String(myString1).Between(myString2, myString3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringEmptyValid(t *testing.T) {
	var v *Validation

	v = Is(String("").Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = ""

	v = Is(String(myString1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringEmptyInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("a").Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String(" ").Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(String(myString1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringBlankValid(t *testing.T) {
	var v *Validation

	v = Is(String("").Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String(" ").Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v = Is(String(myString1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringBlankInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("a").Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(String(myString1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPassingValid(t *testing.T) {

	var v *Validation

	v = Is(String("text").Passing(func(val string) bool {
		return val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v = Is(String(myString1).Passing(func(val MyString) bool {
		return val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPassingInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("text1").Passing(func(val string) bool {
		return val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"

	v = Is(String(myString1).Passing(func(val MyString) bool {
		return val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringInSliceValid(t *testing.T) {

	var v *Validation

	v = Is(String("up").InSlice([]string{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = Is(String(myString1).InSlice([]MyString{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringInSliceInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("up").InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = Is(String(myString1).InSlice([]MyString{"down", "indle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringMatchingToValid(t *testing.T) {

	var v *Validation

	regex, _ := regexp.Compile("pre-.+")

	v = Is(String("pre-approved").MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "pre-approved"

	v = Is(String(myString1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringMatchingToInvalid(t *testing.T) {

	var v *Validation

	regex, _ := regexp.Compile("pre-.+")

	v = Is(String("pre_approved").MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "pre_approved"

	v = Is(String(myString1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringMaxBytesValid(t *testing.T) {

	var v *Validation

	v = Is(String("123456").MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("12345").MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(String(myString1).MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringMaxBytesInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("1234567").MaxBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "1234567"

	v = Is(String(myString1).MaxBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringMinBytesValid(t *testing.T) {

	var v *Validation

	v = Is(String("123456").MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("1234567").MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(String(myString1).MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringMinBytesInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("12345").MinBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(String(myString1).MinBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringOfByteLengthValid(t *testing.T) {
	var v *Validation

	v = Is(String("123456").OfByteLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(String(myString1).OfByteLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringOfByteLengthInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("12345").OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("1234567").OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(String(myString1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringOfByteLengthBetweenValid(t *testing.T) {

	var v *Validation

	v = Is(String("123456").OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("12345678").OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("1234567890").OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(String(myString1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringOfByteLengthBetweenInvalid(t *testing.T) {

	var v *Validation

	v = Is(String("12345").OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("12345678901").OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(String(myString1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringMaxLengthValid(t *testing.T) {
	var v *Validation

	// "虎視眈々" has 4 runes, 12 bytes
	v = Is(String("虎視眈々").MaxLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("虎視眈々").MaxLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringMaxLengthInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").MaxLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringMinLengthValid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").MinLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("虎視眈々").MinLength(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringMinLengthInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").MinLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"5\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringOfLengthValid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").OfLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringOfLengthInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").OfLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("虎視眈々").OfLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"5\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringOfLengthBetweenValid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").OfLengthBetween(4, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(String("虎視眈々").OfLengthBetween(2, 4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringOfLengthBetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(String("虎視眈々").OfLengthBetween(5, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"5\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(String("虎視眈々").OfLengthBetween(1, 3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"1\" and \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(String("1").EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(String("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(String("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(String("1").EqualTo("0").Or().EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(String("1").EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(String("1").Not().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(String("1").Not().EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(String("1").EqualTo("1").Or().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(String("1").EqualTo("0").Or().EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(String("1").EqualTo("0").Or().EqualTo("1").EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(String("1").EqualTo("0").Or().EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(String("1").EqualTo("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(String("1").EqualTo("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorStringOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(String("1").EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(String("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(String("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(String("1").EqualTo("0").Or().EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(String("1").EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(String("1").Not().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(String("1").Not().EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(String("1").EqualTo("1").Or().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(String("1").EqualTo("0").Or().EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(String("1").EqualTo("0").Or().EqualTo("1").EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(String("1").EqualTo("0").Or().EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(String("1").EqualTo("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(String("1").EqualTo("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
