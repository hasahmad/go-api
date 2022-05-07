package handlers

import (
	"net/http"

	"github.com/hasahmad/go-api/internal/helpers"
)

func (h *Handlers) GetAllMemberOrgUnitsHandler(w http.ResponseWriter, r *http.Request) {
	memberId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	result, err := h.Repositories.MemberOrgUnits.FindByMemberId(r.Context(), memberId, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) GetMemberOrgUnitHandler(w http.ResponseWriter, r *http.Request) {
	memberId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	result, err := h.Repositories.MemberOrgUnits.FindActiveByMemberId(r.Context(), memberId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}

func (h *Handlers) UpdateMemberOrgUnitHandler(w http.ResponseWriter, r *http.Request) {
	memberId, err := helpers.ReadUUIDParam(r)
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	newOrgUnitId, err := helpers.ReadUUIDParamByKey(r, "org_unit_id")
	if err != nil {
		helpers.BadRequestResponse(h.Logger, w, r, err)
		return
	}

	result, err := h.Repositories.MemberOrgUnits.UpdateMemberOrgUnit(r.Context(), memberId, newOrgUnitId)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"detail": result}, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}
