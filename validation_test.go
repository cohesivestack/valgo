package valgo_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationMergeMessages(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v0 := valgo.Is(valgo.String("up", "status").EqualTo("down")).
		Is(valgo.String("", "name").Not().Blank()).
		AddErrorMessage("status", "The status is not valid").
		AddErrorMessage("base", "Record has errors")

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v0.Errors()["status"].Messages()[0])
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])

	v1 := valgo.Is(valgo.String("up", "status").Not().EqualTo("up")).
		Is(valgo.String("", "name").Not().Blank()).
		Is(valgo.Number(0, "position").Not().Zero()).
		// Same error message as in v0 should not be added twice
		AddErrorMessage("status", "The status is not valid").
		AddErrorMessage("status", "The status is incorrect")

	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Status can't be equal to \"up\"",
		v1.Errors()["status"].Messages()[0])
	assert.Equal(t,
		"The status is not valid",
		v1.Errors()["status"].Messages()[1])
	assert.Equal(t,
		"The status is incorrect",
		v1.Errors()["status"].Messages()[2])
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
	t.Parallel()

	v0 := valgo.Is(valgo.String("up", "status").EqualTo("up"))
	assert.True(t, v0.Valid())
	assert.Empty(t, v0.Errors())

	v1 := valgo.Is(valgo.String("", "name").Not().Blank())
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

func TestValidationIn(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.In("address",
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()))

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("address.line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("address.line2").Messages()[0])

	v.Is(valgo.String("", "line1").Not().Blank())

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("address.line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("address.line2").Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("line1").Messages()[0])
}

func TestValidationInDeeply(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.In("address",
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()).
			In("phone",
				valgo.Is(valgo.String("", "code").Not().Empty()).
					Is(valgo.String("", "number").Not().Empty())),
	)

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("address.line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("address.line2").Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.ErrorByKey("address.phone.code").Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.ErrorByKey("address.phone.number").Messages()[0])
}

func TestValidationInRow(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.InRow("addresses", 0,
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()),
	).InRow("addresses", 1,
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()))

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[0].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[0].line2").Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[1].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[1].line2").Messages()[0])

	v.Is(valgo.String("", "addresses").Not().Blank())

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[0].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[0].line2").Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[1].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[1].line2").Messages()[0])
	assert.Equal(t,
		"Addresses can't be blank",
		v.ErrorByKey("addresses").Messages()[0])
}

func TestValidationInRowDeeply(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.InRow("addresses", 0,
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()).
			InRow("phones", 0,
				valgo.Is(valgo.String("", "code").Not().Empty()).
					Is(valgo.String("", "number").Not().Empty())).
			InRow("phones", 1,
				valgo.Is(valgo.String("", "code").Not().Empty()).
					Is(valgo.String("", "number").Not().Empty())),
	).InRow("addresses", 1,
		valgo.Is(valgo.String("", "line1").Not().Blank()).
			Is(valgo.String("", "line2").Not().Blank()).
			InRow("phones", 0,
				valgo.Is(valgo.String("", "code").Not().Empty()).
					Is(valgo.String("", "number").Not().Empty())).
			InRow("phones", 1,
				valgo.Is(valgo.String("", "code").Not().Empty()).
					Is(valgo.String("", "number").Not().Empty())),
	)

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[0].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[0].line2").Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.ErrorByKey("addresses[0].phones[0].code").Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.ErrorByKey("addresses[0].phones[0].number").Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.ErrorByKey("addresses[0].phones[1].code").Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.ErrorByKey("addresses[0].phones[1].number").Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.ErrorByKey("addresses[1].line1").Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.ErrorByKey("addresses[1].line2").Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.ErrorByKey("addresses[1].phones[0].code").Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.ErrorByKey("addresses[1].phones[0].number").Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.ErrorByKey("addresses[1].phones[1].code").Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.ErrorByKey("addresses[1].phones[1].number").Messages()[0])
}

func TestLastValidationIsNotAlteringPreviousOne(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	// With validation
	v := valgo.Is(valgo.String("up", "status").EqualTo("down")).
		Is(valgo.String("Elon Musk", "name").Not().Blank())

	assert.False(t, v.Valid())

	assert.False(t, v.IsValid("status"))
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v.ErrorByKey("status").Messages()[0])

	assert.True(t, v.IsValid("name"))
	_, ok := v.Errors()["name"]
	assert.False(t, ok)

	// Adding error message
	v = valgo.AddErrorMessage("status", "Something is wrong").
		Is(valgo.String("Elon Musk", "name").Not().Blank())

	assert.False(t, v.Valid())

	assert.False(t, v.IsValid("status"))
	assert.Equal(t,
		"Something is wrong",
		v.ErrorByKey("status").Messages()[0])

	assert.True(t, v.IsValid("name"))
	_, ok = v.Errors()["name"]
	assert.False(t, ok)
}

func TestValidation_Valid(t *testing.T) {
	t.Parallel()

	val := valgo.Is(valgo.Number(21, "age").GreaterThan(18)).
		Is(valgo.String("singl", "status").InSlice([]string{"married", "single"}))

	require.False(t, val.Valid())

	out, err := json.Marshal(val.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"status":["Status is not valid"]}`, string(out))
}

func TestValidation_IsValid(t *testing.T) {
	t.Parallel()

	val := valgo.Is(valgo.Number(16, "age").GreaterThan(18)).
		Is(valgo.String("single", "status").InSlice([]string{"married", "single"}))

	require.False(t, val.IsValid("age"))
}

func TestValidation_Merge(t *testing.T) {
	t.Parallel()

	type Record struct {
		Name   string
		Status string
	}

	validatePreStatus := func(status string) *valgo.Validation {
		regex := regexp.MustCompile("pre-.+")

		return valgo.Check(valgo.String(status, "status").Not().Blank().MatchingTo(regex))
	}

	r := Record{"Classified", ""}

	val := valgo.Is(
		valgo.String(r.Name, "name").Not().Blank()).Is(
		valgo.String(r.Status, "status").Not().Blank())

	val.Merge(validatePreStatus(r.Status))

	require.False(t, val.Valid())

	out, err := json.Marshal(val.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"status":["Status can't be blank","Status must match to \"pre-.+\""]}`, string(out))
}
