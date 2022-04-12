package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateRolesTable, downCreateRolesTable)
}

func upCreateRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS roles (
		"role_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"role_name" varchar(150) NOT NULL,
		"role_description" TEXT,
		UNIQUE("role_name"),
		CONSTRAINT "roles_pkey" PRIMARY KEY ("role_id")
	)`)
	return err
}

func downCreateRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE roles`)
	return err
}
