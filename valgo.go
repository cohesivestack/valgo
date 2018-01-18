package valgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var locales map[string]locale

var defaultLocale locale

func init() {
	err := SetDefaultLocale("en")
	if err != nil {
		log.Fatal(err)
	}
}

func getLocales() map[string]locale {
	if locales == nil {
		locales = map[string]locale{}
	}
	return locales
}

func SetDefaultLocale(code string) error {

	if _locale, exist := getLocales()[code]; exist {
		defaultLocale = _locale
		return nil
	} else {
		return errors.New(fmt.Sprintf("There is not a locale registered with code %s", code))
	}

}

func AddOrReplaceLocaleFromYaml(code string, filePath string) error {
	yamlFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	_locale := locale{}
	err = yaml.Unmarshal(yamlFile, &_locale)
	if err != nil {
		return err
	}

	getLocales()[code] = _locale

	return nil
}

func AddOrReplaceLocale(code string, messages map[string]string) {
	_locale := locale{
		Messages: messages,
	}

	getLocales()[code] = _locale
}

func Localized(code string) (*localized, error) {
	if _locale, exist := getLocales()[code]; exist {
		return &localized{
			_locale: _locale,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Doesn't exist a registered locale with code '%s'", code))
	}
}

func newValidator(_locale locale, value interface{}) *Validator {
	validator := &Validator{
		currentIndex: 0,
		currentValue: value,
		currentValid: true,
		valid:        true,
		_locale:      defaultLocale,
	}
	validator.currentName = fmt.Sprintf("value%v", validator.currentIndex)
	validator.currentTitle = validator.currentName
	return validator
}

func Is(value interface{}) *Validator {
	return newValidator(defaultLocale, value)
}
