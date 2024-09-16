-- name: FindAdminUserById :one
SELECT *
FROM adminUsers
WHERE id = ? LIMIT 1;

-- name: UpdateAdminUser :exec
UPDATE adminUsers
SET userName     = ?,
    email        = ?,
    passwordHash = ?,
    status       = ?,
    WHERE id = ?;

-- name: CreateAdminUser :execresult
INSERT INTO adminUsers (id,
                        userName,
                        email,
                        passwordHash,
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
