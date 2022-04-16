package models

import (
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
	PeriodType     string        `db:"period_type" json:"period_type"`
	StartDate      time.Time     `db:"start_date" json:"start_date"`
	EndDate        time.Time     `db:"end_date" json:"end_date"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt      time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt      null.Time     `db:"deleted_at" json:"deleted_at"`
}
