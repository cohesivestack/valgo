package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestMatchingToValid(t *testing.T) {
	valgo.ResetMessages()

	pattern := `^[a-z]+\[[0-9]+\]$`
	value := "vitalik[10]"

	for description, value := range map[string]interface{}{
		"literal": value,
		"pointer": &value,
	} {
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(value).MatchingTo(pattern)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestMatchingToInvalid(t *testing.T) {
	valgo.ResetMessages()

	pattern := `^[a-z]+\[[0-9]+\]$`
	value := "Vitalik[10]"

	for description, value := range map[string]interface{}{
		"integer": 10,
		"struct":  MyStruct{FieldString: value},
		"slice":   []string{value},
		"literal": value,
		"pointer": &value,
	} {
		v := valgo.Is(value).MatchingTo(pattern)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" must match to \"%v\"", pattern), msg)
		}
	}
}
