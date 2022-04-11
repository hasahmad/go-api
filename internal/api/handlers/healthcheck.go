package handlers

import (
	"fmt"
	"net/http"

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

	err := helpers.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(h.Logger, w, r, err)
	}
}
