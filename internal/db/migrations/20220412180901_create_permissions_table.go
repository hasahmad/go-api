package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreatePermissionsTable, downCreatePermissionsTable)
}

func upCreatePermissionsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS permissions (
		"permission_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"permission_name" varchar(150) NOT NULL,
		"permission_description" text,
		UNIQUE ("permission_name"),
		CONSTRAINT "permissions_pkey" PRIMARY KEY ("permission_id")
	)`)
	return err
}

func downCreatePermissionsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE permissions`)
	return err
}
