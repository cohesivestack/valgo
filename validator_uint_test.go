package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorUintConstructor(t *testing.T) {
	v := Is(Uint(uint(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint8(uint8(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint16(uint16(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint32(uint32(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint64(uint64(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Byte(byte(1)).Not().Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorUintNot(t *testing.T) {
	v := Is(Uint(uint(1)).Not().EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorUintEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).EqualTo(myUint2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(1)).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 1
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).EqualTo(myUint2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintGreaterThanValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(3)).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 3
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).GreaterThan(myUint2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorUintGreaterThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Uint(uint(2)).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).GreaterThan(myUint2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintGreaterOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint(uint(3)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).GreaterOrEqualTo(myUint2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintGreaterOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 3

	v = Is(Uint(MyUint(myUint1)).GreaterOrEqualTo(myUint2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintLessThanValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 3

	v = Is(Uint(MyUint(myUint1)).LessThan(myUint2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintLessThanInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Uint(uint(3)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).LessThan(myUint2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintLessOrEqualToValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint(uint(1)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).LessOrEqualTo(myUint2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintLessOrEqualToInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(3)).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 3
	var myUint2 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).LessOrEqualTo(myUint2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintBetweenValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint(uint(4)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Uint(uint(6)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 2
	var myUint3 MyUint = 6

	v = Is(Uint(MyUint(myUint1)).Between(myUint2, myUint3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintBetweenInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Uint(uint(7)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2
	var myUint2 MyUint = 3
	var myUint3 MyUint = 6

	v = Is(Uint(MyUint(myUint1)).Between(myUint2, myUint3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintZeroValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(0)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 0

	v = Is(Uint(uint(myUint1)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidatorUintZeroInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 1

	v = Is(Uint(uint(myUint1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintPassingValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(1)).Passing(func(val uint) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 1

	v = Is(Uint(MyUint(myUint1)).Passing(func(val MyUint) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorUintPassingInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(1)).Passing(func(val uint) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 1

	v = Is(Uint(MyUint(myUint1)).Passing(func(val MyUint) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintInSliceValid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(2)).InSlice([]uint{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 2

	v = Is(Uint(MyUint(myUint1)).InSlice([]MyUint{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorUintInSliceInvalid(t *testing.T) {
	var v *Validation

	v = Is(Uint(uint(4)).InSlice([]uint{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyUint uint
	var myUint1 MyUint = 4

	v = Is(Uint(MyUint(myUint1)).InSlice([]MyUint{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorUintOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Uint(uint(1)).Not().EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Uint(uint(1)).Not().EqualTo(uint(1)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)).EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)).EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorUintOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Uint(uint(1)).Not().EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Uint(uint(1)).Not().EqualTo(uint(1)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Uint(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)).EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Uint(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)).EqualTo(uint(0)))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(1)).Or().EqualTo(uint(0)))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Uint(uint(1)).EqualTo(uint(1)).EqualTo(uint(0)).Or().EqualTo(uint(1)))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
