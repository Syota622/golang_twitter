-- name: CreateUser :exec
INSERT INTO users (email, password_hash, is_active, activation_token) VALUES ($1, $2, $3, $4);
