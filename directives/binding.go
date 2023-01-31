package directives

import (
	"context"
	"fmt"
	//"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)


var validate *validator.Validate;


func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string)(interface{}, error) {
	fmt.Println("constraint : ", constraint)
	//val, _ := next(ctx);
	/*
	if err != nil {
		panic(err)
	}

	fieldName := *graphql.GetPathContext(ctx).Field

	err = validate.Var(val, constraint)
	if err != nil {
		error := fmt.Errorf("%s%+v", fieldName, err.Error())
		return val,error 
	}
	*/
	return next(ctx)

}
