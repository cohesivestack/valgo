package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorTypedNot(t *testing.T) {

	v := Is(Typed(10).Not().Passing(func(val int) bool {
		return val == 11
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedNilValid(t *testing.T) {

	var v *Validation

	var nilInterface interface{}
	v = Is(Typed(nilInterface).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Typed[interface{}](nil).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	var a *int
	v = Is(Typed(a).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	type X struct{}
	var x *X
	v = Is(Typed(x).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedNilInvalid(t *testing.T) {

	var v *Validation

	v = Is(Typed(0).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	type X struct{}
	x := X{}

	v = Is(Typed(&x).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTypedPassingValid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Typed(valTen).Passing(func(val int) bool {
		return val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	type Status string
	status := Status("running")

	// Test valid status
	v = Is(Typed(status).Passing(func(s Status) bool {
		return s == "running" || s == "paused"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedPassingInvalid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(Typed(valTen).Passing(func(val int) bool {
		return val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	type Status string
	status := Status("stopped")
	v = Is(Typed(status).Passing(func(s Status) bool {
		return s == "running" || s == "paused"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTypedOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Typed(true).Not().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Typed(true).Not().Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	// Testing Or operation with two valid conditions
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Typed(true).Not().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Typed(true).Not().Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Typed(true).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == true }).Or().Passing(func(val bool) bool { return val == false }))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Typed(true).Passing(func(val bool) bool { return val == true }).Passing(func(val bool) bool { return val == false }).Or().Passing(func(val bool) bool { return val == true }))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedOrElseOperatorWithIs(t *testing.T) {
	var v *Validation

	// Testing OrElse with left side valid - should short-circuit (key behavior)
	var input *string
	v = Is(Typed(input).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left side invalid - should continue to right side
	testStr := "test"
	v = Is(Typed(&testStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left invalid and right side fails
	otherStr := "other"
	v = Is(Typed(&otherStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with both sides invalid
	v = Is(Typed(&otherStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with Not() - left valid should short-circuit
	v = Is(Typed(&testStr).Not().Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "other" }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with Not() - left invalid should continue to right
	v = Is(Typed(input).Not().Nil().OrElse().Passing(func(s *string) bool { return s == nil }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTypedOrElseOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Testing OrElse with left side valid - should short-circuit
	var input *string
	v = Check(Typed(input).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left side invalid - should continue to right side
	testStr := "test"
	v = Check(Typed(&testStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left invalid and right side fails
	otherStr := "other"
	v = Check(Typed(&otherStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with both sides invalid
	v = Check(Typed(&otherStr).Nil().OrElse().Passing(func(s *string) bool { return s != nil && *s == "test" }))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
}
