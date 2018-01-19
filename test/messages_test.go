package test

import (
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestCustomMessageTemplate(t *testing.T) {
	v := valgo.Is(" ").NotBlank("The field \"{{Title}}\" can't be empty")
	assert.Contains(t, v.Errors()[0].Messages, "The field \"value0\" can't be empty")

	// Must not replace default blank message
	v = valgo.Is(" ").NotBlank()
	assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be blank")
}
