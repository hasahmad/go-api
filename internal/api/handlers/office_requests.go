package handlers

import (
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/dto"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/validator"
)

func (h *Handlers) GetAllOfficeRequestsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := h.Repositories.OfficeRequests.FindAll(r.Context(), []goqu.Expression{}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) GetOfficeRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	result, err := h.Repositories.OfficeRequests.FindById(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) CreateOfficeRequestHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateOfficeRequestRequest
	err := helpers.ReadJSON(w, r, &input)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	v := input.Validate(validator.New())
	if !v.Valid() {
		helpers.FailedValidationResponse(h.Logger, w, r, v.Errors)
		return
	}

	record := models.NewOfficeRequest(
		input.OfficeID,
		input.OrgUnitID,
		input.PeriodID,
		input.StartDate,
		input.EndDate,
	)

	result, err := h.Repositories.OfficeRequests.Insert(r.Context(), record)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) UpdateOfficeRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	_, err = h.Repositories.OfficeRequests.FindById(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	var input dto.UpdateOfficeRequestRequest
	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	v := validator.New()
	updateInput, err := input.ToJson(v)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	if !v.Valid() {
		helpers.FailedValidationResponse(h.Logger, w, r, v.Errors)
		return
	}

	result, err := h.Repositories.OfficeRequests.Update(r.Context(), id, updateInput)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) DeleteOfficeRequestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	_, err = h.Repositories.OfficeRequests.FindById(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = h.Repositories.OfficeRequests.Delete(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": "Successfully deleted"}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}
