package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateDepartmentsTable, downCreateDepartmentsTable)
}

func upCreateDepartmentsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS departments (
		"department_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"department_name" varchar(250) NOT NULL,
		"department_code" varchar(150) NOT NULL,
		"sort_order" int,
		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE ("department_name"),
		UNIQUE ("department_code"),
		CONSTRAINT "departments_pkey" PRIMARY KEY ("department_id")
	)`)
	return err
}

func downCreateDepartmentsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE departments`)
	return err
}
