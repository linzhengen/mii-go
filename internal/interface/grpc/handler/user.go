package handler

import (
	"context"

	"github.com/linzhengen/mii-go/internal/usecase"
	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func (u userHandler) GetUser(ctx context.Context, request *v1user.GetUserRequest) (*v1user.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) ListUser(ctx context.Context, request *v1user.ListUserRequest) (*v1user.ListUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) CreateUser(ctx context.Context, request *v1user.CreateUserRequest) (*v1user.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdateUser(ctx context.Context, request *v1user.UpdateUserRequest) (*v1user.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdatePasswordUser(ctx context.Context, request *v1user.UpdatePasswordUserRequest) (*v1user.UpdatePasswordUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(userUseCase usecase.UserUseCase) v1user.UserServiceServer {
	return &userHandler{userUseCase: userUseCase}
}
