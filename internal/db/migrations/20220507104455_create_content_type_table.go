package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateContentTypesTable, downCreateContentTypesTable)
}

func upCreateContentTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS content_types (
		"content_type_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"value" varchar(250) NOT NULL,
		"label" varchar(250),
		"description" varchar(250),
		"parent_content_type" varchar(250),

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE("value", "parent_content_type"),
		CONSTRAINT
			"content_types_pkey" PRIMARY KEY ("content_type_id")
	)`)
	return err
}

func downCreateContentTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE content_types`)
	return err
}
