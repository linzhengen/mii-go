package user

import "time"

type Status string

const (
	Active   Status = "active"
	InActive Status = "inActive"
)

func (s Status) IsAllowedValue() bool {
	switch s {
	case Active, InActive:
		return true
	}
	return false
}

type User struct {
	ID       string
	Name     string
	Password string
	Email    string
	Status   Status
	Created  time.Time
	Updated  *time.Time
	Deleted  *time.Time
}
