package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)




func TestCreateRecipeCategory(t *testing.T){
	
	arg := CreateCategoryParams{
		ID: uuid.New().String(),
		Title: "Category Test",
		Image: "cat_1.png",
		Active: sql.NullBool{
			Valid: true,	
			Bool: true,
		},
	}

	cat, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cat)

}
