package user

import (
	"context"

	"github.com/linzhengen/mii-go/internal/domain/user"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql/sqlc"
)

type repositoryImpl struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) user.Repository {
	return &repositoryImpl{q: q}
}

func (r repositoryImpl) FindOne(ctx context.Context, id string) (*user.User, error) {
	u, err := mysql.GetQ(ctx, r.q).FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Status:   user.Status(u.Status),
		Created:  u.Created,
		Updated:  u.Updated,
		Deleted:  u.Deleted,
	}, nil
}

func (r repositoryImpl) Create(ctx context.Context, u *user.User) error {
	_, err := mysql.GetQ(ctx, r.q).CreateUser(ctx, sqlc.CreateUserParams{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Status:   string(u.Status),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r repositoryImpl) Update(ctx context.Context, u *user.User) error {
	return mysql.GetQ(ctx, r.q).UpdateUser(ctx, sqlc.UpdateUserParams{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Status:   string(u.Status),
		ID:       u.ID,
	})
}
