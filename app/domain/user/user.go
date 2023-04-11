package user

import "time"

type User struct {
	ID       string
	Name     string
	Password string
	Email    string
	Status   string
	Created  time.Time
	Updated  *time.Time
	Deleted  *time.Time
}
