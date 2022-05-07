package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTicketsTable, downCreateTicketsTable)
}

func upCreateTicketsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS tickets (
		"ticket_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"submitted_by_user_id" UUID NOT NULL,
		"title" varchar(250) NOT NULL,
		"ticket_type_id" UUID NOT NULL,
		"ticket_status_id" UUID NOT NULL,
		"description" text,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		CONSTRAINT
			"tickets_pkey" PRIMARY KEY ("ticket_id"),
		CONSTRAINT
			"fk_tickets_ticket_type" FOREIGN KEY ("ticket_type_id")
				REFERENCES types ("type_id"),
		CONSTRAINT
			"fk_tickets_ticket_status" FOREIGN KEY ("ticket_status_id")
				REFERENCES statuses ("status_id")
	)`)
	return err
}

func downCreateTicketsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE tickets`)
	return err
}
