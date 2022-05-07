package models

import (
	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID `db:"role_id" json:"role_id"`
	PermissionID uuid.UUID `db:"permission_id" json:"permission_id"`
}

func RolePermissionCols() []string {
	return []string{
		"role_id",
		"permission_id",
	}
}
