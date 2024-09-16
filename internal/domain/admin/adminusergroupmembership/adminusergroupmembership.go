package adminusergroupmembership

import "time"

type AdminUserGroupMembership struct {
	ID               string
	AdminUserID      string
	AdminUserGroupID string
	UpdatedAt        time.Time
	CreatedAt        time.Time
}
