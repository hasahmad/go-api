package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Member struct {
	MemberID   uuid.UUID   `db:"member_id" json:"member_id" goqu:"defaultifempty,skipupdate"`
	MemberCode string      `db:"member_code" json:"member_code"`
	FirstName  string      `db:"first_name" json:"first_name"`
	MiddleName null.String `db:"middle_name" json:"middle_name"`
	LastName   null.String `db:"last_name" json:"last_name"`
	Tanzeem    string      `db:"tanzeem" json:"tanzeem"`
	IsWaqfNau  bool        `db:"is_waqf_nau" json:"is_waqf_nau"`
	IsMusi     bool        `db:"is_musi" json:"is_musi"`
	CreatedAt  time.Time   `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt  time.Time   `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt  null.Time   `db:"deleted_at" json:"deleted_at"`
	// extra calculated properties
	OrgUnitID uuid.NullUUID `db:"org_unit_id" json:"org_unit_id" goqu:"skipupdate"`
	User      *User         `db:"user" json:"user,omitempty"`
	Email     null.String   `db:"email" json:"email"`
	Emails    []string      `db:"-" json:"emails"`
	OrgUnit   *OrgUnit      `db:"-" json:"org_unit"`
	OrgUnits  []OrgUnit     `db:"-" json:"org_units,omitempty"`
}

func NewMember(
	memberCode string,
	firstName string,
	middleName null.String,
	lastName null.String,
	tanzeem string,
	isWaqfNau bool,
	isMusi bool,
) Member {
	return Member{
		MemberID:   uuid.New(),
		MemberCode: memberCode,
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Tanzeem:    tanzeem,
		IsWaqfNau:  isWaqfNau,
		IsMusi:     isMusi,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  null.TimeFromPtr(nil),
	}
}

func MemberCols() []string {
	return []string{
		"member_id",
		"member_code",
		"first_name",
		"middle_name",
		"last_name",
		"tanzeem",
		"is_waqf_nau",
		"is_musi",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func MemberColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range MemberCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
