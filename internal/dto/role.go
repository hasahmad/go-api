package dto

import (
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateRoleDto struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

type UpdateRoleDto struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

func (u UpdateRoleDto) ToJson(v *validator.Validator) (helpers.Envelope, bool, error) {
	shouldUpdate := false
	result := helpers.Envelope{}

	if u.RoleName != "" {
		shouldUpdate = true
		result["role_name"] = u.RoleName
	}

	if u.RoleDescription != "" {
		shouldUpdate = true
		result["role_description"] = u.RoleDescription
	}

	return result, shouldUpdate, nil
}
