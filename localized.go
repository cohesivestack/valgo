package valgo

type localized struct {
	_locale *locale
}

// Create a new localized Validation session.
func (l *localized) New() *Validation {
	return newValidation(l._locale)
}
