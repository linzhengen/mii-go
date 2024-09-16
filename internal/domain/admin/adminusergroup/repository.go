package adminusergroup

import "context"

type Repository interface {
	FineOne(ctx context.Context, id string) (*AdminUserGroup, error)
	Create(ctx context.Context, u *AdminUserGroup) error
	Update(ctx context.Context, u *AdminUserGroup) error
}
