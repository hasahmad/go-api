package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateMembersTable, downCreateMembersTable)
}

func upCreateMembersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS members (
		"member_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"member_code" varchar(150) NOT NULL,
		"first_name" varchar(150) NOT NULL,
		"middle_name" varchar(150),
		"last_name" varchar(150),
		"tanzeem" varchar(50) NOT NULL DEFAULT 'K',
		"is_waqf_nau" bool NOT NULL DEFAULT False,
		"is_musi" bool NOT NULL DEFAULT False,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"members_pkey" PRIMARY KEY ("member_id")
	)`)
	return err
}

func downCreateMembersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE members`)
	return err
}
