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
