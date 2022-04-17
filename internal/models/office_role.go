package models

import (
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
