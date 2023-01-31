package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)


func RandomCategory(t *testing.T) RecipeCategory{
	arg := CreateCategoryParams{
		ID: uuid.New().String(),
		Title: fmt.Sprintf("Category Test %d", time.Now().Minute()),
		Image: "cat_default_image.png",
		Active: sql.NullBool{
			Valid: true,	
			Bool: true,
		},
	}

	cat, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cat)
	
	return cat;
}

func TestCreateRecipeCategory(t *testing.T){
	RandomCategory(t);
}

func TestGetCategories(t *testing.T){
	n := 10;

	for i := 0; i < n; i++ {
		RandomCategory(t)
	}

	cats, err := testQueries.GetCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, cats)

	for i := 0; i < n; i++ {
		require.NotEmpty(t, cats[i]);
	}

}
