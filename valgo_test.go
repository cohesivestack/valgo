package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.False(t, Is("Elon Musk").Empty().Valid())

	assert.True(t, Is("").Empty().Valid())
}

func TestBlank(t *testing.T) {
	assert.False(t, Is("Elon Musk").Blank().Valid())

	assert.True(t, Is(" ").Blank().Valid())

	assert.True(t, Is("").Blank().Valid())
}
