package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateMemberOrgUnitsTable, downCreateMemberOrgUnitsTable)
}

func upCreateMemberOrgUnitsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS member_org_units (
		"member_org_unit_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"member_id" UUID NOT NULL,
		"org_unit_id" UUID NOT NULL,
		"primary_org_unit" bool NOT NULL DEFAULT False,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"member_org_units_pkey" PRIMARY KEY ("member_org_unit_id"),
		CONSTRAINT
			"fk_member_org_units_member" FOREIGN KEY ("member_id")
				REFERENCES members ("member_id"),
		CONSTRAINT
			"fk_member_org_units_org_unit" FOREIGN KEY ("org_unit_id")
				REFERENCES org_units ("org_unit_id")
	)`)
	return err
}

func downCreateMemberOrgUnitsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE member_org_units`)
	return err
}
