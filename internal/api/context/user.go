package context

import (
	"context"

	"github.com/hasahmad/go-api/internal/models"
)

type userKey string

const userContextKey = userKey("user")

func SetUser(ctx context.Context, user models.User) context.Context {
	if user.IsAnonymousUser() {
		return ctx
	}

	if ctx.Value(userContextKey) != nil {
		return ctx
	}

	return context.WithValue(ctx, userContextKey, user)
}

func GetUser(ctx context.Context) *models.User {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
