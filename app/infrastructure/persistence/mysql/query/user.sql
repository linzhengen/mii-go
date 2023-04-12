-- name: FindUserById :one
SELECT *
FROM users
WHERE id = ? LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET name     = ?,
    password = ?,
    email    = ?,
    status   = ?,
    updated  = ?
WHERE id = ?;