package user

import "context"

type Repository interface {
	FineOne(ctx context.Context, id string) (*User, error)
}
