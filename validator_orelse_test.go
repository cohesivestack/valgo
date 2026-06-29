package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorOrElseOperatorWithIs(t *testing.T) {
	var v *Validation

	// -------------------------------------------------------------------------
	// Basic OrElse behavior (acts like Or, but cuts the chain on success)
	// -------------------------------------------------------------------------

	// True OrElse True (cuts on first True)
	v = Is(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True OrElse False (cuts on first True; right side is not evaluated)
	v = Is(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False OrElse True (evaluates right side)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False OrElse False (fails with OR-joined message)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False OrElse False (with Not) - fails with OR-joined message (redundant subject preserved)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().Not().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 can't be equal to \"1\"",
		v.Errors()["value_0"].Messages()[0])

	// -------------------------------------------------------------------------
	// OrElse cuts the entire remainder of the chain (including any AND tail)
	// -------------------------------------------------------------------------

	// True OrElse False . False  (cuts; AND tail not evaluated)
	v = Is(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)).EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True OrElse (False Or False) (cuts; OR chain on the right not evaluated)
	v = Is(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)).Or().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True (Or True) OrElse False (cuts based on left OR-chain)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).OrElse().EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True (Or True) OrElse False . False (cuts; RHS + AND tail not evaluated)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).OrElse().EqualTo(int(2)).EqualTo(int(3)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// -------------------------------------------------------------------------
	// OrElse does NOT cut when left is false; RHS is evaluated normally
	// -------------------------------------------------------------------------

	// False OrElse True . True  (RHS evaluated; AND tail evaluated)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(1)).EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False OrElse True . False (OR-group passes via RHS; AND tail fails => only AND message)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(1)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False OrElse (False Or True) . False (RHS OR-group passes; AND tail fails => only AND message)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).Or().EqualTo(int(1)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// -------------------------------------------------------------------------
	// OrElse inside an OR chain: message joining must match Or() behavior
	// -------------------------------------------------------------------------

	// False OrElse False Or False (fails with single OR-joined message)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False OrElse False (same join, different placement of OrElse)
	v = Is(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).OrElse().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// False OrElse False Or True (passes)
	v = Is(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// True OrElse (False Or True) (cuts; RHS OR-chain not evaluated)
	v = Is(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)).Or().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// -------------------------------------------------------------------------
	// Document current precedence rule: Or/OrElse cannot "rescue" a failure that
	// happened before the Or/OrElse is introduced (same behavior as Or tests)
	// -------------------------------------------------------------------------

	// False . True OrElse True (still fails because first condition fails before OrElse boundary)
	v = Is(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).OrElse().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorOrElseOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check is non-short-circuited by default, but OrElse must still cut the chain on success.

	// -------------------------------------------------------------------------
	// Basic OrElse behavior
	// -------------------------------------------------------------------------

	// True OrElse False (cuts; RHS not evaluated)
	v = Check(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False OrElse True (evaluates RHS)
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(1)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// False OrElse False (fails with OR-joined message)
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False OrElse False (with Not)
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().Not().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 can't be equal to \"1\"",
		v.Errors()["value_0"].Messages()[0])

	// -------------------------------------------------------------------------
	// OrElse cuts the entire remainder of the chain even under Check()
	// -------------------------------------------------------------------------

	// True OrElse False . False (cuts; AND tail not evaluated; no errors)
	v = Check(Int(int(1)).EqualTo(int(1)).OrElse().EqualTo(int(0)).EqualTo(int(2)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// (False Or True) OrElse False . False (cuts based on left OR-chain; RHS + tail not evaluated)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(1)).OrElse().EqualTo(int(2)).EqualTo(int(3)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// -------------------------------------------------------------------------
	// When OrElse does not cut (left is false), Check continues collecting errors
	// -------------------------------------------------------------------------

	// False OrElse True . False => OR passes via RHS, AND tail fails => only AND message
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(1)).EqualTo(int(2)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// False OrElse False . False => OR group fails, AND tail still evaluated under Check => 2 messages
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\" or Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[1])

	// False OrElse False Or False . False => OR group fails, AND tail evaluated => 2 messages
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).Or().EqualTo(int(3)).EqualTo(int(4)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 2, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
	assert.Equal(t,
		"Value 0 must be equal to \"4\"",
		v.Errors()["value_0"].Messages()[1])

	// -------------------------------------------------------------------------
	// OrElse inside an OR chain: message joining must match Or() behavior
	// -------------------------------------------------------------------------

	// False OrElse False Or False (single OR-joined message)
	v = Check(Int(int(1)).EqualTo(int(0)).OrElse().EqualTo(int(2)).Or().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// False Or False OrElse False (same join, different placement)
	v = Check(Int(int(1)).EqualTo(int(0)).Or().EqualTo(int(2)).OrElse().EqualTo(int(3)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"; Value 0 must be equal to \"2\"; or Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// -------------------------------------------------------------------------
	// Document current precedence rule: OrElse cannot rescue failures that occur
	// before the OrElse boundary is introduced (same behavior as Or tests)
	// -------------------------------------------------------------------------

	// False . True OrElse True (still fails because first condition fails before OrElse boundary)
	v = Check(Int(int(1)).EqualTo(int(0)).EqualTo(int(1)).OrElse().EqualTo(int(1)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t, 1, len(v.Errors()["value_0"].Messages()))
	assert.Equal(t,
		"Value 0 must be equal to \"0\"",
		v.Errors()["value_0"].Messages()[0])
}
