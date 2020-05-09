package valgo

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchingToValid(t *testing.T) {
	ResetMessages()

	pattern := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	value := "vitalik[10]"

	v := IsString(value).MatchingTo(pattern)
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestMatchingToInvalid(t *testing.T) {
	ResetMessages()

	pattern := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	value := "Vitalik[10]"

	v := IsString(value).MatchingTo(pattern)

	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Contains(t, v.Errors()["value_0"].Messages(),
		fmt.Sprintf("Value 0 must match to \"%v\"", pattern))
}
