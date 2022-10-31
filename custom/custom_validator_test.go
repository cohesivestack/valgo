package custom

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidator(t *testing.T) {
	valgo.TeardownTest()

	errorMessages, err := valgo.GetLocaleMessages(valgo.GetDefaultLocaleCode())
	assert.NoError(t, err)

	errorMessages["not_valid_secret"] = "{{title}} is invalid."
	valgo.SetLocaleMessages(valgo.GetDefaultLocaleCode(), errorMessages)

	v := valgo.Is(SecretWord("loose", "secret").Correct())

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Len(t, v.Errors()["secret"].Messages(), 1)
	assert.Contains(t, v.Errors()["secret"].Messages(), "Secret is invalid.")

	v = valgo.Is(SecretWord("cohesive").Correct())
	assert.True(t, v.Valid())
	assert.Len(t, v.Errors(), 0)

}

type ValidatorSecretWord struct {
	context *valgo.ValidatorContext
}

func (validator *ValidatorSecretWord) Correct(template ...string) *ValidatorSecretWord {
	validator.context.Add(
		func() bool {
			return validator.context.Value().(string) == "cohesive" ||
				validator.context.Value().(string) == "stack"
		},
		"not_valid_secret", template...)

	return validator
}

func (validator *ValidatorSecretWord) Context() *valgo.ValidatorContext {
	return validator.context
}

func SecretWord(value string, nameAndTitle ...string) *ValidatorSecretWord {
	return &ValidatorSecretWord{context: valgo.NewContext(value, nameAndTitle...)}
}
