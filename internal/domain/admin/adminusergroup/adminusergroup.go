package adminusergroup

import "time"

type Status string

const (
	Active   Status = "Active"
	Inactive Status = "Inactive"
)

type AdminUserGroup struct {
	ID          string
	GroupName   string
	Status      Status
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
