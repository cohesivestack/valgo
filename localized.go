package valgo

type localized struct {
	_locale *locale
}

func (l *localized) New() *Validation {
	return newValidation(l._locale)
}
