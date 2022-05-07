package handlers

import (
	"fmt"
	"net/http"

	"github.com/doug-martin/goqu/v9"
	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/hasahmad/go-api/internal/dto"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/validator"
	"gopkg.in/guregu/null.v4"
)

func (h *Handlers) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repositories.Users.FindAll(r.Context(), []goqu.Expression{}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": users}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	user, err := h.Repositories.Users.FindById(r.Context(), userId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": user}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserDto
	err := helpers.ReadJSON(w, r, &input)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	v := validator.New()
	models.ValidateEmail(v, input.Email)
	models.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		helpers.FailedValidationResponse(h.Logger, w, r, v.Errors)
		return
	}

	user, err := models.NewUser(input.Username, input.Password)
	user.Email = null.StringFrom(input.Email)
	user.FirstName = null.StringFrom(input.FirstName)
	user.LastName = null.StringFrom(input.LastName)

	u, err := h.Repositories.Users.Insert(r.Context(), *user)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": u}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	user, err := h.Repositories.Users.FindById(r.Context(), userId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	var input dto.UpdateUserDto
	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	v := validator.New()
	userInput, shouldUpdate, err := input.ToJson(v)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	if !v.Valid() {
		helpers.FailedValidationResponse(h.Logger, w, r, v.Errors)
		return
	}

	if !shouldUpdate {
		helpers.BadRequestResponse(h.Logger, w, r, fmt.Errorf("no data provided"))
		return
	}

	u, err := h.Repositories.Users.Update(r.Context(), userId, user.Version, userInput)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": u}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	_, err = h.Repositories.Users.FindById(r.Context(), userId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = h.Repositories.Users.Delete(r.Context(), userId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": "Successfully deleted"}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) GetUserRolesHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	user_roles, err := h.Repositories.UserRoles.FindAll(
		r.Context(),
		[]goqu.Expression{
			goqu.Ex{"user_id": userId},
		},
		nil,
	)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": user_roles}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) CreateUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	roleId, err := helpers.ReadUUIDParamByKey(r, "role_id")
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	user_role, err := h.Repositories.UserRoles.Insert(
		r.Context(),
		models.UserRole{
			UserID: userId,
			RoleID: roleId,
		},
	)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": user_role}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) DeleteUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	roleId, err := helpers.ReadUUIDParamByKey(r, "role_id")
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	err = h.Repositories.UserRoles.Delete(
		r.Context(),
		roleId,
		userId,
	)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": "successfully deleted"}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := apicontext.GetUser(r.Context())
	user_roles, err := h.Repositories.Roles.FindByUserId(
		r.Context(),
		user.UserID,
	)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	for i, ur := range user_roles {
		permissions, err := h.Repositories.Permissions.FindByRoleId(
			r.Context(),
			ur.RoleID,
		)

		if err != nil {
			helpers.ServerErrorResponse(h.Logger, w, r, err)
			return
		}

		user_roles[i].Permissions = permissions
	}

	user.Roles = user_roles

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": user}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}
