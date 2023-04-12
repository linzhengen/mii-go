package usecase

import (
	"context"
	"database/sql"
	"github.com/linzhengen/mii-go/app/domain/user"
)

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*user.User, error)
}

type userUseCase struct {
	db      *sql.DB
	useRepo user.Repository
}

func (u userUseCase) GetUser(ctx context.Context, id string) (*user.User, error) {
	return u.useRepo.FineOne(ctx, id)
}
