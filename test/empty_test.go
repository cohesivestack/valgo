package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	for _, value := range []string{"Vitalik Buterin", " "} {
		v = valgo.Is(value).Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty", fmt.Sprintf("not assert using %s", value))
		}
	}
}

func TestNotEmptyString(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").NotEmpty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}

	for _, value := range []string{"Vitalik Buterin", " "} {
		v = valgo.Is(value).NotEmpty()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestEmptyNumber(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{0, 0.0} {
		v := valgo.Is(value).Empty()
		assert.True(t, v.Valid(), fmt.Sprintf("not assert using %v", value))
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %v", value))
	}

	for _, value := range []interface{}{1, 1.1} {
		v := valgo.Is(value).Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %v", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty", fmt.Sprintf("not assert using %v", value))
		}
	}
}

func TestNotEmptyNumber(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{0, 0.0} {
		v := valgo.Is(value).NotEmpty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %v", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty", fmt.Sprintf("not assert using %v", value))
		}
	}

	for _, value := range []interface{}{1, 1.1} {
		v := valgo.Is(value).NotEmpty()

		assert.True(t, v.Valid(), fmt.Sprintf("not assert using %v", value))
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %v", value))
	}
}

func TestEmptySlice(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{}).Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is([]int{0}).Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")
	}
}

func TestNotEmptySlice(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{}).NotEmpty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}

	v = valgo.Is([]int{0}).NotEmpty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestEmptyMap(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{}).Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(map[string]int{"a": 0}).Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")
	}
}

func TestNotEmptyMap(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{}).NotEmpty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}

	v = valgo.Is(map[string]int{"a": 0}).NotEmpty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
