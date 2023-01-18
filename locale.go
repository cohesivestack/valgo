package valgo

const (
	LocaleCodeEn = "en"
	LocaleCodeEs = "es"
)

const localeDefault = LocaleCodeEn

type Locale struct {
	Messages map[string]string
}

func getLocale(code string) *Locale {
	switch code {
	case LocaleCodeEs:
		return getLocaleEs()
	default:
		return getLocaleEn()
	}
}

func (_locale *Locale) merge(locale *Locale) {
	for k, v := range locale.Messages {
		_locale.Messages[k] = v
	}
}
