package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorIntConstructor(t *testing.T) {
	v := Is(Int(1).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int8(int8(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int16(int16(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int32(int32(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int64(int64(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Rune(rune(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntNot(t *testing.T) {
	v := Is(Int(int(1)).Not().EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).EqualTo(myInt2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(1)).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 1
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).EqualTo(myInt2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntGreaterThanValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(3)).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 3
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).GreaterThan(myInt2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntGreaterThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Int(int(2)).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).GreaterThan(myInt2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int(int(3)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).GreaterOrEqualTo(myInt2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 3

	v = Is(Int(MyInt(myInt1)).GreaterOrEqualTo(myInt2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntLessThanValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 3

	v = Is(Int(MyInt(myInt1)).LessThan(myInt2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntLessThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Int(int(3)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).LessThan(myInt2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntLessOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int(int(1)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).LessOrEqualTo(myInt2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(3)).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 3
	var myInt2 MyInt = 2

	v = Is(Int(MyInt(myInt1)).LessOrEqualTo(myInt2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntBetweenValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int(int(4)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Int(int(6)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 2
	var myInt3 MyInt = 6

	v = Is(Int(MyInt(myInt1)).Between(myInt2, myInt3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntBetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Int(int(7)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2
	var myInt2 MyInt = 3
	var myInt3 MyInt = 6

	v = Is(Int(MyInt(myInt1)).Between(myInt2, myInt3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntZeroValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(0)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 0

	v = Is(Int(int(myInt1)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntZeroInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 1

	v = Is(Int(int(myInt1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPositiveValid(t *testing.T) {

	v := Is(Int(int(1)).Positive())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorIntPositiveInvalid(t *testing.T) {

	v := Is(Int(int(0)).Positive())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be positive",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Int(int(-1)).Positive())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be positive",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntNegativeValid(t *testing.T) {

	v := Is(Int(int(-1)).Negative())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntNegativeInvalid(t *testing.T) {

	v := Is(Int(int(0)).Negative())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be negative",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Int(int(1)).Negative())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be negative",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntPassingValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(1)).Passing(func(val int) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 1

	v = Is(Int(MyInt(myInt1)).Passing(func(val MyInt) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntPassingInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(1)).Passing(func(val int) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 1

	v = Is(Int(MyInt(myInt1)).Passing(func(val MyInt) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntInSliceValid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(2)).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 2

	v = Is(Int(MyInt(myInt1)).InSlice([]MyInt{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorIntInSliceInvalid(t *testing.T) {
	var v *Validation

	v = Is(Int(int(4)).InSlice([]int{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 4

	v = Is(Int(MyInt(myInt1)).InSlice([]MyInt{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorIntOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Int(int(1)).Not().EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Int(int(1)).Not().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorIntOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Int(int(1)).Not().EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Int(int(1)).Not().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
