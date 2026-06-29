package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type validatorContextLocaleFallbackValidator struct {
	context *ValidatorContext
}

func newValidatorContextLocaleFallbackValidator(locales ...*Locale) *validatorContextLocaleFallbackValidator {
	context := NewContext("value", "field", "Field")
	context.WithLocaleFallback(locales...)

	return &validatorContextLocaleFallbackValidator{context: context}
}

func (validator *validatorContextLocaleFallbackValidator) Invalid(errorKey string) *validatorContextLocaleFallbackValidator {
	validator.context.Add(
		func() bool {
			return false
		},
		errorKey,
	)

	return validator
}

func (validator *validatorContextLocaleFallbackValidator) Context() *ValidatorContext {
	return validator.context
}

func TestValidatorContextWithLocaleFallbackUsesFallbackForMissingKey(t *testing.T) {
	locale := &Locale{
		"custom_error": "{{title}} uses the fallback message",
	}

	validation := New()
	originalLocale := validation._locale

	validation.Is(newValidatorContextLocaleFallbackValidator(locale).Invalid("custom_error"))

	assert.Contains(t, validation.Errors()["field"].Messages(), "Field uses the fallback message")
	assert.NotSame(t, originalLocale, validation._locale)
	assert.NotContains(t, *originalLocale, "custom_error")
}

func TestValidatorContextWithLocaleFallbackDoesNotOverrideActiveLocale(t *testing.T) {
	activeLocale := &Locale{
		"custom_error": "{{title}} uses the active locale message",
	}
	fallbackLocale := &Locale{
		"custom_error": "{{title}} uses the fallback message",
	}

	validation := New(Options{Locale: activeLocale})
	originalLocale := validation._locale

	validation.Is(newValidatorContextLocaleFallbackValidator(fallbackLocale).Invalid("custom_error"))

	assert.Contains(t, validation.Errors()["field"].Messages(), "Field uses the active locale message")
	assert.Same(t, originalLocale, validation._locale)
}

func TestValidatorContextWithLocaleFallbackDoesNotOverrideBuiltInLocale(t *testing.T) {
	locale := &Locale{
		ErrorKeyBlank: "{{title}} uses the fallback message",
	}

	validation := New()
	originalLocale := validation._locale

	validation.Is(newValidatorContextLocaleFallbackValidator(locale).Invalid(ErrorKeyBlank))

	assert.Contains(t, validation.Errors()["field"].Messages(), "Field must be blank")
	assert.Same(t, originalLocale, validation._locale)
}

func TestValidatorContextWithLocaleFallbackKeepsFirstFallbackEntry(t *testing.T) {
	firstLocale := &Locale{
		"custom_error": "{{title}} uses the first fallback message",
	}
	secondLocale := &Locale{
		"custom_error": "{{title}} uses the second fallback message",
	}

	validation := New().Is(
		newValidatorContextLocaleFallbackValidator(firstLocale, secondLocale).Invalid("custom_error"),
	)

	assert.Contains(t, validation.Errors()["field"].Messages(), "Field uses the first fallback message")
}
