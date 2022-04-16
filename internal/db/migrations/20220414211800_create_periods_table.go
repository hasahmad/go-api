package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreatePeriodsTable, downCreatePeriodsTable)
}

func upCreatePeriodsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS periods (
		"period_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"parent_period_id" UUID,
		"period_group" varchar(50) NOT NULL,
		"period_type" varchar(250),
		"start_date" date NOT NULL,
		"end_date" date NOT NULL,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"periods_pkey" PRIMARY KEY ("period_id"),
		CONSTRAINT
			"fk_parent_period_id" FOREIGN KEY ("parent_period_id")
				REFERENCES periods ("period_id")
	)`)
	return err
}

func downCreatePeriodsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE periods`)
	return err
}
