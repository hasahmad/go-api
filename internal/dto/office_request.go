package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
	"gopkg.in/guregu/null.v4"
)

type CreateOfficeRequestDto struct {
	OfficeID  uuid.UUID `json:"office_id"`
	OrgUnitID uuid.UUID `json:"org_unit_id"`
	PeriodID  uuid.UUID `json:"period_id"`
	StartDate null.Time `json:"start_date"`
	EndDate   null.Time `json:"end_date"`
}

func (r CreateOfficeRequestDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.OfficeID.String() != "", "office_id", "must be provided")
	v.Check(r.OrgUnitID.String() != "", "org_unit_id", "must be provided")
	v.Check(r.PeriodID.String() != "", "period_id", "must be provided")

	return v
}

type UpdateOfficeRequestDto struct {
	OfficeID  uuid.UUID `json:"office_id"`
	OrgUnitID uuid.UUID `json:"org_unit_id"`
	PeriodID  uuid.UUID `json:"period_id"`
	StartDate null.Time `json:"start_date"`
	EndDate   null.Time `json:"end_date"`
}

func (r UpdateOfficeRequestDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.OfficeID.String() != "" {
		shouldUpdate = true
		result["office_id"] = r.OfficeID.String()
	}

	if r.OrgUnitID.String() != "" {
		shouldUpdate = true
		result["org_unit_id"] = r.OrgUnitID.String()
	}

	if r.PeriodID.String() != "" {
		shouldUpdate = true
		result["period_id"] = r.PeriodID.String()
	}

	if r.StartDate.Valid {
		shouldUpdate = true
		result["start_date"] = r.StartDate.Time
	}

	if r.EndDate.Valid {
		shouldUpdate = true
		result["end_date"] = r.EndDate.Time
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
