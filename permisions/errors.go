package permisions

import (
	"errors"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrAuth             = errors.New("authentication failed")
)
