package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type OfficeRequest struct {
	OfficeRequestID uuid.UUID `db:"office_request_id" json:"office_request_id" goqu:"defaultifempty,skipupdate"`
	OfficeID        uuid.UUID `db:"office_id" json:"office_id"`
	OrgUnitID       uuid.UUID `db:"org_unit_id" json:"org_unit_id"`
	PeriodID        uuid.UUID `db:"period_id" json:"period_id"`
	StartDate       null.Time `db:"start_date" json:"start_date"`
	EndDate         null.Time `db:"end_date" json:"end_date"`
	CreatedAt       time.Time `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt       null.Time `db:"deleted_at" json:"deleted_at"`
}