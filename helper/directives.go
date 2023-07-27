package helper

import (
	"context"
	"fmt"
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

	claims, err := ParseJWT(token.(string))

	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}

	fmt.Println("testt")
	fmt.Println(claims["iss"])
	ctx = context.WithValue(ctx, "UserID", claims["iss"])

	return next(ctx)
}
