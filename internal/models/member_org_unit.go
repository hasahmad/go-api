package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MemberOrgUnit struct {
	MemberOrgUnitID uuid.UUID `db:"member_org_unit_id" json:"member_org_unit_id" goqu:"defaultifempty,skipupdate"`
	MemberID        uuid.UUID `db:"member_id" json:"member_id"`
	OrgUnitID       uuid.UUID `db:"org_unit_id" json:"org_unit_id"`
	PrimaryOrgUnit  bool      `db:"primary_org_unit" json:"primary_org_unit"`
	CreatedAt       time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt       null.Time `db:"deleted_at" json:"deleted_at"`
	// extra calculated properties
	Member  *Member  `db:"-" json:"member,omitempty"`
	OrgUnit *OrgUnit `db:"-" json:"org_unit,omitempty"`
}

func NewMemberOrgUnit(
	memberId uuid.UUID,
	orgUnitId uuid.UUID,
	primaryOrgUnit bool,
) MemberOrgUnit {
	return MemberOrgUnit{
		MemberOrgUnitID: uuid.New(),
		MemberID:        memberId,
		OrgUnitID:       orgUnitId,
		PrimaryOrgUnit:  primaryOrgUnit,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       null.TimeFromPtr(nil),
	}
}

func MemberOrgUnitCols() []string {
	return []string{
		"member_org_unit_id",
		"member_id",
		"org_unit_id",
		"primary_org_unit",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func MemberOrgUnitColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range MemberOrgUnitCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
