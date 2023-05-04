-- name: FindRoleByUserId :many
SELECT *
FROM userRoles
WHERE userId = ?;


-- name: UpdateUserRole :exec
UPDATE userRoles
SET userId  = ?,
    roleId  = ?,
    updated = now()
WHERE id = ?;

-- name: CreateUserRole :execresult
INSERT INTO userRoles (id,
                       userId,
                       roleId,
                       created,
                       updated)
VALUES (?,
        ?,
        ?,
        now(),
        now());
