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
	SortOrder      null.Int  `db:"sort_order" json:"sort_order"`
	CreatedAt      time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt      null.Time `db:"deleted_at" json:"deleted_at"`
}

func NewDepartment(
	departmentName string,
	departmentCode string,
	sortOrder int,
) Department {
	sort := null.IntFromPtr(nil)
	if sortOrder != 0 {
		sort = null.IntFrom(int64(sortOrder))
	}

	return Department{
		DepartmentID:   uuid.New(),
		DepartmentName: departmentName,
		DepartmentCode: departmentCode,
		SortOrder:      sort,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      null.TimeFromPtr(nil),
	}
}
