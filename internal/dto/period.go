package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreatePeriodRequest struct {
	ParentPeriodID string `json:"parent_period_id"`
	PeriodGroup    string `json:"period_group"`
	PeriodType     string `json:"period_type"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
}

func (r CreatePeriodRequest) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.StartDate != "", "period_group", "must be provided")
	v.Check(r.StartDate != "", "start_date", "must be provided")
	v.Check(r.EndDate != "", "end_date", "must be provided")

	return v
}

type UpdatePeriodRequest struct {
	ParentPeriodID string     `json:"parent_period_id"`
	PeriodGroup    string     `json:"period_group"`
	PeriodType     string     `json:"period_type"`
	StartDate      *time.Time `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
}

func (r UpdatePeriodRequest) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.ParentPeriodID != "" {
		shouldUpdate = true
		result["parent_period_id"] = r.ParentPeriodID
	}

	if r.PeriodGroup != "" {
		shouldUpdate = true
		result["period_group"] = r.PeriodGroup
	}

	if r.PeriodType != "" {
		shouldUpdate = true
		result["period_type"] = r.PeriodType
	}

	if r.StartDate != nil {
		shouldUpdate = true
		result["start_date"] = r.StartDate
	}

	if r.EndDate != nil {
		shouldUpdate = true
		result["end_date"] = r.EndDate
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
