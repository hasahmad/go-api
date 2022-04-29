package models

import (
	"fmt"
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
	// extra calculated properties
	UserOfficeRequest []UserOfficeRequest `json:"user_requests,omitempty"`
}

func NewOfficeRequest(
	officeID uuid.UUID,
	orgUnitID uuid.UUID,
	periodId uuid.UUID,
	startDate null.Time,
	endDate null.Time,
) OfficeRequest {
	return OfficeRequest{
		OfficeRequestID: uuid.New(),
		OfficeID:        officeID,
		OrgUnitID:       orgUnitID,
		PeriodID:        periodId,
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       null.TimeFromPtr(nil),
	}
}

func OfficeRequestCols() []string {
	return []string{
		"office_request_id",
		"office_id",
		"org_unit_id",
		"period_id",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func OfficeRequestColsMap(keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range OfficeRequestCols() {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
