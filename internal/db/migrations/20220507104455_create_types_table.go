package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTypesTable, downCreateTypesTable)
}

func upCreateTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS types (
		"type_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"value" varchar(250) NOT NULL,
		"label" varchar(250),
		"description" varchar(250),
		"parent_type" varchar(250),

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE("value", "parent_type"),
		CONSTRAINT
			"types_pkey" PRIMARY KEY ("type_id")
	)`)
	return err
}

func downCreateTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE types`)
	return err
}
