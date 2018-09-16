package test

import (
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomMessageTemplate(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(" ").Not().Blank("The field \"{{Title}}\" can't be blank. :-)").Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "The field \"Value 0\" can't be blank. :-)")
}

func TestCustomMessageTemplateIsNotReplacingOtherValidations(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(" ").Not().Blank("The field \"{{Title}}\" can't be blank. :-)").Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "The field \"Value 0\" can't be blank. :-)")

	// Should not replace the template for other validations
	assert.Contains(t, v.ErrorItems()[0].Messages, "Value 0 must be empty")
}

func TestCustomMessageTemplateIsNotReplacingDefaultValidation(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(" ").Not().Blank("The field \"{{Title}}\" can't be blank. :-)").Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "The field \"Value 0\" can't be blank. :-)")

	// Should not replace default blank message
	v = valgo.Is(" ").Not().Blank()
	assert.Contains(t, v.ErrorItems()[0].Messages, "Value 0 can't be blank")
}

func TestCustomMessagesAddingANewLocale(t *testing.T) {
	valgo.ResetMessages()

	valgo.AddOrReplaceMessages("liberland", map[string]string{
		"not_blank": "Liberland say \"{{Title}}\" can't be blank",
		"empty":     "Liberland say \"{{Title}}\" must be empty",
	})

	localized, err := valgo.Localized("liberland")
	assert.NoError(t, err)

	v := localized.Is(" ").Not().Blank().Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "Liberland say \"Value 0\" can't be blank")
	assert.Contains(t, v.ErrorItems()[0].Messages, "Liberland say \"Value 0\" must be empty")

}

func TestCustomMessagesReplacingExistingLocale(t *testing.T) {
	valgo.ResetMessages()

	valgo.AddOrReplaceMessages("en", map[string]string{
		"not_blank": "An improved english say \"{{Title}}\" can't be blank",
		"empty":     "An improved english say \"{{Title}}\" must be empty",
	})

	localized, err := valgo.Localized("en")
	assert.NoError(t, err)

	v := localized.Is(" ").Not().Blank().Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "An improved english say \"Value 0\" can't be blank")
	assert.Contains(t, v.ErrorItems()[0].Messages, "An improved english say \"Value 0\" must be empty")

}

func TestGetMessagesIsACopy(t *testing.T) {
	valgo.ResetMessages()

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	messages, err := valgo.GetMessagesCopy("en")
	assert.NoError(t, err)

	assert.True(t, len(messages) > 1)

	// Check that is a real copy
	messages["not_blank"] = "This message should not be assigned"

	v := valgo.Is(" ").Not().Blank()
	assert.Contains(t, v.ErrorItems()[0].Messages, "Value 0 can't be blank")
	assert.NotContains(t, v.ErrorItems()[0].Messages, "This message should not be assigned")
}

func TestACopyCanBeUsedToReplaceALocale(t *testing.T) {
	valgo.ResetMessages()

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	messages, err := valgo.GetMessagesCopy("en")
	assert.NoError(t, err)
	assert.True(t, len(messages) > 1)

	messages["not_blank"] = "This message is changed for test purposes"

	valgo.AddOrReplaceMessages("en", messages)

	v := valgo.Is(" ").Not().Blank().Empty()
	assert.Contains(t, v.ErrorItems()[0].Messages, "This message is changed for test purposes")
	assert.Contains(t, v.ErrorItems()[0].Messages, "Value 0 must be empty")
}

func TestWrongMessageKey(t *testing.T) {
	valgo.ResetMessages()

	err := valgo.SetDefaultLocale("en")
	assert.NoError(t, err)

	valgo.AddOrReplaceMessages("en", map[string]string{})

	v := valgo.Is(" ").Not().Blank()
	assert.Contains(t, v.ErrorItems()[0].Messages,
		"ERROR: THERE IS NOT A MESSAGE WITH THE KEY \"not_blank\"!")

}

func TestMissingMessageKey(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("USD").Passing(func(_v *valgo.CustomValidator, _t ...string) {
		if _v.Value().AsString() != "BTC" {
			// Here the missing key
			_v.Invalidate(" ", _t, nil)
		}
	})

	assert.Contains(t, v.ErrorItems()[0].Messages,
		"ERROR: MISSING MESSAGE KEY OR TEMPLATE STRING!")
}
