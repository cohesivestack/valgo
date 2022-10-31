package valgo

import (
	"fmt"
	"log"
)

type locale struct {
	Messages map[string]string
}

var locales map[string]*locale

var defaultLocaleCode string

func init() {
	setDefaultEnglishMessages()
	setDefaultSpanishMessages()

	err := SetDefaultLocale("en")
	if err != nil {
		log.Fatal(err)
	}
}

func getLocales() map[string]*locale {
	if locales == nil {
		locales = map[string]*locale{}
	}
	return locales
}

func getDefaultLocale() *locale {
	return getLocales()[defaultLocaleCode]
}

// Get the default locale code.
func GetDefaultLocaleCode() string {
	return defaultLocaleCode
}

// Set the default locale.
func SetDefaultLocale(code string) error {
	if _, exist := getLocales()[code]; exist {
		defaultLocaleCode = code
		return nil
	} else {
		return fmt.Errorf("there is not a locale registered with code %s", code)
	}

}

// Add or change the messages of a specific locale.
func SetLocaleMessages(code string, messages map[string]string) {
	getLocales()[code] = &locale{Messages: messages}
}

// Get the messages of a specific locale.
func GetLocaleMessages(code string) (messages map[string]string, err error) {
	if _, exist := getLocales()[code]; exist {
		messages = map[string]string{}
		for k, v := range getLocales()[code].Messages {
			messages[k] = v
		}
	} else {
		err = fmt.Errorf("there is not a locale registered with code %s", code)
	}
	return
}
