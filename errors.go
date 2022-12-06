package valgo

import (
	"errors"
	"fmt"
)

var ErrLocaleDoesntExist = errors.New("doesn't exist a registered locale with code")

func localeDoesNotExist(code string) error {
	return fmt.Errorf("%w '%s'", ErrLocaleDoesntExist, code)
}
