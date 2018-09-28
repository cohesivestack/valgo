package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	v := valgo.IsString("golang").InSlice([]interface{}{"swift", "golang", "kotlin"})
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringInSliceInvalid(t *testing.T) {
	for i, slice := range [][]interface{}{
		[]interface{}{"dart", "typescript"},
		[]interface{}{},
		[]interface{}{"perl"}} {

		v := valgo.IsString("golang").InSlice(slice)
		m := fmt.Sprintf("not assert using options '%v'", i)

		assert.False(t, v.Valid(), m)
		if assert.NotEmpty(t, v.Errors(), m) {
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				"golang is not a value valid for Value 0", m)
		}
	}
}
