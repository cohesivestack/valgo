package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorComparablePNot(t *testing.T) {

	val := 10

	v := Is(ComparableP(&val).Not().EqualTo(11))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePEqualToValid(t *testing.T) {

	var v *Validation

	val := 10

	v = Is(ComparableP(&val).EqualTo(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with string
	text := "hello"
	v = Is(ComparableP(&text).EqualTo("hello"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with struct (comparable)
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	v = Is(ComparableP(&p1).EqualTo(Point{X: 1, Y: 2}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePEqualToInvalid(t *testing.T) {

	var v *Validation

	ival := 11
	v = Is(ComparableP(&ival).EqualTo(10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"10\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with string
	text := "hello"
	v = Is(ComparableP(&text).EqualTo("world"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"world\"",
		v.Errors()["value_0"].Messages()[0])

	// Test with struct (comparable)
	type Point struct {
		X, Y int
	}
	p1 := Point{X: 1, Y: 2}
	v = Is(ComparableP(&p1).EqualTo(Point{X: 1, Y: 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"{1 3}\"",
		v.Errors()["value_0"].Messages()[0])

	// Nil pointer
	var pn *int
	v = Is(ComparableP(pn).EqualTo(11))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"11\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePNilValid(t *testing.T) {

	var v *Validation

	var a *int
	v = Is(ComparableP(a).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	type X struct{}
	var x *X
	v = Is(ComparableP(x).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePNilInvalid(t *testing.T) {

	var v *Validation

	val := 0
	v = Is(ComparableP(&val).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	type X struct{}
	x := X{}

	v = Is(ComparableP(&x).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePPassingValid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(ComparableP(&valTen).Passing(func(val *int) bool {
		return val != nil && *val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePPassingInvalid(t *testing.T) {

	var v *Validation

	valTen := 10

	v = Is(ComparableP(&valTen).Passing(func(val *int) bool {
		return val != nil && *val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePPassingWithCustomType(t *testing.T) {

	var v *Validation

	type Status string
	status := Status("running")

	// Test valid status
	v = Is(ComparableP(&status).Passing(func(s *Status) bool {
		return s != nil && (*s == "running" || *s == "paused")
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test invalid status
	status = Status("stopped")
	v = Is(ComparableP(&status).Passing(func(s *Status) bool {
		return s != nil && (*s == "running" || *s == "paused")
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePInSliceValid(t *testing.T) {

	var v *Validation

	val := 10
	v = Is(ComparableP(&val).InSlice([]int{5, 10, 15}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Test with string
	text := "hello"
	v = Is(ComparableP(&text).InSlice([]string{"world", "hello", "test"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyInt int
	var myInt1 MyInt = 10

	v = Is(ComparableP(&myInt1).InSlice([]MyInt{5, 10, 15}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePInSliceInvalid(t *testing.T) {

	var v *Validation

	val := 10
	v = Is(ComparableP(&val).InSlice([]int{5, 15, 20}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Test with string
	text := "hello"
	v = Is(ComparableP(&text).InSlice([]string{"world", "test", "example"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Nil pointer
	var pn *int
	v = Is(ComparableP(pn).InSlice([]int{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorComparablePOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	one := 1

	// Testing Or operation with two valid conditions
	v = Is(ComparableP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(ComparableP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(ComparableP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(ComparableP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())
}

func TestValidatorComparablePOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	one := 1

	// Testing Or operation with two valid conditions
	v = Check(ComparableP(&one).EqualTo(1).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(ComparableP(&one).EqualTo(0).Or().EqualTo(1))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(ComparableP(&one).EqualTo(1).Or().EqualTo(0))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(ComparableP(&one).EqualTo(0).Or().EqualTo(0))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())
}

func TestValidatorComparablePOrElseOperatorWithIs(t *testing.T) {
	var v *Validation

	pending := "pending"
	active := "active"
	unknown := "unknown"
	validStatuses := []string{"active", "inactive"}

	// Testing OrElse with left side valid - should short-circuit (key behavior)
	v = Is(ComparableP(&pending).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left side invalid - should continue to right side
	v = Is(ComparableP(&active).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left invalid and right side fails
	v = Is(ComparableP(&unknown).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with both sides invalid
	v = Is(ComparableP(&unknown).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with Not() - left valid should short-circuit
	v = Is(ComparableP(&active).Not().EqualTo("pending").OrElse().InSlice([]string{"pending"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with Not() - left invalid should continue to right
	v = Is(ComparableP(&pending).Not().EqualTo("pending").OrElse().InSlice([]string{"pending"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorComparablePOrElseOperatorWithCheck(t *testing.T) {
	var v *Validation

	pending := "pending"
	active := "active"
	unknown := "unknown"
	validStatuses := []string{"active", "inactive"}

	// Testing OrElse with left side valid - should short-circuit
	v = Check(ComparableP(&pending).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left side invalid - should continue to right side
	v = Check(ComparableP(&active).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing OrElse with left invalid and right side fails
	v = Check(ComparableP(&unknown).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing OrElse with both sides invalid
	v = Check(ComparableP(&unknown).EqualTo("pending").OrElse().InSlice(validStatuses))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
}
