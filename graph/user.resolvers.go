package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	db "github.com/brkss/golang-graphql/db/sqlc"
	"github.com/brkss/golang-graphql/graph/model"
	"github.com/brkss/golang-graphql/utils"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterUserInput) (*model.AuthResponse, error) {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	arg := db.RegisterUserParams{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Email:    input.Email,
		Username: input.Username,
		Password: hash,
	}

	// user needed while generating token !
	_, err = r.Store.RegisterUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	token := "token placeholder !"
	return &model.AuthResponse{
		Status: true,
		Token: &token,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.LoginUserInput) (*model.AuthResponse, error) {

	user, err := r.Store.GetUserByIdent(ctx, input.Ident)
	if err != nil {
		msg := "User not found !"		
		return &model.AuthResponse{
			Status: false,
			Message: &msg,
		}, nil
	}

	err = utils.VerifyPassword(user.Password, input.Password)
	if err != nil {
		msg := "Invalid Password !"
		return &model.AuthResponse{
			Status: false,
			Message: &msg,
		}, nil 
	}

	// generate token !
	token := "token placeholder !"
	return &model.AuthResponse{
		Status: true,
		Token: &token,
	}, nil

}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var validate *validator.Validate
