package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var OrgUnitLevels = Levels

type OrgUnit struct {
	OrgUnitID           uuid.UUID     `db:"org_unit_id" json:"org_unit_id" goqu:"defaultifempty,skipupdate"`
	OrgUnitName         string        `db:"org_unit_name" json:"org_unit_name"`
	OrgUnitCode         string        `db:"org_unit_code" json:"org_unit_code"`
	OrgUnitLevel        string        `db:"org_unit_level" json:"org_unit_level"`
	ParentOrgUnitID     uuid.NullUUID `db:"parent_org_unit_id" json:"parent_org_unit_id"`
	MergedIntoOrgUnitID string        `db:"merged_into_org_unit_id" json:"merged_into_org_unit_id"`
	SplitFromOrgUnitID  string        `db:"split_from_org_unit_id" json:"split_from_org_unit_id"`
	Timezone            string        `db:"timezone" json:"timezone"`
	SortOrder           null.Int      `db:"sort_order" json:"sort_order"`
	CreatedAt           time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt           time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt           null.Time     `db:"deleted_at" json:"deleted_at"`
}

func NewOrgUnit(
	orgUnitName string,
	orgUnitCode string,
	orgUnitLevel string,
	sortOrder int,
) OrgUnit {
	sort := null.IntFromPtr(nil)
	if sortOrder != 0 {
		sort = null.IntFrom(int64(sortOrder))
	}

	return OrgUnit{
		OrgUnitID:    uuid.New(),
		OrgUnitName:  orgUnitName,
		OrgUnitCode:  orgUnitCode,
		OrgUnitLevel: orgUnitLevel,
		SortOrder:    sort,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    null.TimeFromPtr(nil),
	}
}

func OrgUnitCols() []string {
	return []string{
		"org_unit_id",
		"org_unit_name",
		"org_unit_code",
		"org_unit_level",
		"parent_org_unit_id",
		"merged_into_org_unit_id",
		"split_from_org_unit_id",
		"timezone",
		"sort_order",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func OrgUnitColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range OrgUnitCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
