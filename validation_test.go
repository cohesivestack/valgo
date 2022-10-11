package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationMergeMessages(t *testing.T) {

	v0 := Is(String("up", "status").EqualTo("down")).
		Is(String("", "name").Not().Blank()).
		AddErrorMessage("status", "The status is not valid").
		AddErrorMessage("base", "Record has errors")

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v0.Errors()["status"].Messages()[0])
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])

	v1 := Is(String("up", "status").Not().EqualTo("up")).
		Is(String("", "name").Not().Blank()).
		Is(Number(0, "position").Not().Zero()).
		AddErrorMessage("status", "The status is incorrect")

	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Status can't be equal to \"up\"",
		v1.Errors()["status"].Messages()[0])
	assert.Equal(t,
		"The status is incorrect",
		v1.Errors()["status"].Messages()[1])
	assert.Equal(t,
		"Name can't be blank",
		v1.Errors()["name"].Messages()[0])
	assert.Equal(t,
		"Position must not be zero",
		v1.Errors()["position"].Messages()[0])

	v0.Merge(v1)

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v0.Errors()["status"].Messages()[0])
	assert.Equal(t,
		"The status is not valid",
		v0.Errors()["status"].Messages()[1])
	assert.Equal(t,
		"Status can't be equal to \"up\"",
		v0.Errors()["status"].Messages()[2])
	assert.Equal(t,
		"The status is incorrect",
		v0.Errors()["status"].Messages()[3])
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])
	assert.Equal(t,
		"Record has errors",
		v0.Errors()["base"].Messages()[0])
	assert.Equal(t,
		"Position must not be zero",
		v0.Errors()["position"].Messages()[0])

}

func TestValidationMergeInvalidate(t *testing.T) {

	v0 := Is(String("up", "status").EqualTo("up"))
	assert.True(t, v0.Valid())
	assert.Empty(t, v0.Errors())

	v1 := Is(String("", "name").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v1.Errors()["name"].Messages()[0])

	// v0 is initially valid, but merging to v1 must be invalidated
	v0.Merge(v1)

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])
}
