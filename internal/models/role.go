package models

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Role struct {
	RoleID          uuid.UUID   `db:"role_id" json:"role_id" goqu:"defaultifempty,skipupdate"`
	RoleName        string      `db:"role_name" json:"role_name"`
	RoleDescription null.String `db:"role_description" json:"role_description" goqu:"defaultifempty"`
	// extra calculated properties
	Permissions []Permission `db:"-" json:"permissions,omitempty"`
}

func RoleCols() []string {
	return []string{
		"role_id",
		"role_name",
		"role_description",
	}
}
