package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	db "github.com/brkss/golang-graphql/db/sqlc"
	"github.com/brkss/golang-graphql/graph/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterUserInput) (*model.RegisterUserResponse, error) {
	
	err := validate.Var(input.Email, "require,email")
	if err != nil {
		return nil, err
	}

	arg := db.RegisterUserParams{
		ID: uuid.New().String(),
		Name: input.Name,
		Email: input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	user, err := r.Store.RegisterUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	response := model.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
	} 	
	return &model.RegisterUserResponse{
		Status: true,
		User: &response,
	}, nil
}
