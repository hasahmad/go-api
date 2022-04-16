package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Office struct {
	OfficeID        uuid.UUID     `db:"office_id" json:"office_id" goqu:"defaultifempty,skipupdate"`
	OfficeName      string        `db:"office_name" json:"office_name"`
	OfficeLevel     string        `db:"office_level" json:"office_level"`
	Tanzeem         string        `db:"tanzeem" json:"tanzeem"`
	DepartmentID    uuid.NullUUID `db:"department_id" json:"department_id"`
	MultipleAllowed bool          `db:"multiple_allowed" json:"multiple_allowed"`
	Reportable      bool          `db:"reportable" json:"reportable"`
	Electable       bool          `db:"electable" json:"electable"`
	SortOrder       int           `db:"sort_order" json:"sort_order"`
	CreatedAt       time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt       time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt       null.Time     `db:"deleted_at" json:"deleted_at"`
	// extra calculated properties
	Roles []Role `db:"-" json:"roles,omitempty"`
}
