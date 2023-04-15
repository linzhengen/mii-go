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
    updated  = now()
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (id,
                   name,
                   password,
                   email,
                   status,
                   created,
                   updated)
VALUES (?,
        ?,
        ?,
        ?,
        ?,
        now(),
        now());
