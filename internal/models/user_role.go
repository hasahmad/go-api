package models

import (
	"github.com/google/uuid"
)

type UserRole struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	RoleID uuid.UUID `db:"role_id" json:"role_id"`
}

func UserRoleCols() []string {
	return []string{
		"user_id",
		"role_id",
	}
}
