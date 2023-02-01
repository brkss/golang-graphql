package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	middleware "github.com/brkss/golang-graphql/middlewares"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver)(interface {}, error){
	payload := middleware.GetPayload(ctx)
	if payload == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}	
	return next(ctx)
}
