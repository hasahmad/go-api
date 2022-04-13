package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddRoleAndPermissions, downAddRoleAndPermissions)
}

func upAddRoleAndPermissions(tx *sql.Tx) error {
	rolesStr := ""
	roles := []string{
		"ADMIN", "MANAGER", "USER",
	}
	for i, v := range roles {
		comma := ","
		if i == len(roles)-1 {
			comma = ""
		}
		rolesStr = fmt.Sprintf("%s ('%s')%s", rolesStr, v, comma)
	}

	rolesQuery := fmt.Sprintf("INSERT INTO roles (role_name) VALUES %s", rolesStr)
	_, err := tx.Exec(rolesQuery)
	if err != nil {
		return err
	}

	crud := []string{
		"BROWSE",
		"READ",
		"EDIT",
		"ADD",
		"DELETE",
	}

	permissionsStr := ""
	permissions := []string{
		"ROLE", "PERMISSION", "USER",
	}
	for i, v := range permissions {
		for j, c := range crud {
			comma := ","
			if (i + j) == (len(crud) - 1 + len(permissions) - 1) {
				comma = ""
			}
			permissionsStr = fmt.Sprintf("%s ('%s-%s')%s", permissionsStr, c, v, comma)
		}
	}

	permissionsQuery := fmt.Sprintf("INSERT INTO permissions (permission_name) VALUES %s", permissionsStr)
	_, err = tx.Exec(permissionsQuery)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO role_permissions(role_id, permission_id)
	SELECT r.role_id, p.permission_id
	FROM roles r, permissions p
	WHERE r.role_name = 'ADMIN'
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddRoleAndPermissions(tx *sql.Tx) error {
	return nil
}
