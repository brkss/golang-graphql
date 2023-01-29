package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/brkss/golang-graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateClient - createClient"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todo := model.Todo{
		ID:   "id-1",
		Text: "finish graphql example !",
		Done: false,
		User: &model.User{
			ID:   "user-1",
			Name: "Brahim!",
		},
	}
	response := []*model.Todo{
		&todo,
	}
	return response, nil
	//panic(fmt.Errorf("not implemented: Todos - todos"))
}

// GetClients is the resolver for the getClients field.
func (r *queryResolver) GetClients(ctx context.Context) ([]*model.Client, error) {
	panic(fmt.Errorf("not implemented: GetClients - getClients"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }