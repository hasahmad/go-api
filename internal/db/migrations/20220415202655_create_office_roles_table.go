package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOfficeRolesTable, downCreateOfficeRolesTable)
}

func upCreateOfficeRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS office_roles (
		"office_role_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"office_id" UUID NOT NULL,
		"role_id" UUID NOT NULL,
		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"office_roles_pkey" PRIMARY KEY ("office_role_id"),
		CONSTRAINT
			"fk_office_roles_office" FOREIGN KEY ("office_id")
				REFERENCES offices ("office_id"),
		CONSTRAINT
			"fk_office_roles_role" FOREIGN KEY ("role_id")
				REFERENCES roles ("role_id")
	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX office_id_index on office_roles (office_id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX role_id_index on office_roles (role_id)")
	if err != nil {
		return err
	}

	return err
}

func downCreateOfficeRolesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE office_roles`)
	return err
}
