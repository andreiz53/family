package webUtils

import (
	"context"
	"family/types"
)

func User(ctx context.Context) types.AuthenticatedUser {
	u, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)
	if !ok || !u.LoggedIn {
		return types.AuthenticatedUser{}
	}
	return u
}
func LoggedIn(ctx context.Context) bool {
	u, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)
	return u.LoggedIn && ok
}
