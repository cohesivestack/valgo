package valgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var locales map[string]locale

var defaultLocale locale

func init() {
	locales = map[string]locale{}
	SetDefaultLocale("en")
}

func Is(value interface{}) *Validator {
	return newValidator(defaultLocale, value)
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

func SetDefaultLocale(code string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(ex)

	if _locale, exist := locales[code]; exist {
		defaultLocale = _locale

		// Parse config yml file
		filePath, _ := filepath.Abs(path.Join(dir, "messages", fmt.Sprintf("%s.yml", code)))

		yamlFile, _ := ioutil.ReadFile(filePath)

		defaultLocale := locale{}
		err = yaml.Unmarshal(yamlFile, &defaultLocale)
		if err != nil {
			panic(err)
		}
	} else {
		filePath, _ := filepath.Abs(path.Join(dir, "messages", fmt.Sprintf("%s.yml", code)))
		AddOrReplaceLocale(filePath, code)

		defaultLocale = locales[code]
	}

}

func AddOrReplaceLocale(code string, filePath string) error {
	yamlFile, _ := ioutil.ReadFile(filePath)

	_locale := locale{}
	err := yaml.Unmarshal(yamlFile, &_locale)
	if err != nil {
		return err
	}

	locales[code] = _locale

	return nil
}

func Localized(code string) (*localized, error) {
	if _locale, exist := locales[code]; exist {
		return &localized{
			_locale: _locale,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Doesn't exist a registered locale with code '%s'", code))
	}
}
