package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MemberContact struct {
	MemberContactID uuid.UUID `db:"member_contact_id" json:"member_contact_id" goqu:"defaultifempty,skipupdate"`
	MemberID        uuid.UUID `db:"member_id" json:"member_id"`
	Email           string    `db:"email" json:"email"`
	PrimaryEmail    bool      `db:"primary_email" json:"primary_email"`
	CreatedAt       time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt       null.Time `db:"deleted_at" json:"deleted_at"`
	// extra calculated properties
	Member *Member `db:"-" json:"member,omitempty"`
}

func NewMemberContact(
	memberId uuid.UUID,
	email string,
	primary bool,
) MemberContact {
	return MemberContact{
		MemberContactID: uuid.New(),
		MemberID:        memberId,
		Email:           email,
		PrimaryEmail:    primary,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       null.TimeFromPtr(nil),
	}
}
