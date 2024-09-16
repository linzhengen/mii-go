package adminusergroupmembership

import "context"

type Repository interface {
	FineOne(ctx context.Context, id string) (*AdminUserGroupMembership, error)
	Create(ctx context.Context, u *AdminUserGroupMembership) error
	Update(ctx context.Context, u *AdminUserGroupMembership) error
}
