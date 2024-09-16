package adminuser

import "time"

type Status string

const (
	Active   Status = "active"
	InActive Status = "inActive"
)

type AdminUser struct {
	ID           string
	UserName     string
	Email        string
	PasswordHash string
	Status       Status
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
