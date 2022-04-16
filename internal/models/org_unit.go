package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var OrgUnitLevels = []string{
	"",
}

type OrgUnit struct {
	OrgUnitID           uuid.UUID     `db:"org_unit_id" json:"org_unit_id" goqu:"defaultifempty,skipupdate"`
	OrgUnitName         string        `db:"org_unit_name" json:"org_unit_name"`
	OrgUnitCode         string        `db:"org_unit_code" json:"org_unit_code"`
	OrgUnitLevel        string        `db:"org_unit_level" json:"org_unit_level"`
	ParentOrgUnitID     uuid.NullUUID `db:"parent_org_unit_id" json:"parent_org_unit_id"`
	MergedIntoOrgUnitID string        `db:"merged_into_org_unit_id" json:"merged_into_org_unit_id"`
	SplitFromOrgUnitID  string        `db:"split_from_org_unit_id" json:"split_from_org_unit_id"`
	Timezone            string        `db:"timezone" json:"timezone"`
	SortOrder           int           `db:"sort_order" json:"sort_order"`
	CreatedAt           time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt           time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt           null.Time     `db:"deleted_at" json:"deleted_at"`
}
