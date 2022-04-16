package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOrgUnitsTable, downCreateOrgUnitsTable)
}

func upCreateOrgUnitsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS org_units (
		"org_unit_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"org_unit_name" varchar(250) UNIQUE NOT NULL,
		"org_unit_code" varchar(150) UNIQUE NOT NULL,
		"org_unit_level" varchar(150) NOT NULL,
		"parent_org_unit_id" UUID,
		"merged_into_org_unit_id" UUID,
		"split_from_org_unit_id" UUID,
		"timezone" varchar(150),
		"sort_order" int,
		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"org_units_pkey" PRIMARY KEY ("org_unit_id"),
		CONSTRAINT
			"fk_org_units_parent_org_unit" FOREIGN KEY(parent_org_unit_id)
				REFERENCES org_units(org_unit_id),
		CONSTRAINT
			"fk_org_units_merged_into_org_unit" FOREIGN KEY (merged_into_org_unit_id)
				REFERENCES org_units (org_unit_id),
		CONSTRAINT
			"fk_org_units_split_from_org_unit" FOREIGN KEY (split_from_org_unit_id)
				REFERENCES org_units (org_unit_id)
	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX parent_org_unit_id_index on org_units (parent_org_unit_id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX merged_into_org_unit_id_index on org_units (merged_into_org_unit_id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX split_from_org_unit_id_index on org_units (split_from_org_unit_id)")
	if err != nil {
		return err
	}

	return err
}

func downCreateOrgUnitsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE org_units`)
	return err
}
