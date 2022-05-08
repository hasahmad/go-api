package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ContentType struct {
	ContentTypeID uuid.UUID   `db:"content_type_id" json:"content_type_id" goqu:"defaultifempty,skipupdate"`
	Value         string      `db:"value" json:"value"`
	Label         null.String `db:"label" json:"label" goqu:"defaultifempty"`
	Description   null.String `db:"description" json:"description" goqu:"defaultifempty"`
	ParentType    null.String `db:"parent_type" json:"parent_type" goqu:"defaultifempty"`
	CreatedAt     time.Time   `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt     time.Time   `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt     null.Time   `db:"deleted_at" json:"deleted_at"`
}

func NewContentType(
	value string,
	label string,
	parent string,
	description string,
) ContentType {
	return ContentType{
		ContentTypeID: uuid.New(),
		Value:         value,
		Label:         null.StringFrom(label),
		ParentType:    null.StringFrom(parent),
		Description:   null.StringFrom(description),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     null.TimeFromPtr(nil),
	}
}

func ContentTypeCols() []string {
	return []string{
		"content_type_id",
		"value",
		"label",
		"description",
		"parent_type",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
