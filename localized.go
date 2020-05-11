package valgo

type localized struct {
	_locale *locale
}

func (l *localized) IsString(value string, nameAndTitle ...string) *Validator {
	return l.NewValidator().IsString(value, nameAndTitle...)
}

func (l *localized) CheckString(value string, nameAndTitle ...string) *Validator {
	return l.NewValidator().CheckString(value, nameAndTitle...)
}

func (l *localized) NewValidator() *Validator {
	return newValidator(l._locale)
}
