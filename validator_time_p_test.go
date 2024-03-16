package valgo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidatorPTimeNot(t *testing.T) {
	time1 := time.Now()

	v := Is(TimeP(&time1).Not().EqualTo(time.Now().Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).EqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).EqualTo(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeAfterValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).After(now.Add(-1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeAfterInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).After(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(TimeP(&now).After(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeAfterOrEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).AfterOrEqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(TimeP(&now).AfterOrEqualTo(now.Add(-1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeAfterOrEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).AfterOrEqualTo(now.Add(1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be after or equal to \""+now.Add(1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeBeforeValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).Before(now.Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeBeforeInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).Before(now))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before \""+now.String()+"\"",
		v.Errors()["value_0"].Messages()[0])

	v = Is(TimeP(&now).Before(now.Add(-1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before \""+now.Add(-1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeBeforeOrEqualToValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).BeforeOrEqualTo(now))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(TimeP(&now).BeforeOrEqualTo(now.Add(1 * time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeBeforeOrEqualToInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).BeforeOrEqualTo(now.Add(-1 * time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be before or equal to \""+now.Add(-1*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeBetweenValid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).Between(now.Add(-1*time.Hour), now.Add(1*time.Hour)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeBetweenInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()

	v = Is(TimeP(&now).Between(now.Add(1*time.Hour), now.Add(2*time.Hour)))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \""+now.Add(1*time.Hour).String()+"\" and \""+now.Add(2*time.Hour).String()+"\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeInSliceValid(t *testing.T) {
	var v *Validation
	now := time.Now()
	timeSlice := []time.Time{now.Add(-1 * time.Hour), now, now.Add(1 * time.Hour)}

	v = Is(TimeP(&now).InSlice(timeSlice))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeInSliceInvalid(t *testing.T) {
	var v *Validation
	now := time.Now()
	timeSlice := []time.Time{now.Add(-1 * time.Hour), now.Add(-30 * time.Minute), now.Add(-15 * time.Minute)}

	v = Is(TimeP(&now).InSlice(timeSlice))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimePassingValid(t *testing.T) {

	var v *Validation

	now := time.Now()
	v = Is(TimeP(&now).Passing(func(val *time.Time) bool {
		return val.Equal(now)
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimePassingInvalid(t *testing.T) {

	var v *Validation

	now := time.Now()
	v = Is(TimeP(&now).Passing(func(val *time.Time) bool {
		return val.Equal(now.Add(1 * time.Hour))
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorPTimeZeroValid(t *testing.T) {

	var v *Validation

	zeroTime := time.Time{}

	v = Is(TimeP(&zeroTime).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorPTimeZeroInvalid(t *testing.T) {

	var v *Validation

	nonZeroTime := time.Now()

	v = Is(TimeP(&nonZeroTime).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}
