package models

import (
	"fmt"

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

func PermissionColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range PermissionCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
