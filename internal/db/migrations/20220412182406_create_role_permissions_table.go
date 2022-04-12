package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateRolePermissionsTable, downCreateRolePermissionsTable)
}

func upCreateRolePermissionsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS role_permissions (
		"role_id" UUID NOT NULL,
		"permission_id" UUID NOT NULL,
		CONSTRAINT "role_permissions_pkey" PRIMARY KEY ("role_id", "permission_id")
	)`)
	return err
}

func downCreateRolePermissionsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE role_permissions`)
	return err
}
