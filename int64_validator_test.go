package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64EqualToValid(t *testing.T) {
	ResetMessages()

	v := IsInt64(int64(10)).EqualTo(int64(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(1)).EqualTo(int64(10))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}

func TestInt64GreaterThanValid(t *testing.T) {
	ResetMessages()

	v := IsInt64(int64(10)).GreaterThan(int64(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(1)).GreaterThan(int64(10))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}

func TestInt64LessThanValid(t *testing.T) {
	ResetMessages()

	v := IsInt64(int64(5)).LessThan(int64(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(10)).LessThan(int64(1))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}

func TestInt64GreaterOrEqualThanValid(t *testing.T) {
	ResetMessages()

	v := IsInt64(int64(10)).GreaterOrEqualThan(int64(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(10)).GreaterOrEqualThan(int64(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(1)).GreaterOrEqualThan(int64(10))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}

func TestInt64LessOrEqualThanValid(t *testing.T) {
	ResetMessages()

	v := IsInt64(int64(5)).LessOrEqualThan(int64(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(5)).LessOrEqualThan(int64(5))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = IsInt64(int64(10)).LessOrEqualThan(int64(1))
	assert.False(t, v.Valid())
	assert.Len(t, v.Errors(), 1)
}
