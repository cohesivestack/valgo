package valgo

import (
	"encoding/json"
	"fmt"
	"regexp"
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

func TestValidationIn(t *testing.T) {

	v := In("address",
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()))

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["address.line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["address.line2"].Messages()[0])

	v.Is(String("", "line1").Not().Blank())

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["address.line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["address.line2"].Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["line1"].Messages()[0])
}

func TestValidationInDeeply(t *testing.T) {

	v := In("address",
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()).
			In("phone",
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())),
	)

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["address.line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["address.line2"].Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.Errors()["address.phone.code"].Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.Errors()["address.phone.number"].Messages()[0])
}

func TestValidationInRow(t *testing.T) {

	v := InRow("addresses", 0,
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()),
	).InRow("addresses", 1,
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()))

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[0].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[0].line2"].Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[1].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[1].line2"].Messages()[0])

	v.Is(String("", "addresses").Not().Blank())

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[0].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[0].line2"].Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[1].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[1].line2"].Messages()[0])
	assert.Equal(t,
		"Addresses can't be blank",
		v.Errors()["addresses"].Messages()[0])
}

func TestValidationInRowDeeply(t *testing.T) {

	v := InRow("addresses", 0,
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()).
			InRow("phones", 0,
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())).
			InRow("phones", 1,
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())),
	).InRow("addresses", 1,
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()).
			InRow("phones", 0,
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())).
			InRow("phones", 1,
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())),
	)

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[0].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[0].line2"].Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.Errors()["addresses[0].phones[0].code"].Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.Errors()["addresses[0].phones[0].number"].Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.Errors()["addresses[0].phones[1].code"].Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.Errors()["addresses[0].phones[1].number"].Messages()[0])
	assert.Equal(t,
		"Line 1 can't be blank",
		v.Errors()["addresses[1].line1"].Messages()[0])
	assert.Equal(t,
		"Line 2 can't be blank",
		v.Errors()["addresses[1].line2"].Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.Errors()["addresses[1].phones[0].code"].Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.Errors()["addresses[1].phones[0].number"].Messages()[0])
	assert.Equal(t,
		"Code can't be empty",
		v.Errors()["addresses[1].phones[1].code"].Messages()[0])
	assert.Equal(t,
		"Number can't be empty",
		v.Errors()["addresses[1].phones[1].number"].Messages()[0])
}

func TestLastValidationIsNotAlteringPreviousOne(t *testing.T) {

	// With validation
	v := Is(String("up", "status").EqualTo("down")).
		Is(String("Elon Musk", "name").Not().Blank())

	assert.False(t, v.Valid())

	assert.False(t, v.IsValid("status"))
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v.Errors()["status"].Messages()[0])

	assert.True(t, v.IsValid("name"))
	_, ok := v.Errors()["name"]
	assert.False(t, ok)

	// Adding error message
	v = AddErrorMessage("status", "Something is wrong").
		Is(String("Elon Musk", "name").Not().Blank())

	assert.False(t, v.Valid())

	assert.False(t, v.IsValid("status"))
	assert.Equal(t,
		"Something is wrong",
		v.Errors()["status"].Messages()[0])

	assert.True(t, v.IsValid("name"))
	_, ok = v.Errors()["name"]
	assert.False(t, ok)
}

func ExampleValidation_Valid() {
	val := Is(Number(21, "age").GreaterThan(18)).
		Is(String("singl", "status").InSlice([]string{"married", "single"}))

	if !val.Valid() {
		out, _ := json.MarshalIndent(val.Error(), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "status": [
	//     "Status is not valid"
	//   ]
	// }
}

func ExampleValidation_IsValid() {
	val := Is(Number(16, "age").GreaterThan(18)).
		Is(String("single", "status").InSlice([]string{"married", "single"}))

	if !val.IsValid("age") {
		fmt.Println("Warning: someone underage is trying to sign up")
	}

	// output: Warning: someone underage is trying to sign up
}

func ExampleValidation_Merge() {
	type Record struct {
		Name   string
		Status string
	}

	validatePreStatus := func(status string) *Validation {
		regex, _ := regexp.Compile("pre-.+")

		return Check(String(status, "status").Not().Blank().MatchingTo(regex))
	}

	r := Record{"Classified", ""}

	val := Is(
		String(r.Name, "name").Not().Blank()).Is(
		String(r.Status, "status").Not().Blank())

	val.Merge(validatePreStatus(r.Status))

	if !val.Valid() {
		out, _ := json.MarshalIndent(val.Error(), "", "  ")
		fmt.Println(string(out))
	}

	// output: {
	//   "status": [
	//     "Status can't be blank",
	//     "Status must match to \"pre-.+\""
	//   ]
	// }
}
