package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MemberEmail struct {
	MemberEmailID uuid.UUID `db:"member_email_id" json:"member_email_id" goqu:"defaultifempty,skipupdate"`
	MemberID      uuid.UUID `db:"member_id" json:"member_id"`
	Email         string    `db:"email" json:"email"`
	PrimaryEmail  bool      `db:"primary_email" json:"primary_email"`
	CreatedAt     time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt     null.Time `db:"deleted_at" json:"deleted_at"`
	// extra calculated properties
	Member *Member `db:"-" json:"member,omitempty"`
}

func NewMemberEmail(
	memberId uuid.UUID,
	email string,
	primary bool,
) MemberEmail {
	return MemberEmail{
		MemberEmailID: uuid.New(),
		MemberID:      memberId,
		Email:         email,
		PrimaryEmail:  primary,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     null.TimeFromPtr(nil),
	}
}

func MemberEmailCols() []string {
	return []string{
		"member_email_id",
		"member_id",
		"email",
		"primary_email",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
