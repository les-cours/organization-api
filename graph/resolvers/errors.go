package resolvers

import (
	"errors"
	"google.golang.org/grpc/status"
)

func ErrApi(err error) error {
	return errors.New(status.Convert(err).Message())
}

var ErrPermissionDenied = errors.New("you're not allowed to do that. This resource is off-limits")
