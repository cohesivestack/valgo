package valgo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidatorTimeNot(t *testing.T) {

	v := Is(Time(time.Now()).Not().EqualTo(time.Now().Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).EqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).EqualTo(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeAfterValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now.Add(1 * time.Hour)).After(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeAfterInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).After(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Time(now).After(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeAfterOrEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).AfterOrEqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Time(now.Add(1 * time.Hour)).AfterOrEqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeAfterOrEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).AfterOrEqualTo(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after or equal to \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeBeforeValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).Before(now.Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeBeforeInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).Before(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(Time(now.Add(1 * time.Hour)).Before(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeBeforeOrEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).BeforeOrEqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Time(now).BeforeOrEqualTo(now.Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeBeforeOrEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now.Add(1 * time.Hour)).BeforeOrEqualTo(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before or equal to \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeBetweenValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).Between(now.Add(-1*time.Hour), now.Add(1*time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeBetweenInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(Time(now).Between(now.Add(1*time.Hour), now.Add(2*time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \""+now.Add(1*time.Hour).String()+"\" and \""+now.Add(2*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeInSliceValid(t *testing.T) {
	var v *Validation
	now := time.Now()
	timeSlice := []time.Time{now.Add(-1 * time.Hour), now, now.Add(1 * time.Hour)}

	v = Is(Time(now).InSlice(timeSlice))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeInSliceInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()
	timeSlice := []time.Time{now.Add(-1 * time.Hour), now.Add(-30 * time.Minute), now.Add(-15 * time.Minute)}

	v = Is(Time(now).InSlice(timeSlice))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimePassingValid(t *testing.T) {

	var v *Validation

	now := time.Now()
	v = Is(Time(now).Passing(func(val time.Time) bool {
		return val.Equal(now)
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimePassingInvalid(t *testing.T) {

	var v *Validation

	now := time.Now()
	v = Is(Time(now).Passing(func(val time.Time) bool {
		return val.Equal(now.Add(1 * time.Hour))
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeZeroValid(t *testing.T) {

	var v *Validation

	zeroTime := time.Time{}

	v = Is(Time(zeroTime).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorTimeZeroInvalid(t *testing.T) {

	var v *Validation

	nonZeroTime := time.Now()

	v = Is(Time(nonZeroTime).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorTimeOrOperatorWithIs(t *testing.T) {
	var v *Validation

	var _true = true
	var _false = false

	timeZero := time.Time{}
	timeOne := time.Time{}.Add(time.Second)

	// Testing Or operation with two valid conditions
	v = Is(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Is(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Is(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Is(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Is(Time(timeOne).EqualTo(timeOne).EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Is(Time(timeOne).Not().EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Is(Time(timeOne).Not().EqualTo(timeOne).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Is(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Is(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Is(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne).EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Is(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne).EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Is(Time(timeOne).EqualTo(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Is(Time(timeOne).EqualTo(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())

}

func TestValidatorTimeOrOperatorWithCheck(t *testing.T) {
	var v *Validation

	// Check are Non-Short-circuited operations

	var _true = true
	var _false = false

	timeZero := time.Time{}
	timeOne := time.Time{}.Add(time.Second)

	// Testing Or operation with two valid conditions
	v = Check(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, _true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left invalid and right valid conditions
	v = Check(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with left valid and right invalid conditions
	v = Check(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing Or operation with two invalid conditions
	v = Check(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, _false || false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing And operation (default when no Or() function is used) with left valid and right invalid conditions
	v = Check(Time(timeOne).EqualTo(timeOne).EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing combination of Not and Or operators with left valid and right invalid conditions
	v = Check(Time(timeOne).Not().EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, !false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing combination of Not and Or operators with left invalid and right valid conditions
	v = Check(Time(timeOne).Not().EqualTo(timeOne).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, !true || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the first condition being valid
	v = Check(Time(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, true || _false || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing multiple Or operations in sequence with the last condition being valid
	v = Check(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, _false || false || true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid Or operation then valid And operation
	v = Check(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne).EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, false || _true && true, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing valid Or operation then invalid And operation
	v = Check(Time(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne).EqualTo(timeZero))
	assert.False(t, v.Valid())
	assert.Equal(t, false || true && false, v.Valid())
	assert.NotEmpty(t, v.Errors())

	// Testing valid And operation then invalid Or operation
	v = Check(Time(timeOne).EqualTo(timeOne).EqualTo(timeOne).Or().EqualTo(timeZero))
	assert.True(t, v.Valid())
	assert.Equal(t, _true && true || false, v.Valid())
	assert.Empty(t, v.Errors())

	// Testing invalid And operation then valid Or operation
	v = Check(Time(timeOne).EqualTo(timeOne).EqualTo(timeZero).Or().EqualTo(timeOne))
	assert.True(t, v.Valid())
	assert.Equal(t, true && false || true, v.Valid())
	assert.Empty(t, v.Errors())
}
