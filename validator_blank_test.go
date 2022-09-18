package valgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringBlank(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := Is(String(value).Blank())
		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestStringBlankInvalid(t *testing.T) {
	v := Is(String("Vitalik Buterin").Blank())
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t,
			v.Errors()["value_0"].Messages(),
			"Value 0 must be blank")
	}
}

func TestStringNotBlankValid(t *testing.T) {
	v := Is(String("Vitalik Buterin").Not().Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringNotBlankInvalid(t *testing.T) {
	for _, value := range []string{" ", ""} {
		v := Is(String(value).Not().Blank())
		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t,
				v.Errors()["value_0"].Messages(),
				"Value 0 can't be blank",
				fmt.Sprintf("not assert using %s", value))
		}
	}
}

// Benchmarks
func BenchmarkStringBlank(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Is(String(" ").Blank())
	}
}

func BenchmarkStringBlankInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Is(String("Vitalik Buterin").Blank())
	}
}

func BenchmarkNotBlank(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Is(String("Vitalik Buterin").Not().Blank())
	}
}

func BenchmarkNotBlankInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Is(String(" ").Not().Blank())
	}
}
