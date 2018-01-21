package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomMessageTemplate(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(" ").NotBlank("The field \"{{Title}}\" can't be blank. :-)").Empty()
	assert.Contains(t, v.Errors()[0].Messages, "The field \"value0\" can't be blank. :-)")

	// Should not replace the template for other validations
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")

	// Should not replace default blank message
	v = valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank")
}
