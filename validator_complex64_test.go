package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplex64EqualToValid(t *testing.T) {
	ResetMessages()

	v := Is(Complex64(complex64(10)).EqualTo(complex64(10)))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Complex64(complex64(1)).EqualTo(complex64(10)))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}
