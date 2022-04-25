package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var (
	PeriodGroups = []string{
		"",
	}
	PeriodTypes = []string{
		"",
	}
)

type Period struct {
	PeriodID       uuid.UUID     `db:"period_id" json:"period_id" goqu:"defaultifempty,skipupdate"`
	ParentPeriodID uuid.NullUUID `db:"parent_period_id" json:"parent_period_id"`
	PeriodGroup    string        `db:"period_group" json:"period_group"`
	PeriodType     null.String   `db:"period_type" json:"period_type"`
	StartDate      time.Time     `db:"start_date" json:"start_date"`
	EndDate        time.Time     `db:"end_date" json:"end_date"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt      time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt      null.Time     `db:"deleted_at" json:"deleted_at"`
}

func NewPeriod(
	periodGroup string,
	periodType string,
	startDate time.Time,
	endDate time.Time,
) Period {
	return Period{
		PeriodID:    uuid.New(),
		PeriodGroup: periodGroup,
		PeriodType:  null.StringFrom(periodType),
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   null.TimeFromPtr(nil),
	}
}

func PeriodCols() []string {
	return []string{
		"period_id",
		"parent_period_id",
		"period_group",
		"period_type",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func PeriodColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range PeriodCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
