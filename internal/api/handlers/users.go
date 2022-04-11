package handlers

import (
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/helpers"
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
