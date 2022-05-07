package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
	"gopkg.in/guregu/null.v4"
)

type CreateTicketDto struct {
	Title        string        `json:"title"`
	TicketTypeID uuid.NullUUID `json:"ticket_type_id"`
	Description  null.String   `json:"description"`
}

func (c CreateTicketDto) Validate() bool {
	if c.Title == "" {
		return false
	}

	return true
}

type UpdateTicketDto struct {
	Title          string        `json:"title"`
	TicketTypeID   uuid.NullUUID `json:"ticket_type_id"`
	TicketStatusID uuid.NullUUID `json:"ticket_status_id"`
	Description    null.String   `json:"description"`
}

func (u UpdateTicketDto) ToJson(v *validator.Validator) (helpers.Envelope, bool, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if u.Title != "" {
		shouldUpdate = true
		result["title"] = u.Title
	}

	if u.Description.Valid {
		shouldUpdate = true
		result["description"] = u.Description
	}

	if u.TicketTypeID.UUID.String() != "" {
		shouldUpdate = true
		result["ticket_type_id"] = u.TicketTypeID
	}

	if u.TicketStatusID.UUID.String() != "" {
		shouldUpdate = true
		result["ticket_status_id"] = u.TicketStatusID
	}

	return result, shouldUpdate, nil
}
