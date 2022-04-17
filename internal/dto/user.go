package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Version   int    `json:"-"`
}

func (c CreateUserRequest) Validate() bool {
	if c.Username == "" {
		return false
	}
	if c.Email == "" {
		return false
	}
	if c.Password == "" {
		return false
	}

	return true
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Version   int    `json:"-"`
}

func (u UpdateUserRequest) ToJson(v *validator.Validator) (helpers.Envelope, bool, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
		"version":    u.Version + 1,
	}

	if u.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
		if err != nil {
			return result, shouldUpdate, err
		}

		if v != nil {
			models.ValidatePasswordPlaintext(v, u.Password)
		}

		shouldUpdate = true
		result["password"] = string(passwordHash)
	}

	if u.FirstName != "" {
		shouldUpdate = true
		result["first_name"] = u.FirstName
	}

	if u.LastName != "" {
		shouldUpdate = true
		result["last_name"] = u.LastName
	}

	if u.Email != "" {
		shouldUpdate = true
		if v != nil {
			models.ValidateEmail(v, u.Email)
		}
		result["email"] = u.Email
	}

	return result, shouldUpdate, nil
}
