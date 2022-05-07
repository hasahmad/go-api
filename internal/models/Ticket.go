package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Ticket struct {
	TicketID          uuid.UUID     `db:"ticket_id" json:"ticket_id" goqu:"defaultifempty,skipupdate"`
	SubmittedByUserID uuid.UUID     `db:"submitted_by_user_id" json:"submitted_by_user_id"`
	Title             string        `db:"title" json:"title"`
	TicketTypeID      uuid.NullUUID `db:"ticket_type_id" json:"ticket_type_id" goqu:"defaultifempty"`
	TicketStatusID    uuid.NullUUID `db:"ticket_status_id" json:"ticket_status_id" goqu:"defaultifempty"`
	Description       null.String   `db:"description" json:"description" goqu:"defaultifempty"`
	CreatedAt         time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt         time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt         null.Time     `db:"deleted_at" json:"deleted_at"`
}

func NewTicket(
	SubmittedByUserID uuid.UUID,
	title string,
	ticketTypeID uuid.NullUUID,
	ticketStatusID uuid.NullUUID,
) Ticket {
	return Ticket{
		TicketID:          uuid.New(),
		SubmittedByUserID: SubmittedByUserID,
		Title:             title,
		TicketTypeID:      ticketStatusID,
		TicketStatusID:    ticketStatusID,
		Description:       null.StringFrom(""),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         null.TimeFromPtr(nil),
	}
}

func TicketCols() []string {
	return []string{
		"ticket_id",
		"submitted_by_user_id",
		"title",
		"ticket_type_id",
		"ticket_status_id",
		"description",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
