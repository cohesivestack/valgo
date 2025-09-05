package valgo

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorStringPNot(t *testing.T) {

	text1 := "text1"

	v := Is(StringP(&text1).Not().EqualTo("text2"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPEqualToValid(t *testing.T) {
	var v *Validation

	text1 := "text"

	v = Is(StringP(&text1).EqualTo("text"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"
	var myString2 MyString = "text"

	v = Is(StringP(&myString1).EqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPEqualToInvalid(t *testing.T) {
	var v *Validation

	_text1 := "text1"
	text1 := &_text1

	v = Is(StringP(text1).EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"
	var myString2 MyString = "text2"

	v = Is(StringP(&myString1).EqualTo(myString2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPGreaterThanValid(t *testing.T) {
	var v *Validation

	text1 := "ab"

	v = Is(StringP(&text1).GreaterThan("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).GreaterThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPGreaterThanInvalid(t *testing.T) {
	var v *Validation

	_text1 := "aa"
	text1 := &_text1

	v = Is(StringP(text1).GreaterThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	*text1 = "aa"

	v = Is(StringP(text1).GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).GreaterThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	text1 := "aa"

	v = Is(StringP(&text1).GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "ab"

	v = Is(StringP(&text1).GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).GreaterOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_text1 := "aa"
	text1 := &_text1

	v = Is(StringP(text1).GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = Is(StringP(&myString1).GreaterOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPLessThanValid(t *testing.T) {
	var v *Validation

	text1 := "aa"

	v = Is(StringP(&text1).LessThan("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v = Is(StringP(&myString1).LessThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPLessThanInvalid(t *testing.T) {
	var v *Validation

	_text1 := "aa"
	text1 := &_text1

	v = Is(StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	*text1 = "ab"

	v = Is(StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).LessThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPLessOrEqualToValid(t *testing.T) {
	var v *Validation

	text1 := "aa"

	v = Is(StringP(&text1).LessOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&text1).LessOrEqualTo("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).LessOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	_text1 := "ab"
	text1 := &_text1

	v = Is(StringP(text1).LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil
	v = Is(StringP(text1).LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v = Is(StringP(&myString1).LessOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPBetweenValid(t *testing.T) {
	var v *Validation

	text1 := "aa"

	v = Is(StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "ac"

	v = Is(StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "ae"

	v = Is(StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"
	var myString3 MyString = "ae"

	v = Is(StringP(&myString1).Between(myString2, myString3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPBetweenInvalid(t *testing.T) {
	var v *Validation

	_text1 := "aa"
	text1 := &_text1

	v = Is(StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])

	*text1 = "af"

	v = Is(StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"
	var myString3 MyString = "ae"

	v = Is(StringP(&myString1).Between(myString2, myString3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPEmptyValid(t *testing.T) {
	var v *Validation

	_text1 := ""
	text1 := &_text1

	v = Is(StringP(text1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = ""

	v = Is(StringP(&myString1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPEmptyInvalid(t *testing.T) {
	var v *Validation

	_text1 := "a"
	text1 := &_text1

	v = Is(StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	*text1 = " "

	v = Is(StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(StringP(&myString1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPEmptyOrNilValid(t *testing.T) {
	var v *Validation

	_text1 := ""
	text1 := &_text1

	v = Is(StringP(text1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = nil

	v = Is(StringP(text1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = ""

	v = Is(StringP(&myString1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPEmptyOrNilInvalid(t *testing.T) {
	var v *Validation

	text1 := "a"

	v = Is(StringP(&text1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	text1 = " "

	v = Is(StringP(&text1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(StringP(&myString1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPBlankValid(t *testing.T) {
	var v *Validation

	_text1 := ""
	text1 := &_text1

	v = Is(StringP(text1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	*text1 = " "

	v = Is(StringP(text1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v = Is(StringP(&myString1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPBlankInvalid(t *testing.T) {
	var v *Validation

	_text1 := "a"
	text1 := &_text1

	v = Is(StringP(text1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(StringP(&myString1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPBlankOrNilValid(t *testing.T) {
	var v *Validation

	_text1 := ""
	text1 := &_text1

	v = Is(StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	*text1 = " "

	v = Is(StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = nil

	v = Is(StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v = Is(StringP(&myString1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPBlankOrNilInvalid(t *testing.T) {
	var v *Validation

	text1 := "a"

	v = Is(StringP(&text1).BlankOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v = Is(StringP(&myString1).BlankOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPPassingValid(t *testing.T) {

	var v *Validation

	text1 := "text"

	v = Is(StringP(&text1).Passing(func(val *string) bool {
		return *val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v = Is(StringP(&myString1).Passing(func(val *MyString) bool {
		return *val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPPassingInvalid(t *testing.T) {

	var v *Validation

	text1 := "text1"

	v = Is(StringP(&text1).Passing(func(val *string) bool {
		return *val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"

	v = Is(StringP(&myString1).Passing(func(val *MyString) bool {
		return *val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPInSliceValid(t *testing.T) {

	var v *Validation

	text1 := "up"

	v = Is(StringP(&text1).InSlice([]string{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = Is(StringP(&myString1).InSlice([]MyString{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPInSliceInvalid(t *testing.T) {

	var v *Validation

	_text1 := "up"
	text1 := &_text1

	v = Is(StringP(text1).InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v = Is(StringP(&myString1).InSlice([]MyString{"down", "indle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPMatchingToValid(t *testing.T) {

	var v *Validation

	regex, _ := regexp.Compile("pre-.+")

	text1 := "pre-approved"

	v = Is(StringP(&text1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "pre-approved"

	v = Is(StringP(&myString1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPMatchingToInvalid(t *testing.T) {

	var v *Validation

	regex, _ := regexp.Compile("pre-.+")

	_text1 := "pre_approved"
	text1 := &_text1

	v = Is(StringP(text1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "pre_approved"

	v = Is(StringP(&myString1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPMaxBytesValid(t *testing.T) {

	var v *Validation

	text1 := "123456"

	v = Is(StringP(&text1).MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "12345"

	v = Is(StringP(&text1).MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(StringP(&myString1).MaxBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPMaxBytesInvalid(t *testing.T) {

	var v *Validation

	_text1 := "1234567"
	text1 := &_text1

	v = Is(StringP(text1).MaxBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).MaxBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "1234567"

	v = Is(StringP(&myString1).MaxBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPMinBytesValid(t *testing.T) {

	var v *Validation

	text1 := "123456"

	v = Is(StringP(&text1).MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "1234567"

	v = Is(StringP(&text1).MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(StringP(&myString1).MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	myString1 = "1234567"

	v = Is(StringP(&myString1).MinBytes(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPMinBytesInvalid(t *testing.T) {

	var v *Validation

	_text1 := "12345"
	text1 := &_text1

	v = Is(StringP(text1).MinBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).MinBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(StringP(&myString1).MinBytes(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPOfByteLengthValid(t *testing.T) {

	var v *Validation

	text1 := "123456"

	v = Is(StringP(&text1).OfByteLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(StringP(&myString1).OfByteLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPOfByteLengthInvalid(t *testing.T) {

	var v *Validation

	_text1 := "12345"
	text1 := &_text1

	v = Is(StringP(text1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*text1 = "1234567"

	v = Is(StringP(text1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(StringP(&myString1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	myString1 = "1234567"

	v = Is(StringP(&myString1).OfByteLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPOfByteLengthBetweenValid(t *testing.T) {

	var v *Validation

	text1 := "123456"

	v = Is(StringP(&text1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "12345678"

	v = Is(StringP(&text1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	text1 = "1234567890"

	v = Is(StringP(&text1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v = Is(StringP(&myString1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	myString1 = "12345678"
	v = Is(StringP(&myString1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	myString1 = "1234567890"
	v = Is(StringP(&myString1).OfByteLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorStringPOfByteLengthBetweenInvalid(t *testing.T) {

	var v *Validation

	_text1 := "12345"
	text1 := &_text1

	v = Is(StringP(text1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	*text1 = "12345678901"

	v = Is(StringP(text1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	text1 = nil

	v = Is(StringP(text1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v = Is(StringP(&myString1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	myString1 = "12345678901"
	v = Is(StringP(&myString1).OfByteLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPMaxLengthValid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).MaxLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&text).MaxLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).MaxLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textLatin).MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).MaxLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textAccent).MaxLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).MaxLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPMaxLengthInvalid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).MaxLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	textP := (*string)(nil)
	v = Is(StringP(textP).MaxLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).MaxLength(4))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"4\"",
		v.Errors()["value_0"].Messages()[0])

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).MaxLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).MaxLength(4))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"4\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPMinLengthValid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).MinLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&text).MinLength(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).MinLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textLatin).MinLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).MinLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textAccent).MinLength(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).MinLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPMinLengthInvalid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).MinLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"5\"",
		v.Errors()["value_0"].Messages()[0])

	textP := (*string)(nil)
	v = Is(StringP(textP).MinLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"5\"",
		v.Errors()["value_0"].Messages()[0])

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).MinLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"5\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPOfLengthValid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).OfLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).OfLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).OfLength(4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).OfLength(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPOfLengthInvalid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).OfLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&text).OfLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"5\"",
		v.Errors()["value_0"].Messages()[0])

	textP := (*string)(nil)
	v = Is(StringP(textP).OfLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"5\"",
		v.Errors()["value_0"].Messages()[0])

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).OfLength(4))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"4\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&textLatin).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).OfLength(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&textAccent).OfLength(5))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"5\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).OfLength(4))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"4\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPOfLengthBetweenValid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).OfLengthBetween(4, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&text).OfLengthBetween(2, 4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).OfLengthBetween(4, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textLatin).OfLengthBetween(5, 5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).OfLengthBetween(3, 5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(StringP(&textAccent).OfLengthBetween(4, 4))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).OfLengthBetween(4, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringPOfLengthBetweenInvalid(t *testing.T) {
	var v *Validation

	text := "虎視眈々"
	v = Is(StringP(&text).OfLengthBetween(5, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"5\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&text).OfLengthBetween(1, 3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"1\" and \"3\"",
		v.Errors()["value_0"].Messages()[0])

	textP := (*string)(nil)
	v = Is(StringP(textP).OfLengthBetween(1, 3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"1\" and \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Latin characters - "Hello" has 5 runes
	textLatin := "Hello"
	v = Is(StringP(&textLatin).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&textLatin).OfLengthBetween(1, 4))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"1\" and \"4\"",
		v.Errors()["value_0"].Messages()[0])

	// Mixed Latin characters with accents - "Café" has 4 runes
	textAccent := "Café"
	v = Is(StringP(&textAccent).OfLengthBetween(5, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"5\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(StringP(&textAccent).OfLengthBetween(1, 3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"1\" and \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "Hello"

	v = Is(StringP(&myString1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringNilIsValid(t *testing.T) {

	var v *Validation

	var valString *string

	v = Is(StringP(valString).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyString string
	var myString1 *MyString

	v = Is(StringP(myString1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorStringNilIsInvalid(t *testing.T) {

	var v *Validation

	valString := "text"

	v = Is(StringP(&valString).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v = Is(StringP(&myString1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorStringPOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	var one = "1"

	// Testing Or operation with two valid conditions
	v = Is(StringP(&one).EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(StringP(&one).EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(StringP(&one).EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(StringP(&one).EqualTo("0").Or().EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(StringP(&one).EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(StringP(&one).Not().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(StringP(&one).Not().EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(StringP(&one).EqualTo("1").Or().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(StringP(&one).EqualTo("0").Or().EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(StringP(&one).EqualTo("0").Or().EqualTo("1").EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(StringP(&one).EqualTo("0").Or().EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(StringP(&one).EqualTo("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(StringP(&one).EqualTo("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorStringPOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	var one = "1"

	// Testing Or operation with two valid conditions
	v = Check(StringP(&one).EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(StringP(&one).EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(StringP(&one).EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(StringP(&one).EqualTo("0").Or().EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(StringP(&one).EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(StringP(&one).Not().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(StringP(&one).Not().EqualTo("1").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(StringP(&one).EqualTo("1").Or().EqualTo("0").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(StringP(&one).EqualTo("0").Or().EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(StringP(&one).EqualTo("0").Or().EqualTo("1").EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(StringP(&one).EqualTo("0").Or().EqualTo("1").EqualTo("0"))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(StringP(&one).EqualTo("1").EqualTo("1").Or().EqualTo("0"))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(StringP(&one).EqualTo("1").EqualTo("0").Or().EqualTo("1"))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
