package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateCommentsTable, downCreateCommentsTable)
}

func upCreateCommentsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
		"comment_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"user_id" UUID NOT NULL,
		"comment_text" text not null,
		"comment_type_id" UUID REFERENCES content_types ("content_type_id"),
		"parent_comment_id" UUID REFERENCES comments ("comment_id"),
		"description" varchar(250),

		"model_type_id" UUID NOT NULL REFERENCES model_types ("model_type_id"),
		"model_type_record_id" UUID NOT NULL,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"comments_pkey" PRIMARY KEY ("comment_id")
	)`)
	return err
}

func downCreateCommentsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE comments`)
	return err
}
