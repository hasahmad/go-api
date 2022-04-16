package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOfficesTable, downCreateOfficesTable)
}

func upCreateOfficesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS offices (
		"office_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"office_name" varchar(250) NOT NULL,
		"department_id" UUID,
		"office_level" varchar(50) NOT NULL,
		"tanzeem" varchar(50) NOT NULL,
		"multiple_allowed" bool NOT NULL DEFAULT False,
		"reportable" bool NOT NULL DEFAULT False,
		"electable" bool NOT NULL DEFAULT False,
		"sort_order" int,
		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"offices_pkey" PRIMARY KEY ("office_id"),
		CONSTRAINT
			"fk_office_department" FOREIGN KEY ("department_id")
				REFERENCES departments ("department_id")
	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX department_id_index on offices (department_id)")
	if err != nil {
		return err
	}

	return err
}

func downCreateOfficesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE offices`)
	return err
}
