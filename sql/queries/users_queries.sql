-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (name, lastname, username, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING *;
