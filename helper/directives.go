package helper

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func AuthDirectives(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	token := ctx.Value("TokenValue")

	if token == nil {
		return nil, &gqlerror.Error{
			Message: "Permission denied",
		}
	}
	ctx = context.WithValue(ctx, "UserID", token)
	return next(ctx)
}
