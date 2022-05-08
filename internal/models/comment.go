package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Comment struct {
	CommentID         uuid.UUID     `db:"comment_id" json:"comment_id" goqu:"defaultifempty,skipupdate"`
	UserID            uuid.UUID     `db:"user_id" json:"user_id"`
	CommentTypeID     uuid.NullUUID `db:"comment_type_id" json:"comment_type_id"`
	ParentCommentID   uuid.NullUUID `db:"parent_comment_id" json:"parent_comment_id"`
	CommentText       string        `db:"comment_text" json:"comment_text"`
	Description       null.String   `db:"description" json:"description" goqu:"defaultifempty"`
	ModelTypeID       uuid.UUID     `db:"model_type_id" json:"model_type_id"`
	ModelTypeRecordID uuid.UUID     `db:"model_type_record_id" json:"model_type_record_id"`
	CreatedAt         time.Time     `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt         time.Time     `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt         null.Time     `db:"deleted_at" json:"deleted_at"`
}

func NewComment(
	commentText string,
	userID uuid.UUID,
	modelTypeID uuid.UUID,
	modelTypeRecordID uuid.UUID,
	description string,
) Comment {
	return Comment{
		CommentID:         uuid.New(),
		UserID:            userID,
		CommentText:       commentText,
		ModelTypeID:       modelTypeID,
		ModelTypeRecordID: modelTypeRecordID,
		Description:       null.StringFrom(description),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         null.TimeFromPtr(nil),
	}
}

func CommentCols() []string {
	return []string{
		"comment_id",
		"user_id",
		"comment_type_id",
		"parent_comment_id",
		"comment_text",
		"description",
		"model_type_id",
		"model_type_record_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
