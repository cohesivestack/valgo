package valgo

const (
	LocaleCodeEn = "en"
	LocaleCodeEs = "es"
)

const localeDefault = LocaleCodeEn

type Locale struct {
	Messages map[string]string
}

func getLocaleWithSkipDefaultOption(code string, skipDefault bool, factoryLocales ...map[string]*Locale) *Locale {

	if len(factoryLocales) > 0 && factoryLocales[0] != nil {

		if locale, exists := factoryLocales[0][code]; exists {
			return locale
		}
		if skipDefault {
			return nil
		}
		return getLocaleEn()

	} else {

		switch code {
		case LocaleCodeEs:
			return getLocaleEs()
		default:
			if skipDefault {
				return nil
			}
			return getLocaleEn()
		}

	}
}

func getLocaleAndSkipDefaultOption(code string, factoryLocales ...map[string]*Locale) *Locale {
	return getLocaleWithSkipDefaultOption(code, true, factoryLocales...)
}

func getLocale(code string, factoryLocales ...map[string]*Locale) *Locale {
	return getLocaleWithSkipDefaultOption(code, false, factoryLocales...)
}

func (_locale *Locale) merge(locale *Locale) {
	for k, v := range locale.Messages {
		_locale.Messages[k] = v
	}
}
