package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Department struct {
	DepartmentID   uuid.UUID `db:"department_id" json:"department_id" goqu:"defaultifempty,skipupdate"`
	DepartmentName string    `db:"department_name" json:"department_name"`
	DepartmentCode string    `db:"department_code" json:"department_code"`
	SortOrder      int       `db:"sort_order" json:"sort_order"`
	CreatedAt      time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt      null.Time `db:"deleted_at" json:"deleted_at"`
}
