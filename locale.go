package valgo

import (
	"errors"
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

func SetDefaultLocale(code string) error {
	if _, exist := getLocales()[code]; exist {
		defaultLocaleCode = code
		return nil
	} else {
		return errors.New(fmt.Sprintf("There is not a locale registered with code %s", code))
	}

}
