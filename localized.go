package valgo

type localized struct {
	_locale locale
}

func (l *localized) Is(value interface{}) *Validator {
	return newValidator(l._locale, value)
}
