package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEmpty(t *testing.T) {
	v := IsString("").Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringEmptyInvalid(t *testing.T) {
	for _, value := range []string{" ", "Vitalik Buterin"} {
		v := IsString(value).Empty()
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.False(t, v.Valid(), m)
		if assert.NotEmpty(t, v.Errors(), m) {
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				"Value 0 must be empty", m)
		}
	}
}

func TestStringNotEmptyValid(t *testing.T) {
	for _, value := range []string{" ", "Vitalik Buterin"} {
		v := IsString(value).Not().Empty()
		m := fmt.Sprintf("not assert using '%s'", value)

		assert.True(t, v.Valid(), m)
		assert.Empty(t, v.Errors(), m)
	}
}

func TestStringNotEmptyInvalid(t *testing.T) {
	v := IsString("").Not().Empty()
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 can't be empty")
}

// Benchmarks

func BenchmarkStringEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = IsString("").Empty()
	}
}

func BenchmarkStringEmptyInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = IsString("Vitalik Buterin").Empty()
	}
}

func BenchmarkNotEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = IsString("Vitalik Buterin").Not().Empty()
	}
}

func BenchmarkNotEmptyInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = IsString("").Not().Empty()
	}
}
