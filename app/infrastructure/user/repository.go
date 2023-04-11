package user

import (
	"context"
	"github.com/linzhengen/mii-go/app/domain/user"
	"github.com/linzhengen/mii-go/app/infrastructure/mysql/sqlc"
)

type repositoryImpl struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) user.Repository {
	return &repositoryImpl{q: q}
}

func (r repositoryImpl) FineOne(ctx context.Context, id string) (*user.User, error) {
	u, err := r.q.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Status:   u.Status,
		Created:  u.Created,
		Updated:  u.Updated,
		Deleted:  u.Deleted,
	}, nil
}
