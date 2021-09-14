package valgo

import (
	"regexp"
	"strings"
)

type StringValidator struct {
	*validatorContext
}

func IsString(value string, nameAndTitle ...string) *StringValidator {
	return NewValidator().IsString(value, nameAndTitle...)
}

func CheckString(value string, nameAndTitle ...string) *StringValidator {
	return NewValidator().CheckString(value, nameAndTitle...)
}

func (v *validatorContext) IsString(value string, nameAndTitle ...string) *StringValidator {
	return v.isString(true, value, nameAndTitle...)
}

func (v *validatorContext) CheckString(value string, nameAndTitle ...string) *StringValidator {
	return v.isString(false, value, nameAndTitle...)
}

func (l *localized) IsString(value string, nameAndTitle ...string) *StringValidator {
	return l.NewValidator().IsString(value, nameAndTitle...)
}

func (l *localized) CheckString(value string, nameAndTitle ...string) *StringValidator {
	return l.NewValidator().CheckString(value, nameAndTitle...)
}

func (v *validatorContext) isString(shortCircuit bool, value string, nameAndTitle ...string) *StringValidator {
	v.currentDataType = DataTypeString
	v.currentValue = value
	v.currentIndex += 1
	v.currentValid = true
	v.shortCircuit = shortCircuit

	sizeNameAndTitle := len(nameAndTitle)
	if sizeNameAndTitle > 0 {
		v.currentName = &nameAndTitle[0]
		if sizeNameAndTitle > 1 {
			v.currentTitle = &nameAndTitle[1]
		}
	}
	return &StringValidator{v}
}

func IsBlank(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func (v *StringValidator) Blank(template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Blank validator requires a string as value")
	} else if !v.assert(IsBlank(v.currentValue.(string))) {
		v.invalidate("blank", map[string]interface{}{"title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}

func IsEmpty(value string) bool {
	return len(value) == 0
}

func (v *StringValidator) Empty(template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Empty validator requires a string as value")
	} else if !v.assert(IsEmpty(v.currentValue.(string))) {
		v.invalidate("empty", map[string]interface{}{"title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}

func IsEqualTo(valueA interface{}, valueB interface{}) bool {
	return valueA == valueB
}

func (v *StringValidator) EqualTo(value interface{}, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(IsEqualTo(v.currentValue, value)) {
		v.invalidate("equal_to", map[string]interface{}{
			"title": v.currentTitle,
			"value": value}, template...)
	}

	v.resetNegative()

	return v
}

func IsMatchingTo(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

func (v *StringValidator) MatchingTo(regex *regexp.Regexp, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Empty validator requires a string as value")
	} else if !v.assert(IsMatchingTo(v.currentValue.(string), regex)) {
		v.invalidate("matching_to", map[string]interface{}{
			"title": v.currentTitle,
			"value": regex.String()}, template...)
	}

	v.resetNegative()

	return v
}

func (v *StringValidator) MaxLength(length int, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("MaxLength validator requires a string as value")
	} else if !v.assert(len(v.currentValue.(string)) <= length) {
		v.invalidate("max_length", map[string]interface{}{"title": v.currentTitle, "length": length}, template...)
	}

	v.resetNegative()

	return v
}

func (v *StringValidator) MinLength(length int, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("MinLength validator requires a string as value")
	} else if !v.assert(len(v.currentValue.(string)) >= length) {
		v.invalidate("min_length", map[string]interface{}{"title": v.currentTitle, "length": length}, template...)
	}

	v.resetNegative()

	return v
}

func (v *StringValidator) Length(length int, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Length validator requires a string as value")
	} else if !v.assert(len(v.currentValue.(string)) == length) {
		v.invalidate("length", map[string]interface{}{"title": v.currentTitle, "length": length}, template...)
	}

	v.resetNegative()

	return v
}

func (v *StringValidator) Not() *StringValidator {
	v.currentNegative = true

	return v
}

func (v *StringValidator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *StringValidator {

	if v.isShortCircuit() {
		return v
	}

	customValidator := CustomValidator{
		validator: v.validatorContext,
	}

	function(&customValidator, template...)

	v.resetNegative()

	return v
}

func IsInSlice(value interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func IsStringInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func (v *StringValidator) InSlice(slice []string, template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if !v.assert(IsStringInSlice(v.currentValue.(string), slice)) {
		v.invalidate("in_slice", map[string]interface{}{
			"title": v.currentTitle,
			"value": v.currentValue}, template...)
	}

	v.resetNegative()

	return v
}

const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

var emailRegex *regexp.Regexp

func init() {
	emailRegex = regexp.MustCompile(emailRegexPattern)
}

func IsAnEmail(value string) bool {
	return IsMatchingTo(value, emailRegex)
}

func (v *StringValidator) AnEmail(template ...string) *StringValidator {
	if v.isShortCircuit() {
		return v
	} else if v.currentDataType != DataTypeString {
		panic("Email validator requires a string as value")
	} else if !v.assert(IsAnEmail(v.currentValue.(string))) {
		v.invalidate("an_email", map[string]interface{}{"title": v.currentTitle}, template...)
	}

	v.resetNegative()

	return v
}
