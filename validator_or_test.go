package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorOrOperatorWithIs(t *testing.T) {
	var v *Validation

	// True Or True
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False
	v = Is(Int(int(1)).EqualTo(int(0)).Or().Not().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 can't be equal to \"1\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or True Or True
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or True
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False Or True
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or False
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False Or False
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or False Or False
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or False (duplicate of previous, but testing again)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False Or True
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False . True ((True Or False) And True)
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True . True ((False Or True) And True)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False . True ((False Or False) And True)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False . False ((False Or False) And False)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or False . False ((True Or False) And False)
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or True . False ((True Or True) And False)
	v = Is(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or True . False ((False Or True) And False)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False . False (duplicate)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// True . True Or False ((True And True) Or False)
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True . False Or True ((True And False) Or True)
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True . False Or False ((True And False) Or False)
	v = Is(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or False ((False And False) Or False)
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . True Or False ((False And True) Or False)
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . True Or True ((False And True) Or True) - fails because first condition fails
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or True ((False And False) Or True) - fails because first condition fails
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or False (duplicate)
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	// True Or True
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or True Or True
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or True
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False Or True
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or False
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False Or False
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or False Or False
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True Or False
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False Or True
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True Or False . True ((True Or False) And True)
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or True . True ((False Or True) And True)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False Or False . True ((False Or False) And True)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False . False ((False Or False) And False)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[1])

	// True Or False . False ((True Or False) And False)
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// True Or True . False ((True Or True) And False)
	v = Check(Int(int(1)).EqualTo(int(1)).Or().EqualTo(int(1)).EqualTo(int(0)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or True . False ((False Or True) And False)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False . False (duplicate)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[1])

	// True . True Or False ((True And True) Or False)
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(1)).Or().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True . False Or True ((True And False) Or True)
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True . False Or False ((True And False) Or False)
	v = Check(Int(int(1)).EqualTo(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or False ((False And False) Or False)
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"2\" or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[1])

	// False . True Or False ((False And True) Or False)
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).Or().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . True Or True ((False And True) Or True) - fails because first condition fails
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).Or().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or True ((False And False) Or True) - fails because first condition fails
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])

	// False . False Or False (duplicate)
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"2\" or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[1])
}
