package valgo

type localized struct {
	_locale *locale
}

func (l *localized) IsString(value string, nameAndTitle ...string) *StringValidator {
	return l.NewValidator().IsString(value, nameAndTitle...)
}

func (l *localized) CheckString(value string, nameAndTitle ...string) *StringValidator {
	return l.NewValidator().CheckString(value, nameAndTitle...)
}

func (l *localized) IsInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return l.NewValidator().IsInt64(value, nameAndTitle...)
}

func (l *localized) CheckInt64(value int64, nameAndTitle ...string) *Int64Validator {
	return l.NewValidator().CheckInt64(value, nameAndTitle...)
}

func (l *localized) NewValidator() Validator {
	return newValidator(l._locale)
}
