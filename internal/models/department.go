package models

import (
	"fmt"
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

func DepartmentCols() []string {
	return []string{
		"department_id",
		"department_name",
		"department_code",
		"sort_order",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func DepartmentColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range DepartmentCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
