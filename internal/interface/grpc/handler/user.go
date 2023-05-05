package handler

import (
	"context"

	"github.com/linzhengen/mii-go/internal/usecase"
	"github.com/linzhengen/mii-go/protobuf/go/v1/user"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func (u userHandler) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) ListUser(ctx context.Context, request *user.ListUserRequest) (*user.ListUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdatePasswordUser(ctx context.Context, request *user.UpdatePasswordUserRequest) (*user.UpdatePasswordUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(userUseCase usecase.UserUseCase) user.UserServiceServer {
	return &userHandler{userUseCase: userUseCase}
}
