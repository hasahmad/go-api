package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateDepartmentDto struct {
	DepartmentName        string `json:"department_name"`
	DepartmentDescription string `json:"department_description"`
	DepartmentCode        string `json:"department_code"`
	SortOrder             int    `json:"sort_order"`
}

func (r CreateDepartmentDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.DepartmentName != "", "department_name", "must be provided")
	v.Check(r.DepartmentCode != "", "department_code", "must be provided")

	return v
}

type UpdateDepartmentDto struct {
	DepartmentName        string `json:"department_name"`
	DepartmentDescription string `json:"department_description"`
	DepartmentCode        string `json:"department_code"`
	SortOrder             int    `json:"sort_order"`
}

func (r UpdateDepartmentDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.DepartmentName != "" {
		shouldUpdate = true
		result["department_name"] = r.DepartmentName
	}

	if r.DepartmentDescription != "" {
		shouldUpdate = true
		result["department_description"] = r.DepartmentDescription
	}

	if r.DepartmentCode != "" {
		shouldUpdate = true
		result["department_code"] = r.DepartmentCode
	}

	if r.SortOrder != 0 {
		shouldUpdate = true
		result["sort_order"] = r.SortOrder
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
