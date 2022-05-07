package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
	"gopkg.in/guregu/null.v4"
)

type CreateOfficeDto struct {
	OfficeName      string        `json:"office_name"`
	OfficeLevel     string        `json:"office_level"`
	Tanzeem         string        `json:"tanzeem"`
	DepartmentID    uuid.NullUUID `json:"department_id"`
	MultipleAllowed bool          `json:"multiple_allowed"`
	Reportable      bool          `json:"reportable"`
	Electable       bool          `json:"electable"`
	SortOrder       null.Int      `json:"sort_order"`
}

func (r CreateOfficeDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.OfficeName != "", "office_name", "must be provided")
	v.Check(r.OfficeLevel != "", "office_level", "must be provided")
	v.Check(r.Tanzeem != "", "tanzeem", "must be provided")

	return v
}

type UpdateOfficeDto struct {
	OfficeName      string        `json:"office_name"`
	OfficeLevel     string        `json:"office_level"`
	Tanzeem         string        `json:"tanzeem"`
	DepartmentID    uuid.NullUUID `json:"department_id"`
	MultipleAllowed *bool         `json:"multiple_allowed"`
	Reportable      *bool         `json:"reportable"`
	Electable       *bool         `json:"electable"`
	SortOrder       null.Int      `json:"sort_order"`
}

func (r UpdateOfficeDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.OfficeName != "" {
		shouldUpdate = true
		result["office_name"] = r.OfficeName
	}

	if r.OfficeLevel != "" {
		shouldUpdate = true
		result["office_level"] = r.OfficeLevel
	}

	if r.Tanzeem != "" {
		shouldUpdate = true
		result["tanzeem"] = r.Tanzeem
	}

	if r.DepartmentID.Valid {
		shouldUpdate = true
		result["department_id"] = r.DepartmentID.UUID
	}

	if r.MultipleAllowed != nil {
		shouldUpdate = true
		result["multiple_allowed"] = r.MultipleAllowed
	}

	if r.Reportable != nil {
		shouldUpdate = true
		result["reportable"] = r.Reportable
	}

	if r.Electable != nil {
		shouldUpdate = true
		result["electable"] = r.Electable
	}

	if r.SortOrder.Valid {
		shouldUpdate = true
		result["sort_order"] = r.SortOrder.Int64
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
