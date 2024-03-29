package models

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Role struct {
	RoleID          uuid.UUID    `db:"role_id" json:"role_id" goqu:"defaultifempty,skipupdate"`
	RoleName        string       `db:"role_name" json:"role_name"`
	RoleDescription null.String  `db:"role_description" json:"role_description" goqu:"defaultifempty"`
	Permissions     []Permission `db:"-" json:"permissions,omitempty"`
}
