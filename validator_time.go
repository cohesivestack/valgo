package valgo

import (
	"time"
)

func isTimeEqualTo(v0 time.Time, v1 time.Time) bool {
	return v0.Equal(v1)
}

func isTimeAfter(v0 time.Time, v1 time.Time) bool {
	return v0.After(v1)
}

func isTimeAfterOrEqualTo(v0 time.Time, v1 time.Time) bool {
	return v0.After(v1) || v0.Equal(v1)
}

func isTimeBefore(v0 time.Time, v1 time.Time) bool {
	return v0.Before(v1)
}

func isTimeBeforeOrEqualTo(v0 time.Time, v1 time.Time) bool {
	return v0.Before(v1) || v0.Equal(v1)
}

func isTimeZero(v time.Time) bool {
	return v.IsZero()
}

func isTimeBetween(v time.Time, min time.Time, max time.Time) bool {
	return (v.After(min) || v.Equal(min)) && (v.Before(max) || v.Equal(max))
}

func isTimeInSlice(v time.Time, slice []time.Time) bool {
	for _, _v := range slice {
		if v.Equal(_v) {
			return true
		}
	}
	return false
}

// The `ValidatorTime` structure provides a set of methods to perform validation
// checks on time.Time values, utilizing Go's native time package.
type ValidatorTime struct {
	context *ValidatorContext
}

// The Time function initiates a new `ValidatorTime` instance to validate a given
// time value. The optional name and title parameters can be used for enhanced
// error reporting. If a name is provided without a title, the name is humanized
// to be used as the title.
//
// For example:
//
//	startTime := time.Now()
//	v := ValidatorTime{}
//	v.Time(startTime, "start_time", "Start Time")
func Time(value time.Time, nameAndTitle ...string) *ValidatorTime {
	return &ValidatorTime{context: NewContext(value, nameAndTitle...)}
}

// The Context method returns the current context of the validator, which can
// be utilized to create custom validations by extending this validator.
func (validator *ValidatorTime) Context() *ValidatorContext {
	return validator.context
}

// The Not method inverts the boolean value associated with the next validator
// method. This can be used to negate the check performed by the next validation
// method in the chain.
//
// For example:
//
//	// Will return false because Not() inverts the boolean value of the Zero() function
//	startTime := time.Now()
//	Is(v.Time(startTime).Not().Zero()).Valid()
func (validator *ValidatorTime) Not() *ValidatorTime {
	validator.context.Not()
	return validator
}

// The EqualTo method validates if the time value is equal to another given time
// value. It uses the equality (`==`) operator from Go for the comparison.
//
// For example:
//
//	timeA := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	timeB := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	Is(v.Time(timeA).EqualTo(timeB)).Valid()
func (validator *ValidatorTime) EqualTo(value time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeEqualTo(validator.context.Value().(time.Time), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// The After method checks if the time value is after a specified time.
//
// For example:
//
//	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	endTime := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
//	Is(v.Time(endTime).After(startTime)).Valid()
func (validator *ValidatorTime) After(value time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeAfter(validator.context.Value().(time.Time), value)
		},
		ErrorKeyAfter, value, template...)

	return validator
}

// The AfterOrEqualTo method checks if the time value is either after or equal to
// a specified time.
//
// For example:
//
//	timeA := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	timeB := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	Is(v.Time(timeA).AfterOrEqualTo(timeB)).Valid()
func (validator *ValidatorTime) AfterOrEqualTo(value time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeAfterOrEqualTo(validator.context.Value().(time.Time), value)
		},
		ErrorKeyAfterOrEqualTo, value, template...)

	return validator
}

// The Before method checks if the time value is before a specified time.
//
// For example:
//
//	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	endTime := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
//	Is(v.Time(startTime).Before(endTime)).Valid()
func (validator *ValidatorTime) Before(value time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeBefore(validator.context.Value().(time.Time), value)
		},
		ErrorKeyBefore, value, template...)

	return validator
}

// The BeforeOrEqualTo method checks if the time value is either before or equal to
// a specified time.
//
// For example:
//
//	timeA := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	timeB := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	Is(v.Time(timeA).BeforeOrEqualTo(timeB)).Valid()
func (validator *ValidatorTime) BeforeOrEqualTo(value time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeBeforeOrEqualTo(validator.context.Value().(time.Time), value)
		},
		ErrorKeyBeforeOrEqualTo, value, template...)

	return validator
}

// The Between method verifies if the time value falls within a given time range, inclusive.
//
// For example:
//
//	minTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	maxTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
//	checkTime := time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC)
//	Is(v.Time(checkTime).Between(minTime, maxTime)).Valid()
func (validator *ValidatorTime) Between(min time.Time, max time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithParams(
		func() bool {
			return isTimeBetween(validator.context.Value().(time.Time), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// The Zero method verifies if the time value is a zero time, which means it hasn't
// been initialized yet.
//
// For example:
//
//	zeroTime := time.Time{}
//	Is(v.Time(zeroTime).Zero()).Valid()
func (validator *ValidatorTime) Zero(template ...string) *ValidatorTime {
	validator.context.Add(
		func() bool {
			return isTimeZero(validator.context.Value().(time.Time))
		},
		ErrorKeyZero, template...)

	return validator
}

// The Passing method allows for custom validation logic by accepting a function
// that returns a boolean indicating whether the validation passed or failed.
//
// For example:
//
//	checkTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	Is(v.Time(checkTime).Passing(func(t time.Time) bool {
//	    return t.Year() == 2023
//	})).Valid()
func (validator *ValidatorTime) Passing(function func(v0 time.Time) bool, template ...string) *ValidatorTime {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(time.Time))
		},
		ErrorKeyPassing, template...)

	return validator
}

// The InSlice method validates if the time value is found within a provided slice
// of time values.
//
// For example:
//
//	timeA := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	timeB := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
//	timeSlice := []time.Time{timeA, timeB}
//	checkTime := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
//	Is(v.Time(checkTime).InSlice(timeSlice)).Valid()
func (validator *ValidatorTime) InSlice(slice []time.Time, template ...string) *ValidatorTime {
	validator.context.AddWithValue(
		func() bool {
			return isTimeInSlice(validator.context.Value().(time.Time), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
