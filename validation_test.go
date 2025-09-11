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
				Is(
					String("", "code").Not().Empty(),
					String("", "number").Not().Empty())), // Passing multiple validators to Is
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
		Is(String("", "line1").Not().Blank(),
			String("", "line2").Not().Blank())) // Passing multiple validators to Is

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
				Is(String("", "code").Not().Empty(),
					String("", "number").Not().Empty())), // Passing multiple validators to Is
	).InRow("addresses", 1,
		Is(String("", "line1").Not().Blank()).
			Is(String("", "line2").Not().Blank()).
			InRow("phones", 0,
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())).
			InRow("phones", 1,
				Is(String("", "code").Not().Empty(),
					String("", "number").Not().Empty())), // Passing multiple validators to Is
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

func TestValidationInCell(t *testing.T) {

	v := InCell("names", 0,
		Is(String("", "name").Not().Blank()),
	).InCell("names", 1,
		Is(String("", "name").Not().Blank()))

	assert.False(t, v.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v.Errors()["names[0]"].Messages()[0])
	assert.Equal(t,
		"Name can't be blank",
		v.Errors()["names[1]"].Messages()[0])
}

func TestValidationInCellDeeply(t *testing.T) {

	v := InCell("addresses", 0,
		Is(String("", "name").Not().Blank()).
			Is(String("", "street").Not().Blank()).
			In("phone",
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())),
	).InCell("addresses", 1,
		Is(String("", "name").Not().Blank()).
			Is(String("", "street").Not().Blank()).
			In("phone",
				Is(String("", "code").Not().Empty()).
					Is(String("", "number").Not().Empty())),
	)

	assert.False(t, v.Valid())

	assert.Len(t, v.Errors()["addresses[0]"].Messages(), 4)
	assert.Len(t, v.Errors()["addresses[1]"].Messages(), 4)

	assert.Contains(t,
		v.Errors()["addresses[0]"].Messages(),
		"Name can't be blank")
	assert.Contains(t,
		v.Errors()["addresses[0]"].Messages(),
		"Street can't be blank")
	assert.Contains(t,
		v.Errors()["addresses[0]"].Messages(),
		"Code can't be empty")
	assert.Contains(t,
		v.Errors()["addresses[0]"].Messages(),
		"Number can't be empty")
	assert.Contains(t,
		v.Errors()["addresses[1]"].Messages(),
		"Name can't be blank")
	assert.Contains(t,
		v.Errors()["addresses[1]"].Messages(),
		"Street can't be blank")
	assert.Contains(t,
		v.Errors()["addresses[1]"].Messages(),
		"Code can't be empty")
	assert.Contains(t,
		v.Errors()["addresses[1]"].Messages(),
		"Number can't be empty")
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

func TestValidationMergeError(t *testing.T) {

	v0 := Is(String("up", "status").EqualTo("up"))
	assert.True(t, v0.Valid())
	assert.Empty(t, v0.Errors())

	v1 := Is(String("", "name").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v1.Errors()["name"].Messages()[0])

	// v0 is initially valid, but merging to v1 must be invalidated
	v0.MergeError(v1.ToValgoError())

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])

	assert.True(t, v0.IsValid("status"))

	// Check that mergin error is not replacing existing errors

	v0 = Is(String("up", "status").EqualTo("down"))
	assert.False(t, v0.IsValid("status"))
	assert.Len(t, v0.Errors(), 1)

	v0.Is(String("", "name").Not().Blank())
	assert.False(t, v0.IsValid("name"))
	assert.Len(t, v0.Errors(), 2)

	v1 = Is(String("", "status").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Status can't be blank",
		v1.Errors()["status"].Messages()[0])

	// v0 is initially valid, but merging to v1 must be invalidated
	v0.MergeError(v1.ToValgoError())

	assert.False(t, v0.IsValid("status"))
	assert.Equal(t,
		"Status must be equal to \"down\"",
		v0.Errors()["status"].Messages()[0])

	assert.False(t, v0.IsValid("status"))
	assert.Equal(t,
		"Status can't be blank",
		v0.Errors()["status"].Messages()[1])

	assert.False(t, v0.IsValid("name"))
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["name"].Messages()[0])
}

func TestValidationMergeErrorIn(t *testing.T) {

	v0 := Is(String("up", "status").EqualTo("up"))
	assert.True(t, v0.Valid())
	assert.Empty(t, v0.Errors())

	v1 := Is(String("", "firstName").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"First name can't be blank",
		v1.Errors()["firstName"].Messages()[0])

	v1.Is(String("", "lastName").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Last name can't be blank",
		v1.Errors()["lastName"].Messages()[0])

	// v0 is initially valid, but merging to v1 Errors must be invalidated
	v0.MergeErrorIn("user", v1.ToValgoError())

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"First name can't be blank",
		v0.Errors()["user.firstName"].Messages()[0])

	assert.Equal(t,
		"Last name can't be blank",
		v0.Errors()["user.lastName"].Messages()[0])

	assert.True(t, v0.IsValid("status"))

}

