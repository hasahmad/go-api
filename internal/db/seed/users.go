package seed

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/models"
	"gopkg.in/guregu/null.v4"
)

func (s Seed) UserSeed() {
	tx, err := s.sql.Begin()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		u, err := models.NewUser(
			fmt.Sprintf("user%d", i),
			"password",
		)
		if err != nil {
			panic(err)
		}

		u.FirstName = null.StringFrom(faker.FirstName())
		u.LastName = null.StringFrom(faker.LastName())
		if i == 0 {
			u.Username = "admin"
			u.Email = null.StringFrom("admin@mail.com")
			u.IsStaff = true
			u.IsSuperuser = true
		} else {
			u.Email = null.StringFrom(faker.Email())
		}

		_, err = tx.Insert("users").Rows(u).Executor().Exec()
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				panic(rollbackErr)
			}
			panic(err)
		}
	}

	var userRole models.Role
	foundUserRole, err := tx.From("roles").Where(goqu.Ex{"role_name": "USER"}).Limit(1).ScanStruct(&userRole)
	if err != nil {
		panic(err)
	}
	if foundUserRole && userRole.RoleID.String() != "" {
		_, err = tx.Query(fmt.Sprintf(`
			INSERT INTO user_roles (user_id, role_id)
			SELECT u.user_id, '%s'
			FROM users u
		`, userRole.RoleID.String()))
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				panic(rollbackErr)
			}

			panic(err)
		}
	}

	var adminRole models.Role
	foundAdmin, err := tx.From("roles").Where(goqu.Ex{"role_name": "ADMIN"}).Limit(1).ScanStruct(&adminRole)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			panic(rollbackErr)
		}
		panic(err)
	}
	if foundAdmin && adminRole.RoleID.String() != "" {
		_, err = tx.Query(fmt.Sprintf(`
			INSERT INTO user_roles (user_id, role_id)
			SELECT u.user_id, '%s'
			FROM users u
			WHERE u.username = 'admin'
		`, adminRole.RoleID.String()))
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				panic(rollbackErr)
			}
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
