-- name: CreateUser :one
INSERT INTO users(
    username, 
    password,
    role
) VALUES (
    $1, $2, $3
) RETURNING *;