package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUserOfficesTable, downCreateUserOfficesTable)
}

func upCreateUserOfficesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS user_office_requests (
		"user_office_request_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"office_request_id" UUID NOT NULL,
		"request_type" varchar(150) NOT NULL,
		"request_status" varchar(150) NOT NULL,
		"user_id" UUID NOT NULL,
		"office_id" UUID NOT NULL,
		"org_unit_id" UUID NOT NULL,
		"period_id" UUID NOT NULL,
		"start_date" date,
		"end_date" date,

		"request_by_user_id" UUID,
		"approved_by_user_id" UUID,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"user_office_requests_pkey" PRIMARY KEY ("user_office_request_id"),
		CONSTRAINT
			"fk_user_office_requests_office_request" FOREIGN KEY ("office_request_id")
				REFERENCES office_requests ("office_request_id"),
		CONSTRAINT
			"fk_user_office_requests_office" FOREIGN KEY ("office_id")
				REFERENCES offices ("office_id"),
		CONSTRAINT
			"fk_user_office_requests_user" FOREIGN KEY ("user_id")
				REFERENCES users ("user_id"),
		CONSTRAINT
			"fk_user_office_requests_org_unit" FOREIGN KEY ("org_unit_id")
				REFERENCES org_units ("org_unit_id"),
		CONSTRAINT
			"fk_user_office_requests_period" FOREIGN KEY ("period_id")
				REFERENCES periods ("period_id"),
		CONSTRAINT
			"fk_user_office_requests_request_by_user" FOREIGN KEY ("request_by_user_id")
				REFERENCES users ("user_id"),
		CONSTRAINT
			"fk_user_office_requests_approved_by_user_id" FOREIGN KEY ("approved_by_user_id")
				REFERENCES users ("user_id")
	)`)
	return err
}

func downCreateUserOfficesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE user_office_requests`)
	return err
}
