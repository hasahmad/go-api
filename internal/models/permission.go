package models

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Permission struct {
	PermissionID          uuid.UUID   `db:"permission_id" json:"permission_id" goqu:"defaultifempty,skipupdate"`
	PermissionName        string      `db:"permission_name" json:"permission_name"`
	PermissionDescription null.String `db:"permission_description" json:"permission_description" goqu:"defaultifempty"`
}

func PermissionCols() []string {
	return []string{
		"permission_id",
		"permission_name",
		"permission_description",
	}
}
