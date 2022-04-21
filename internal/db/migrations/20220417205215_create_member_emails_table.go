package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateMemberEmailsTable, downCreateMemberEmailsTable)
}

func upCreateMemberEmailsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS member_emails (
		"member_email_id" UUID NOT NULL DEFAULT uuid_generate_v1(),
		"member_id" UUID NOT NULL,
		"email" varchar(150) NOT NULL,
		"primary_email" bool NOT NULL DEFAULT False,

		"created_at" timestamptz DEFAULT NOW(),
		"updated_at" timestamptz DEFAULT NOW(),
		"deleted_at" timestamptz,
		UNIQUE("email"),
		CONSTRAINT
			"member_emails_pkey" PRIMARY KEY ("member_email_id"),
		CONSTRAINT
			"fk_member_emails_member" FOREIGN KEY ("member_id")
				REFERENCES members ("member_id")
	)`)
	return err
}

func downCreateMemberEmailsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE member_emails`)
	return err
}
