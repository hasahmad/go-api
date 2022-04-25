package models

import (
	"fmt"

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

func RoleColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range RoleCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
