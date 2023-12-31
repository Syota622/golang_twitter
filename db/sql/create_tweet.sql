-- name: CreateTweet :one
INSERT INTO tweets (user_id, text) VALUES ($1, $2) RETURNING id, user_id, text, created_at;
