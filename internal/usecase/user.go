package usecase

import (
	"context"
	"database/sql"

	"github.com/linzhengen/mii-go/internal/domain/trans"
	"github.com/linzhengen/mii-go/internal/domain/user"
	"github.com/linzhengen/mii-go/pkg/hash"
	"github.com/linzhengen/mii-go/pkg/uuid"
)

func NewUserUseCase(
	db *sql.DB,
	userRepo user.Repository,
	transRepo trans.Repository,
) UserUseCase {
	return &userUseCase{
		db:        db,
		userRepo:  userRepo,
		transRepo: transRepo,
	}
}

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*user.User, error)
	CreateUser(ctx context.Context, name, password, email string) error
}

type userUseCase struct {
	db        *sql.DB
	userRepo  user.Repository
	transRepo trans.Repository
}

func (uc userUseCase) GetUser(ctx context.Context, id string) (*user.User, error) {
	return uc.userRepo.FindOne(ctx, id)
}

func (uc userUseCase) CreateUser(ctx context.Context, name, password, email string) error {
	return uc.transRepo.ExecTrans(ctx, func(ctx context.Context) error {
		return uc.userRepo.Create(ctx, &user.User{
			ID:       uuid.MustString(),
			Name:     name,
			Password: hash.MD5String(password),
			Email:    email,
			Status:   user.Active,
		})
	})
}
