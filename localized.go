package valgo

type localized struct {
	_locale *locale
}

func (l *localized) NewValidator() Validator {
	return newValidator(l._locale)
}
