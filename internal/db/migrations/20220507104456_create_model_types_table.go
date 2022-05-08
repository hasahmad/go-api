package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateModelTypesTable, downCreateModelTypesTable)
}

func upCreateModelTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS model_types (
		"model_type_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"name" varchar(250) NOT NULL,
		"label" varchar(250),
		"description" varchar(250),

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE("name"),
		CONSTRAINT
			"model_types_pkey" PRIMARY KEY ("model_type_id")
	)`)
	return err
}

func downCreateModelTypesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE model_types`)
	return err
}
