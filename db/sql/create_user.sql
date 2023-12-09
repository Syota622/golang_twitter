-- name: CreateUser :exec
INSERT INTO users (email, password_hash) VALUES ($1, $2);
