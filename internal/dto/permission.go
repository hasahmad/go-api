package dto

import (
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreatePermissionDto struct {
	PermissionName        string `json:"permission_name"`
	PermissionDescription string `json:"permission_description"`
}

type UpdatePermissionDto struct {
	PermissionName        string `json:"permission_name"`
	PermissionDescription string `json:"permission_description"`
}

func (u UpdatePermissionDto) ToJson(v *validator.Validator) (helpers.Envelope, bool, error) {
	shouldUpdate := false
	result := helpers.Envelope{}

	if u.PermissionName != "" {
		shouldUpdate = true
		result["permission_name"] = u.PermissionName
	}

	if u.PermissionDescription != "" {
		shouldUpdate = true
		result["permission_description"] = u.PermissionDescription
	}

	return result, shouldUpdate, nil
}
