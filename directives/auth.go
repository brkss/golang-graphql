package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	middleware "github.com/brkss/golang-graphql/middlewares"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver)(interface {}, error){
	fmt.Print("auth directie called !")

	payload := middleware.GetPayload(ctx)
	fmt.Printf("%+v\n", payload)
	return next(ctx)
}
