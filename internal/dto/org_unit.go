package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateOrgUnitRequest struct {
	OrgUnitName         string `json:"org_unit_name"`
	OrgUnitCode         string `json:"org_unit_code"`
	OrgUnitLevel        string `json:"org_unit_level"`
	ParentOrgUnitID     string `json:"parent_org_unit_id"`
	MergedIntoOrgUnitID string `json:"merged_into_org_unit_id"`
	SplitFromOrgUnitID  string `json:"split_from_org_unit_id"`
	Timezone            string `json:"timezone"`
	SortOrder           int    `json:"sort_order"`
}

func (r CreateOrgUnitRequest) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.OrgUnitName != "", "org_unit_name", "must be provided")
	v.Check(r.OrgUnitCode != "", "org_unit_code", "must be provided")
	v.Check(r.OrgUnitLevel != "", "org_unit_level", "must be provided")

	return v
}

type UpdateOrgUnitRequest struct {
	OrgUnitName         string `json:"org_unit_name"`
	OrgUnitCode         string `json:"org_unit_code"`
	OrgUnitLevel        string `json:"org_unit_level"`
	ParentOrgUnitID     string `json:"parent_org_unit_id"`
	MergedIntoOrgUnitID string `json:"merged_into_org_unit_id"`
	SplitFromOrgUnitID  string `json:"split_from_org_unit_id"`
	Timezone            string `json:"timezone"`
	SortOrder           int    `json:"sort_order"`
}

func (r UpdateOrgUnitRequest) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.OrgUnitName != "" {
		shouldUpdate = true
		result["org_unit_name"] = r.OrgUnitName
	}

	if r.OrgUnitCode != "" {
		shouldUpdate = true
		result["org_unit_code"] = r.OrgUnitCode
	}

	if r.OrgUnitLevel != "" {
		shouldUpdate = true
		result["org_unit_level"] = r.OrgUnitLevel
	}

	if r.ParentOrgUnitID != "" {
		shouldUpdate = true
		result["parent_org_unit_id"] = r.ParentOrgUnitID
	}

	if r.MergedIntoOrgUnitID != "" {
		shouldUpdate = true
		result["merged_into_org_unit_id"] = r.MergedIntoOrgUnitID
	}

	if r.SplitFromOrgUnitID != "" {
		shouldUpdate = true
		result["split_from_org_unit_id"] = r.SplitFromOrgUnitID
	}

	if r.Timezone != "" {
		shouldUpdate = true
		result["timezone"] = r.Timezone
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
