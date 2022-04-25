package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type OfficeRole struct {
	OfficeRoleID uuid.UUID `db:"office_role_id" json:"office_role_id" goqu:"defaultifempty,skipupdate"`
	OfficeID     uuid.UUID `db:"office_id" json:"office_id"`
	RoleID       uuid.UUID `db:"role_id" json:"role_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt    null.Time `db:"deleted_at" json:"deleted_at"`
}

func NewOfficeRole(
	officeID uuid.UUID,
	roleID uuid.UUID,
) OfficeRole {
	return OfficeRole{
		OfficeRoleID: uuid.New(),
		OfficeID:     officeID,
		RoleID:       roleID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    null.TimeFromPtr(nil),
	}
}

func OfficeRoleCols() []string {
	return []string{
		"office_role_id",
		"office_id",
		"role_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func OfficeRoleColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range OfficeRoleCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
