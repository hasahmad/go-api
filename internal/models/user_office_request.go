package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var (
	UserOfficeRequestTypes = []string{
		"ADD",      // add new member to amila
		"DISABLE",  // disable member in amila
		"ENABLE",   // enable member in amila
		"ADD_TEMP", // add new as temporarily member to amila
		"UPDATE",   // replace/update member in amila
		"DELETE",   // remove member from amila
	}
	UserOfficeRequestStatuses = []string{
		"DRAFT",
		"SUBMITTED",
		"APPROVED",
		"REJECTED",
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
	// extra calculated properties
	TotalOfficesApproved  null.Int       `db:"-" json:"total_offices_approved,omitempty"`
	TotalOfficesRequested null.Int       `db:"-" json:"total_offices_requested,omitempty"`
	OfficeRequest         *OfficeRequest `db:"-" json:"office_request,omitempty"`
	Office                *Office        `db:"-" json:"office,omitempty"`
	User                  *User          `db:"-" json:"user,omitempty"`
	OrgUnit               *OrgUnit       `db:"-" json:"org_unit,omitempty"`
	Period                *Period        `db:"-" json:"period,omitempty"`
}
