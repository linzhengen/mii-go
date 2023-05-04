-- name: FindRoleById :one
SELECT *
FROM roles
WHERE id = ? LIMIT 1;

-- name: UpdateRole :exec
UPDATE roles
SET name     = ?,
    apiGroup = ?,
    resource = ?,
    updated  = now()
WHERE id = ?;

-- name: CreateRole :execresult
INSERT INTO roles (id,
                   name,
                   apiGroup,
                   resource,
                   created,
                   updated)
VALUES (?,
        ?,
        ?,
        ?,
        now(),
        now());

