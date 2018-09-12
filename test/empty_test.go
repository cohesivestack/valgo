package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStringValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestEmptyStringInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []string{"Vitalik Buterin", " "} {
		v := valgo.Is(value).Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %s", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty", fmt.Sprintf("not assert using %s", value))
		}
	}
}

func TestNotEmptyStringValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []string{"Vitalik Buterin", " "} {
		v := valgo.Is(value).Not().Empty()

		assert.True(t, v.Valid())
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestNotEmptyStringInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("").Not().Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}
}

func TestEmptyNumberValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{0, 0.0} {
		v := valgo.Is(value).Empty()
		assert.True(t, v.Valid(), fmt.Sprintf("not assert using %v", value))
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %v", value))
	}
}

func TestEmptyNumberInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{1, 1.1} {
		v := valgo.Is(value).Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %v", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty", fmt.Sprintf("not assert using %v", value))
		}
	}
}

func TestNotEmptyNumberValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{1, 1.1} {
		v := valgo.Is(value).Not().Empty()

		assert.True(t, v.Valid(), fmt.Sprintf("not assert using %v", value))
		assert.Empty(t, v.Errors(), fmt.Sprintf("not assert using %v", value))
	}
}

func TestNotEmptyNumberInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{0, 0.0} {
		v := valgo.Is(value).Not().Empty()

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors()) {
			assert.Len(t, v.Errors(), 1, fmt.Sprintf("not assert using %v", value))
			assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty", fmt.Sprintf("not assert using %v", value))
		}
	}
}

func TestEmptySliceValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{}).Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestEmptySliceInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{0}).Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")
	}
}

func TestNotEmptySliceValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{0}).Not().Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestNotEmptySliceInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is([]int{}).Not().Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}
}

func TestEmptyMapValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{}).Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestEmptyMapInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{"a": 0}).Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" must be empty")
	}
}

func TestNotEmptyMapValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{"a": 0}).Not().Empty()
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestNotEmptyMapInvalid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is(map[string]int{}).Not().Empty()
	assert.False(t, v.Valid())
	if assert.NotEmpty(t, v.Errors()) {
		assert.Len(t, v.Errors(), 1)
		assert.Contains(t, v.Errors()[0].Messages, "\"value0\" can't be empty")
	}
}
