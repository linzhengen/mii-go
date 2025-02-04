// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: role.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createRole = `-- name: CreateRole :execresult
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
        now())
`

type CreateRoleParams struct {
	ID       string
	Name     string
	Apigroup string
	Resource string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createRole,
		arg.ID,
		arg.Name,
		arg.Apigroup,
		arg.Resource,
	)
}

const findRoleById = `-- name: FindRoleById :one
SELECT id, name, apigroup, resource, created, updated, deleted
FROM roles
WHERE id = ? LIMIT 1
`

func (q *Queries) FindRoleById(ctx context.Context, id string) (*Role, error) {
	row := q.db.QueryRowContext(ctx, findRoleById, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Apigroup,
		&i.Resource,
		&i.Created,
		&i.Updated,
		&i.Deleted,
	)
	return &i, err
}

const updateRole = `-- name: UpdateRole :exec
UPDATE roles
SET name     = ?,
    apiGroup = ?,
    resource = ?,
    updated  = now()
WHERE id = ?
`

type UpdateRoleParams struct {
	Name     string
	Apigroup string
	Resource string
	ID       string
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateRole,
		arg.Name,
		arg.Apigroup,
		arg.Resource,
		arg.ID,
	)
	return err
}