func TestValidationMergeErrorInRow(t *testing.T) {

	v0 := Is(String("up", "status").EqualTo("up"))
	assert.True(t, v0.Valid())
	assert.Empty(t, v0.Errors())

	v1 := Is(String("", "name").Not().Blank())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v1.Errors()["name"].Messages()[0])

	// v0 is initially valid, but merging to v1 Errors must be invalidated
	v0.MergeErrorInRow("user", 0, v1.ToValgoError())

	// v0 is initially valid, but merging to v1 Errors must be invalidated
	v0.MergeErrorInRow("user", 1, v1.ToValgoError())

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["user[0].name"].Messages()[0])

	assert.Equal(t,
		"Name can't be blank",
		v0.Errors()["user[1].name"].Messages()[0])

	assert.True(t, v0.IsValid("status"))

}

func TestValidationCustomTitle(t *testing.T) {
	v0 := Is(String("", "company_name").Not().Empty())

	assert.False(t, v0.Valid())
	assert.Equal(t,
		"Company name can't be empty",
		v0.Errors()["company_name"].Messages()[0])

	v1 := Is(String("", "company_name", "Customer").Not().Empty())
	assert.False(t, v1.Valid())
	assert.Equal(t,
		"Customer can't be empty",
		v1.Errors()["company_name"].Messages()[0])
}

func TestValidationCustomTitlePanic(t *testing.T) {
	v0 := Is(String("", "company_name").Not().Empty())
	assert.False(t, v0.Valid())

	if !v0.Valid() {
		// calling valErr.Title() should not panic even if there is no
		// custom title given
		for _, valErr := range v0.Errors() {
			assert.NotPanics(t, func() {
				valErr.Title()
			})
		}
	}
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

func TestIsFunctionWithMultipleValidators(t *testing.T) {
	v1 := String("testValue", "field1").EqualTo("testValue")
	v2 := Number(100, "field2").GreaterThan(50)

	v := Is(v1, v2)
	assert.True(t, v.Valid(), "Expected multiple validators passed to Is function to be valid")

	vInvalid := Is(
		String("testValue", "field1").EqualTo("testValue"), // This should be valid
		Number(10, "field2").GreaterThan(50),               // This should be invalid
	)
	assert.False(t, vInvalid.Valid())
	assert.Contains(t, vInvalid.Errors(), "field2", "Error map should contain field2")
}

func TestCheckFunctionWithMultipleValidators(t *testing.T) {
	v1 := String("testValue", "field1").Not().EqualTo("wrongValue")
	v2 := Number(42, "field2").LessThan(100)

	validation := Check(v1, v2)
	assert.True(t, validation.Valid())

	vInvalid1 := String("wrongValue", "field1").EqualTo("rightValue") // This should be invalid
	vInvalid2 := Number(150, "field2").LessThan(100)                  // This should be invalid

	validationMixed := Check(vInvalid1, v2) // Mix of valid and invalid
	assert.False(t, validationMixed.Valid())

	validationAllInvalid := Check(vInvalid1, vInvalid2) // All invalid
	assert.False(t, validationAllInvalid.Valid())
	assert.Contains(t, validationAllInvalid.Errors(), "field1", "Error map should contain field1")
	assert.Contains(t, validationAllInvalid.Errors(), "field2", "Error map should contain field2")
}

func TestValidationIf(t *testing.T) {
	// Test If with true condition - should execute function
	v := Is(String("test", "name").Not().Blank()).
		Is(String("", "phone").Not().Blank()).
		Is(String("", "email").EqualTo("test@test.com")).
		If(true, Is(String("", "email").Not().Blank()))
	assert.False(t, v.Valid())
	assert.Contains(t, v.Errors(), "email")
	assert.Contains(t, v.Errors(), "phone")
	assert.Equal(t, "Phone can't be blank", v.Errors()["phone"].Messages()[0])
	assert.Equal(t, "Email must be equal to \"test@test.com\"", v.Errors()["email"].Messages()[0])
	assert.Equal(t, "Email can't be blank", v.Errors()["email"].Messages()[1])
}

func TestValidationIfWithFalseCondition(t *testing.T) {
	// Test If with false condition - should not execute function
	v := Is(String("test", "name").Not().Blank()).
		If(false, Is(String("", "email").Not().Blank()))
	assert.True(t, v.Valid())
	assert.NotContains(t, v.Errors(), "email")
}

func TestValidationWhen(t *testing.T) {
	// Test When with true condition - should execute function
	v := Is(String("test", "name").Not().Blank()).
		When(true, func(val *Validation) {
			val.Is(String("", "email").Not().Blank())
		})
	assert.False(t, v.Valid())
	assert.Contains(t, v.Errors(), "email")
	assert.Equal(t, "Email can't be blank", v.Errors()["email"].Messages()[0])

	// Test When with false condition - should not execute function
	v2 := Is(String("test", "name").Not().Blank()).
		When(false, func(val *Validation) {
			val.Is(String("", "email").Not().Blank())
		})
	assert.True(t, v2.Valid())
	assert.NotContains(t, v2.Errors(), "email")
}
