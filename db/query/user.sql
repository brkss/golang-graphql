-- name: RegisterUser :one
INSERT INTO users
( id, name, email, username, password )
VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetUserByIdent :one
SELECT * FROM users
WHERE 
email = $1 OR 
username = $1
LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;
