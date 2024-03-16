package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorAnyNot(t *testing.T) {

	v := Is(Any(10).Not().EqualTo(11))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorAnyEqualToValid(t *testing.T) {

	var v *Validation

	v = Is(Any(10).EqualTo(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	type X struct{ Value int }
	x := X{Value: 10}
	y := X{Value: 10}
	v = Is(Any(x).EqualTo(y))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	var a *int
	var b *int

	v = Is(Any(a).EqualTo(b))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorAnyEqualToInvalid(t *testing.T) {

	var v *Validation

	v = Is(Any(11).EqualTo(10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"10\"",
		v.Errors()["value_0"].Messages()[0])

	type X struct{ Value int }
	x := X{Value: 10}
	y := X{Value: 11}
	v = Is(Any(x).EqualTo(y))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"{11}\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Any(10).EqualTo(nil))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"<nil>\"",
		v.Errors()["value_0"].Messages()[0])

	// Both nil but different types
	var a *int
	var b *int64

	v = Is(Any(a).EqualTo(b))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"<nil>\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorAnyNilValid(t *testing.T) {

	var v *Validation

	var a *int
	v = Is(Any(a).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	type X struct{}
	var x *X
	v = Is(Any(x).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorAnyNilInvalid(t *testing.T) {

	var v *Validation

	v = Is(Any(0).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	type X struct{}
	x := X{}

	v = Is(Any(&x).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorAnyPassingValid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Any(valTen).Passing(func(val any) bool {
		return val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorAnyPassingInvalid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Any(valTen).Passing(func(val any) bool {
		return val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorAnyOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Any(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Any(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Any(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Any(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Any(true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Any(true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Any(true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Any(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Any(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Any(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Any(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Any(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Any(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorAnyOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Any(true).EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Any(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Any(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Any(true).EqualTo(false).Or().EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Any(true).EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Any(true).Not().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Any(true).Not().EqualTo(true).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Any(true).EqualTo(true).Or().EqualTo(false).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Any(true).EqualTo(false).Or().EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Any(true).EqualTo(false).Or().EqualTo(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Any(true).EqualTo(false).Or().EqualTo(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Any(true).EqualTo(true).EqualTo(true).Or().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Any(true).EqualTo(true).EqualTo(false).Or().EqualTo(true))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
