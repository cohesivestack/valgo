package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_localeDoesNotExist(t *testing.T) {
	t.Parallel()

	assert.Error(t, localeDoesNotExist("n/a"))
}
