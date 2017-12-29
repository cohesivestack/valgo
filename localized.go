package valgo

type localized struct {
	_locale locale
}

func (_localized *localized) Is(value interface{}) *Validator {
	return newValidator(_localized._locale, value)
}
