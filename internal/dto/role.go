package dto

import (
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateRoleRequest struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

type UpdateRoleRequest struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

func (u UpdateRoleRequest) ToJson(v *validator.Validator) (helpers.Envelope, bool, error) {
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
