package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateStatusesTable, downCreateStatusesTable)
}

func upCreateStatusesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS statuses (
		"status_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"value" varchar(250) NOT NULL,
		"label" varchar(250),
		"description" varchar(250),
		"parent_type" varchar(250),

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE("value", "parent_type"),
		CONSTRAINT
			"statuses_pkey" PRIMARY KEY ("status_id")
	)`)
	return err
}

func downCreateStatusesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE statuses`)
	return err
}
