package directives

import (
	"context"
	"fmt"
	//"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)


var validate = validator.New() 

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string)(interface{}, error) {
	//fmt.Println("constraint : ", constraint)
	val, err := next(ctx);
	if err != nil {
		panic(err)
	}

	fieldName := *graphql.GetPathContext(ctx).Field

	fmt.Println("val : ", val)
	fmt.Println("constraint : ", constraint)
	err = validate.Var(val, constraint)
	if err != nil {
		e := fmt.Errorf("%s : invalid entry", fieldName)
		return val, e 
	}
	return next(ctx) 
}
