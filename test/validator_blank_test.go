package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestStringBlank(t *testing.T) {
	for _, value := range []string{"", " "} {
		v := valgo.IsString(value).Blank()
		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestStringBlankInvalid(t *testing.T) {
	v := valgo.IsString("Vitalik Buterin").Blank()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t,
			v.Errors()["value_0"].Messages(),
			"Value 0 must be blank")
	}
}

func TestStringNotBlankValid(t *testing.T) {
	v := valgo.IsString("Vitalik Buterin").Not().Blank()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestStringNotBlankInvalid(t *testing.T) {
	for _, value := range []string{" ", ""} {
		v := valgo.IsString(value).Not().Blank()
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
		_ = valgo.IsString(" ").Blank()
	}
}

func BenchmarkStringBlankInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = valgo.IsString("Vitalik Buterin").Blank()
	}
}

func BenchmarkNotBlank(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = valgo.IsString("Vitalik Buterin").Not().Blank()
	}
}

func BenchmarkNotBlankInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = valgo.IsString(" ").Not().Blank()
	}
}
