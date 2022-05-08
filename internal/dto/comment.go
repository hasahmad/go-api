package dto

import (
	"time"

	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/pkg/validator"
)

type CreateCommentDto struct {
	CommentTypeID     string `json:"comment_type_id"`
	ParentCommentID   string `json:"parent_comment_id"`
	CommentText       string `json:"comment_text"`
	Description       string `json:"description"`
	ModelTypeID       string `json:"model_type_id"`
	ModelTypeRecordID string `json:"model_type_record_id"`
}

func (r CreateCommentDto) Validate(v *validator.Validator) *validator.Validator {
	if v == nil {
		v = validator.New()
	}

	v.Check(r.CommentText != "", "comment_text", "must be provided")
	v.Check(r.ModelTypeID != "", "model_type_id", "must be provided")
	v.Check(r.ModelTypeRecordID != "", "model_type_record_id", "must be provided")

	return v
}

type UpdateCommentDto struct {
	CommentTypeID     string `json:"comment_type_id"`
	ParentCommentID   string `json:"parent_comment_id"`
	CommentText       string `json:"comment_text"`
	Description       string `json:"description"`
	ModelTypeID       string `json:"model_type_id"`
	ModelTypeRecordID string `json:"model_type_record_id"`
}

func (r UpdateCommentDto) ToJson(v *validator.Validator) (helpers.Envelope, error) {
	shouldUpdate := false
	result := helpers.Envelope{
		"updated_at": time.Now(),
	}

	if r.CommentTypeID != "" {
		shouldUpdate = true
		result["comment_type_id"] = r.CommentTypeID
	}

	if r.ParentCommentID != "" {
		shouldUpdate = true
		result["parent_comment_id"] = r.ParentCommentID
	}

	if r.CommentText != "" {
		shouldUpdate = true
		result["comment_text"] = r.CommentText
	}

	if r.Description != "" {
		shouldUpdate = true
		result["description"] = r.Description
	}

	if r.ModelTypeID != "" {
		shouldUpdate = true
		result["model_type_id"] = r.ModelTypeID
	}

	if r.ModelTypeRecordID != "" {
		shouldUpdate = true
		result["model_type_record_id"] = r.ModelTypeRecordID
	}

	if !shouldUpdate {
		v.AddError("input", "no data provided")
	}

	return result, nil
}
