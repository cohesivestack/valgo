package valgo

type localized struct {
	_locale *locale
}

func (l *localized) New() *ValidatorGroup {
	return newValidatorGroup(l._locale)
}
