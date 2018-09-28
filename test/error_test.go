package test

import (
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestNotError(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := valgo.IsString(value).Blank()
		assert.True(t, v.Valid())
		assert.NoError(t, v.Error())
	}
}

func TestError(t *testing.T) {
	v := valgo.IsString("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	assert.Error(t, v.Error())
}
