-- name: CreateUser :one
INSERT INTO users (username,email,hashed_password)
VALUES ($1,$2,$3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
hashed_password = $3
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;