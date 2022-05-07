package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var (
	UserOfficeRequestTypes = map[string]string{
		"ADD":      "ADD",      // add new member to amila
		"DISABLE":  "DISABLE",  // disable member in amila
		"ENABLE":   "ENABLE",   // enable member in amila
		"ADD_TEMP": "ADD_TEMP", // add new as temporarily member to amila
		"UPDATE":   "UPDATE",   // replace/update member in amila
		"DELETE":   "DELETE",   // remove member from amila
	}
	UserOfficeRequestStatuses = map[string]string{
		"DRAFT":     "DRAFT",
		"SUBMITTED": "SUBMITTED",
		"APPROVED":  "APPROVED",
		"REJECTED":  "REJECTED",
	}
)

type UserOfficeRequest struct {
	UserOfficeRequestID uuid.UUID     `db:"user_office_request_id" json:"user_office_request_id" goqu:"defaultifempty,skipupdate"`
	OfficeRequestID     uuid.UUID     `db:"office_request_id" json:"office_request_id"`
	RequestType         string        `db:"request_type" json:"request_type"`
	RequestStatus       string        `db:"request_status" json:"request_status"`
	UserID              uuid.UUID     `db:"user_id" json:"user_id"`
	OfficeID            uuid.UUID     `db:"office_id" json:"office_id"`
	OrgUnitID           uuid.UUID     `db:"org_unit_id" json:"org_unit_id"`
	PeriodID            uuid.UUID     `db:"period_id" json:"period_id"`
	StartDate           null.Time     `db:"start_date" json:"start_date"`
	EndDate             null.Time     `db:"end_date" json:"end_date"`
	RequestByUserID     uuid.NullUUID `db:"request_by_user_id" json:"request_by_user_id"`
	ApprovedByUserID    uuid.NullUUID `db:"approved_by_user_id" json:"approved_by_user_id"`
	CreatedAt           time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt           time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt           null.Time     `db:"deleted_at" json:"deleted_at"`
	IsDefault           bool          `db:"is_default" json:"is_default"`
	// extra calculated properties
	TotalOfficesApproved  *int           `json:"total_offices_approved,omitempty"`
	TotalOfficesRequested *int           `json:"total_offices_requested,omitempty"`
	OfficeRequest         *OfficeRequest `json:"office_request,omitempty"`
	Office                *Office        `json:"office,omitempty"`
	User                  *User          `json:"user,omitempty"`
	OrgUnit               *OrgUnit       `json:"org_unit,omitempty"`
	Period                *Period        `json:"period,omitempty"`
}

func NewUserOfficeRequest(
	userID uuid.UUID,
	officeRequestID uuid.UUID,
	officeID uuid.UUID,
	orgUnitID uuid.UUID,
	periodId uuid.UUID,
	startDate null.Time,
	endDate null.Time,
	requestType string,
) UserOfficeRequest {
	return UserOfficeRequest{
		UserOfficeRequestID: uuid.New(),
		OfficeRequestID:     officeRequestID,
		UserID:              userID,
		OfficeID:            officeID,
		OrgUnitID:           orgUnitID,
		PeriodID:            periodId,
		StartDate:           startDate,
		EndDate:             endDate,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		DeletedAt:           null.TimeFromPtr(nil),
		RequestType:         requestType,
		RequestStatus:       UserOfficeRequestStatuses["DRAFT"],
	}
}

func UserOfficeRequestCols() []string {
	return []string{
		"user_office_request_id",
		"office_request_id",
		"request_type",
		"request_status",
		"user_id",
		"office_id",
		"org_unit_id",
		"period_id",
		"start_date",
		"end_date",
		"request_by_user_id",
		"approved_by_user_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
