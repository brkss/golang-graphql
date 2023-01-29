

-- name: CreateCategory :one
INSERT INTO recipe_categories 
( id, title, image, active ) 
VALUES ( $1, $2, $3, $4 )
RETURNING *;


-- name: GetCategories :one
SELECT * FROM recipe_categories
ORDER BY id;
