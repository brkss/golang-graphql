// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (RecipeCategory, error)
	GetCategories(ctx context.Context) (RecipeCategory, error)
}

var _ Querier = (*Queries)(nil)
