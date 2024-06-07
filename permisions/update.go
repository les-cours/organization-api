package permisions

import (
	"context"
	"github.com/les-cours/learning-api/types"
)

func CanUpdate(ctx context.Context) (*types.UserToken, error) {

	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return nil, ErrAuth
	}

	if !user.Update.LEARNING {
		return nil, ErrPermissionDenied
	}

	return user, nil
}
