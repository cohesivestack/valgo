package valgo

import (
	"time"
)

// ValidatorTimeP is a type that facilitates validation for time pointer variables.
// It retains a context that records details about the validation process.
type ValidatorTimeP struct {
	context *ValidatorContext
}

// TimeP initializes a new ValidatorTimeP instance with the provided time pointer
// and optional name and title arguments for detailed error messages.
//
// Usage example:
//
//	var myTime *time.Time
//	v.TimeP(myTime, "start_time", "Start Time")
func TimeP(value *time.Time, nameAndTitle ...string) *ValidatorTimeP {
	return &ValidatorTimeP{context: NewContext(value, nameAndTitle...)}
}

// Context retrieves the context associated with the validator.
func (validator *ValidatorTimeP) Context() *ValidatorContext {
	return validator.context
}

// Not negates the result of the next validator function in the chain.
//
// Usage example:
//
//	t := time.Now()
//	Is(v.TimeP(&t).Not().Zero()).Valid()  // Will return false since t is not a zero time.
func (validator *ValidatorTimeP) Not() *ValidatorTimeP {
	validator.context.Not()
	return validator
}

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the time is before or equal to time.Now().
//	t := time.Now()
//	isValid := v.Is(v.TimeP(&t).Nil().Or().BeforeOrEqualTo(time.Now())).Valid()
func (validator *ValidatorTimeP) Or() *ValidatorTimeP {
	validator.context.Or()
	return validator
}

// EqualTo validates that the time pointer is equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).EqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) EqualTo(value time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeEqualTo(*(validator.context.Value().(*time.Time)), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// After validates that the time pointer is after the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1.Add(-time.Hour)
//	Is(v.TimeP(&t1).After(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) After(value time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeAfter(*(validator.context.Value().(*time.Time)), value)
		},
		ErrorKeyAfter, value, template...)

	return validator
}

// AfterOrEqualTo validates that the time pointer is after or equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).AfterOrEqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) AfterOrEqualTo(value time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeAfterOrEqualTo(*(validator.context.Value().(*time.Time)), value)
		},
		ErrorKeyAfterOrEqualTo, value, template...)

	return validator
}

// Before validates that the time pointer is before the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1.Add(time.Hour)
//	Is(v.TimeP(&t1).Before(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) Before(value time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeBefore(*(validator.context.Value().(*time.Time)), value)
		},
		ErrorKeyBefore, value, template...)

	return validator
}

// BeforeOrEqualTo validates that the time pointer is before or equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).BeforeOrEqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) BeforeOrEqualTo(value time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeBeforeOrEqualTo(*(validator.context.Value().(*time.Time)), value)
		},
		ErrorKeyBeforeOrEqualTo, value, template...)

	return validator
}

// Between validates that the time pointer is between the specified minimum and maximum time values (inclusive).
//
// Usage example:
//
//	t1 := time.Now()
//	min := t1.Add(-time.Hour)
//	max := t1.Add(time.Hour)
//	Is(v.TimeP(&t1).Between(min, max)).Valid()  // Will return true.
func (validator *ValidatorTimeP) Between(min time.Time, max time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithParams(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeBetween(*(validator.context.Value().(*time.Time)), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Zero validates that the time pointer is pointing to a zero time value.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).Zero()).Valid()  // Will return true as t is nil and thus pointing to a zero time.
func (validator *ValidatorTimeP) Zero(template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeZero(*(validator.context.Value().(*time.Time)))
		},
		ErrorKeyZero, validator.context.Value(), template...)

	return validator
}

// Passing allows for custom validation function to be applied on the time pointer.
//
// Usage example:
//
//	t := time.Now()
//	Is(v.TimeP(&t).Passing(func(v0 *time.Time) bool { return v0.After(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)) })).Valid()  // Custom validation.
func (validator *ValidatorTimeP) Passing(function func(v0 *time.Time) bool, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return function(validator.context.Value().(*time.Time))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// InSlice validates that the time pointer is pointing to a time value present in the specified slice.
//
// Usage example:
//
//	t := time.Now()
//	validTimes := []time.Time{t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)}
//	Is(v.TimeP(&t).InSlice(validTimes)).Valid()  // Will return true.
func (validator *ValidatorTimeP) InSlice(slice []time.Time, template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) != nil && isTimeInSlice(*(validator.context.Value().(*time.Time)), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Nil validates that the time pointer is nil.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).Nil()).Valid()  // Will return true as t is nil.
func (validator *ValidatorTimeP) Nil(template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) == nil
		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}

// NilOrZero validates that the time pointer is either nil or pointing to a zero time value.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).NilOrZero()).Valid()  // Will return true as t is nil.
func (validator *ValidatorTimeP) NilOrZero(template ...string) *ValidatorTimeP {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value().(*time.Time) == nil || isTimeZero(*(validator.context.Value().(*time.Time)))

		},
		ErrorKeyNil, validator.context.Value(), template...)

	return validator
}
