package valgo

import (
	"log"
	"sync"
)

type locale struct {
	Messages map[string]string
}

type existingLocales struct {
	sync.RWMutex
	locales map[string]*locale
}

var locales existingLocales

var defaultLocaleCode string

//nolint:gochecknoinits // by design. must be refactored
func init() {
	locales = existingLocales{locales: make(map[string]*locale)}

	SetDefaultEnglishMessages()
	SetDefaultSpanishMessages()

	if err := SetDefaultLocale("en"); err != nil {
		log.Fatal(err)
	}
}

func getLocale(code string) *locale {
	if len(locales.locales) == 0 {
		return nil
	}

	locales.RLock()

	l, ok := locales.locales[code]
	if !ok {
		return nil
	}

	locales.RUnlock()

	return l
}

func getDefaultLocale() *locale {
	return getLocale(defaultLocaleCode)
}

// GetDefaultLocaleCode Get the default locale code.
func GetDefaultLocaleCode() string {
	return defaultLocaleCode
}

// SetDefaultLocale Set the default locale.
func SetDefaultLocale(code string) error {
	if l := getLocale(code); l == nil {
		return localeDoesNotExist(code)
	}

	defaultLocaleCode = code

	return nil
}

// SetLocaleMessages Add or change the messages of a specific locale.
func SetLocaleMessages(code string, messages map[string]string) {
	locales.Lock()
	locales.locales[code] = &locale{Messages: messages}
	locales.Unlock()
}

// GetLocaleMessages Get the messages of a specific locale.
func GetLocaleMessages(code string) (map[string]string, error) {
	l := getLocale(code)
	if l == nil {
		return nil, localeDoesNotExist(code)
	}

	messages := make(map[string]string)

	for k, v := range l.Messages {
		messages[k] = v
	}

	return messages, nil
}
