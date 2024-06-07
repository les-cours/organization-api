package permisions

import (
	"context"
	"github.com/les-cours/learning-api/types"
)

func Student(ctx context.Context) (*types.UserToken, error) {

	user, ok := ctx.Value("user").(*types.UserToken)

	switch {
	case !ok || *user == (types.UserToken{}):
		return nil, ErrAuth
	case user.UserType != "student":
		return nil, ErrPermissionDenied

	}

	return user, nil
}
