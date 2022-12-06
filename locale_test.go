package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLocalization(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	err := valgo.SetDefaultLocale("es")
	require.NoError(t, err)

	v := valgo.Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, v.ErrorByKey("value_0").Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent
	v = valgo.Is(valgo.String(" ").Empty())
	assert.Contains(t, v.ErrorByKey("value_0").Messages(), "Value 0 debe estar vacío")
}

func TestSeparatedLocalization(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	localized, err := valgo.Localized("es")
	assert.NoError(t, err)

	v := localized.New().Check(valgo.String(" ", "my_value").Not().Blank().Empty())
	assert.Contains(t, v.ErrorByKey("my_value").Messages(), "My value no puede estar en blanco")
	assert.Contains(t, v.ErrorByKey("my_value").Messages(), "My value debe estar vacío")

	// Default localization must not be changed
	v = valgo.Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, v.ErrorByKey("value_0").Messages(), "Value 0 can't be blank")
}

func TestAddLocalization(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	messages, err := valgo.GetLocaleMessages("en")
	assert.NoError(t, err)

	messages[valgo.ErrorKeyNotBlank] = "{{title}} ei tohi olla tühi"

	valgo.SetLocaleMessages("ee", messages)
	assert.NoError(t, err)

	localized, err := valgo.Localized("ee")
	assert.NoError(t, err)

	v := localized.New().Check(valgo.String(" ", "my_value").Not().Blank())
	assert.Contains(t, v.ErrorByKey("my_value").Messages(), "My value ei tohi olla tühi")

	// Default localization must not be changed
	v = valgo.Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, v.ErrorByKey("value_0").Messages(), "Value 0 can't be blank")
}
