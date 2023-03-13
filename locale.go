package valgo

const (
	LocaleCodeEn = "en"
	LocaleCodeEs = "es"
	LocaleCodeDe = "de"
)

const localeCodeDefault = LocaleCodeEn

// Locale is a type alias that represents a map of locale entries.
// The keys in the map are strings that represent the entry's identifier, and
// the values are strings that contain the corresponding localized text
// for that entry
type Locale map[string]string

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
		case LocaleCodeEn:
			return getLocaleEn()
		case LocaleCodeDe:
			return getLocaleDe()
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

func (_locale *Locale) merge(locale *Locale) *Locale {
	if locale != nil {
		for k, v := range *locale {
			(*_locale)[k] = v
		}
	}

	return _locale
}
