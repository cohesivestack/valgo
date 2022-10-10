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
	ResetMessages()

	var v *ValidatorGroup

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
	ResetMessages()

	var v *ValidatorGroup

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
	ResetMessages()

	var v *ValidatorGroup

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
	ResetMessages()

	var v *ValidatorGroup

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
	ResetMessages()

	var v *ValidatorGroup

	valTen := 10

	v = Is(Any(valTen).Passing(func(val any) bool {
		return val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorAnyPassingInvalid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	valTen := 10

	v = Is(Any(valTen).Passing(func(val any) bool {
		return val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}
