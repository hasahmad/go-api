package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOfficeRequestsTable, downCreateOfficeRequestsTable)
}

func upCreateOfficeRequestsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS office_requests (
		"office_request_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"office_id" UUID NOT NULL,
		"org_unit_id" UUID NOT NULL,
		"period_id" UUID NOT NULL,
		"start_date" date,
		"end_date" date,

		"created_by_user_id" UUID,
		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"office_requests_pkey" PRIMARY KEY ("office_request_id"),
		CONSTRAINT
			"fk_office_requests_office" FOREIGN KEY ("office_id")
				REFERENCES offices ("office_id"),
		CONSTRAINT
			"fk_office_requests_org_unit" FOREIGN KEY ("org_unit_id")
				REFERENCES org_units ("org_unit_id"),
		CONSTRAINT
			"fk_office_requests_period" FOREIGN KEY ("period_id")
				REFERENCES periods ("period_id"),
		CONSTRAINT
			"fk_office_requests_created_by_user_id" FOREIGN KEY ("created_by_user_id")
				REFERENCES users ("user_id")
	)`)
	return err
}

func downCreateOfficeRequestsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE office_requests`)
	return err
}
