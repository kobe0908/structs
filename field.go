package structs

import (
	"errors"
)
var (
	errNotExported = errors.New("field is not exported")
	errNotSettable = errors.New("field is not settable")
)

