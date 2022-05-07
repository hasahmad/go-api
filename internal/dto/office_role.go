package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateOfficeRoleDto struct {
	OfficeID uuid.UUID `json:"office_id"`
	RoleID   uuid.UUID `json:"role_id"`
}

func (r CreateOfficeRoleDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.OfficeID.String() != "", "office_id", "must be provided")
	v.Check(r.RoleID.String() != "", "role_id", "must be provided")

	return v
}

type UpdateOfficeRoleDto struct {
	OfficeID uuid.UUID `json:"office_id"`
	RoleID   uuid.UUID `json:"role_id"`
}

func (r UpdateOfficeRoleDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.OfficeID.String() != "" {
		shouldUpdate = true
		result["office_id"] = r.OfficeID.String()
	}

	if r.RoleID.String() != "" {
		shouldUpdate = true
		result["role_id"] = r.RoleID.String()
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
