package handlers

import (
	"fmt"
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/helpers"
)

func (h *Handlers) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := helpers.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": ".env",
			"port":        fmt.Sprintf("%d", h.Config.Server.Port),
		},
	}

	user_offices, err := h.Repositories.Users.FindUserOfficesBy(
		r.Context(),
		[]goqu.Expression{
			goqu.Ex{"u.user_id": "158fa6d1-fd2e-4362-81c9-c75ebfe8936e"},
		},
	)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
		return
	}

	data["user_offices"] = user_offices

	err = helpers.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}
