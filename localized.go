package valgo

type localized struct {
	_locale *locale
}

func (l *localized) IsString(value string, nameAndTitle ...string) *Validator {
	return newValidator(l._locale).IsString(value, nameAndTitle...)
}
