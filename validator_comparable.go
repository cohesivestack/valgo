package valgo

// The Comparable validator's type that keeps its validator context.
// T can be any Go type (pointer, struct, etc.) that is comparable.
type ValidatorComparable[T comparable] struct {
	context *ValidatorContext
}

// Receive a value of type T to validate, where T must be comparable.
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.
//
// Example:
//
//	v.Is(v.Comparable(user).Not().Nil())
func Comparable[T comparable](value T, nameAndTitle ...string) *ValidatorComparable[T] {
	return &ValidatorComparable[T]{context: NewContext(value, nameAndTitle...)}
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorComparable[T]) Context() *ValidatorContext {
	return validator.context
}

// Invert the logical value associated with the next validator function.
// For example:
//
//	// It will return false because `Not()` inverts the boolean value associated with `EqualTo()`
//	v.Is(v.Comparable("a").Not().EqualTo("a")).Valid()
func (validator *ValidatorComparable[T]) Not() *ValidatorComparable[T] {
	validator.context.Not()
	return validator
}

// Or introduces a logical OR boundary in the current validator chain.
//
// Or groups adjacent validation fragments into a single OR-group that is
// evaluated left-to-right until one fragment succeeds. The OR-group succeeds
// if any fragment succeeds; it fails only if all fragments fail.
//
// Precedence: the OR-group is evaluated as a unit before the implicit AND
// that continues the chain. For example:
//
//	A.Or().B.C   == (A OR B) AND C
//
// Error reporting: if the OR-group fails, the error message for that group is
// a single message composed by joining the failing fragments' messages using
// the localized OR list format.
//
// Example:
//
//	// Passes because status equals "running" (EqualTo("paused") OR EqualTo("running")).
//	status := "running"
//	isValid := v.Is(v.Comparable(status).EqualTo("paused").Or().EqualTo("running")).Valid()
func (validator *ValidatorComparable[T]) Or() *ValidatorComparable[T] {
	validator.context.Or()
	return validator
}

// OrElse introduces a logical OR boundary with a cut (short-circuit) in the
// validator chain.
//
// OrElse behaves like Or for building an OR-group, but with an additional rule:
// if the left side (a single fragment, or the entire OR-group accumulated to
// the left of OrElse) succeeds, validation stops and no fragments to the right
// of OrElse are evaluated.
//
// This is primarily used to express "accept X, otherwise validate the rest"
// without repeating X across multiple OR fragments.
//
// Precedence: OrElse still participates in OR-grouping precedence. For example:
//
//	A.OrElse().B.C  == A OR (B AND C)   (with a cut if A succeeds)
//
// Error reporting: if the OR-group fails, its message is composed the same way
// as Or (localized OR list join).
//
// Example:
//
//	// If status equals "pending", the chain succeeds and InSlice is not evaluated.
//	// Otherwise, status must be in the validStatuses slice.
//	status := "pending"
//	validStatuses := []string{"active", "inactive"}
//	isValid := v.Is(v.Comparable(status).EqualTo("pending").OrElse().InSlice(validStatuses)).Valid()
func (validator *ValidatorComparable[T]) OrElse() *ValidatorComparable[T] {
	validator.context.OrElse()
	return validator
}

// Validate if a value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.Comparable(status).EqualTo("running"))
func (validator *ValidatorComparable[T]) EqualTo(value T, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			return validator.context.Value() == value
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a value passes a custom function.
// The function receives a typed T value, enabling compile-time type safety.
//
// For example:
//
//	type Status string
//	status := Status("running")
//	isValid := v.Is(
//	  v.Comparable(status).Passing(func(s Status) bool {
//	    return s == "running" || s == "paused"
//	  }),
//	).Valid()
func (validator *ValidatorComparable[T]) Passing(function func(v T) bool, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			// Value is stored as interface{} inside the context; assert back to T.
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, validator.context.Value(), template...)

	return validator
}

// Validate if a value is present in a slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.Comparable(status).InSlice(validStatus))
func (validator *ValidatorComparable[T]) InSlice(slice []T, template ...string) *ValidatorComparable[T] {
	validator.context.AddWithValue(
		func() bool {
			v := validator.context.Value().(T)
			for _, s := range slice {
				if v == s {
					return true
				}
			}
			return false
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}
