package handler

import (
	"context"

	"github.com/linzhengen/mii-go/internal/domain/user"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/linzhengen/mii-go/internal/usecase"
	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func (u userHandler) GetUser(ctx context.Context, request *v1user.GetUserRequest) (*v1user.GetUserResponse, error) {
	r, err := u.userUseCase.GetUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	return &v1user.GetUserResponse{
		User: &v1user.User{
			Id:      r.ID,
			Name:    r.Name,
			Email:   r.Email,
			Status:  toPbUserStatus(r.Status),
			Created: timestamppb.New(r.Created),
			Updated: timestamppb.New(*r.Updated),
		},
	}, nil
}

func (u userHandler) ListUser(ctx context.Context, request *v1user.ListUserRequest) (*v1user.ListUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) CreateUser(ctx context.Context, request *v1user.CreateUserRequest) (*v1user.CreateUserResponse, error) {
	if err := u.userUseCase.CreateUser(ctx, request.User.Name, request.User.Password, request.User.Email); err != nil {
		return nil, err
	}
	r, err := u.GetUser(ctx, &v1user.GetUserRequest{UserId: request.User.Id})
	if err != nil {
		return nil, err
	}
	return &v1user.CreateUserResponse{User: r.User}, nil
}

func (u userHandler) UpdateUser(ctx context.Context, request *v1user.UpdateUserRequest) (*v1user.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdatePasswordUser(ctx context.Context, request *v1user.UpdatePasswordUserRequest) (*v1user.UpdatePasswordUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func toPbUserStatus(s user.Status) v1user.Status {
	switch s {
	case user.Active:
		return v1user.Status_STATUS_ACTIVE
	case user.InActive:
		return v1user.Status_STATUS_INACTIVE
	default:
		return v1user.Status_STATUS_UNSPECIFIED
	}
}

func NewUserHandler(userUseCase usecase.UserUseCase) v1user.UserServiceServer {
	return &userHandler{userUseCase: userUseCase}
}
