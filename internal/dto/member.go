package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
	"gopkg.in/guregu/null.v4"
)

type CreateMemberDto struct {
	MemberCode string      `json:"member_code"`
	FirstName  string      `json:"first_name"`
	MiddleName null.String `json:"middle_name"`
	LastName   null.String `json:"last_name"`
	Tanzeem    string      `json:"tanzeem"`
	IsWaqfNau  bool        `json:"is_waqf_nau"`
	IsMusi     bool        `json:"is_musi"`
}

func (r CreateMemberDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.MemberCode != "", "member_code", "must be provided")
	v.Check(r.FirstName != "", "first_name", "must be provided")

	return v
}

type UpdateMemberDto struct {
	MemberCode string      `json:"member_code"`
	FirstName  string      `json:"first_name"`
	MiddleName null.String `json:"middle_name"`
	LastName   null.String `json:"last_name"`
	Tanzeem    string      `json:"tanzeem"`
	IsWaqfNau  *bool       `json:"is_waqf_nau"`
	IsMusi     *bool       `json:"is_musi"`
}

func (r UpdateMemberDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.MemberCode != "" {
		shouldUpdate = true
		result["member_code"] = r.MemberCode
	}

	if r.FirstName != "" {
		shouldUpdate = true
		result["first_name"] = r.FirstName
	}

	if r.MiddleName.Valid {
		shouldUpdate = true
		result["middle_name"] = r.MiddleName.String
	}

	if r.LastName.Valid {
		shouldUpdate = true
		result["last_name"] = r.LastName.String
	}

	if r.IsWaqfNau != nil {
		shouldUpdate = true
		result["is_waqf_nau"] = r.IsWaqfNau
	}

	if r.IsMusi != nil {
		shouldUpdate = true
		result["is_musi"] = r.IsMusi
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
