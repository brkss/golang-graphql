package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)


func Auth(ctx context.Context, obj interface{}, next graphql.Resolver)(interface {}, error){
	fmt.Print("auth directie called !")

	return next(ctx)
}



