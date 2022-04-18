package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddMemberIdColInUsersTable, downAddMemberIdColInUsersTable)
}

func upAddMemberIdColInUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	ALTER TABLE users
	ADD COLUMN "member_id" UUID REFERENCES members ("member_id")
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX member_id_index on users (member_id)")
	return err
}

func downAddMemberIdColInUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
	ALTER TABLE users
	DROP COLUMN IF EXISTS member_id
	`)
	return err
}
