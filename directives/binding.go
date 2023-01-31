package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)


var validate = validator.New() 

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string)(interface{}, error) {
	val, err := next(ctx);
	if err != nil {
		panic(err)
	}

	fieldName := *graphql.GetPathContext(ctx).Field

	fmt.Println("val : ", val)
	fmt.Println("constraint : ", constraint)
	err = validate.Var(val, constraint)
	if err != nil {
		err = fmt.Errorf("%s : invalid entry\n%+v", fieldName, err.Error())
		return val, err 
	}
	return next(ctx) 
}
