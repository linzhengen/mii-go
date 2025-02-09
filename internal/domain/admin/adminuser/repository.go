package adminuser

import "context"

type Repository interface {
	FindOne(ctx context.Context, id string) (*AdminUser, error)
	Create(ctx context.Context, u *AdminUser) error
	Update(ctx context.Context, u *AdminUser) error
}
