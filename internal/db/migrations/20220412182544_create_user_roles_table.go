package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUserRolesTable, downCreateUserRolesTable)
}

func upCreateUserRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS user_roles (
		"user_id" UUID NOT NULL,
		"role_id" UUID NOT NULL,
		CONSTRAINT "user_roles_pkey" PRIMARY KEY ("user_id", "role_id")
	)`)
	return err
}

func downCreateUserRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE user_roles`)
	return err
}
