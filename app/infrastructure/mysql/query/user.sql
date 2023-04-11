-- name: FindUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;