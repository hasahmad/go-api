package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ModelType struct {
	ModelTypeID uuid.UUID   `db:"model_type_id" json:"model_type_id" goqu:"defaultifempty,skipupdate"`
	Name        string      `db:"name" json:"name"`
	Label       null.String `db:"label" json:"label" goqu:"defaultifempty"`
	Description null.String `db:"description" json:"description" goqu:"defaultifempty"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt   null.Time   `db:"deleted_at" json:"deleted_at"`
}

func NewModelType(
	name string,
	label string,
	description string,
) ModelType {
	return ModelType{
		ModelTypeID: uuid.New(),
		Name:        name,
		Label:       null.StringFrom(label),
		Description: null.StringFrom(description),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   null.TimeFromPtr(nil),
	}
}

func ModelTypeCols() []string {
	return []string{
		"model_type_id",
		"name",
		"label",
		"description",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
