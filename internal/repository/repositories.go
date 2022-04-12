package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Users           UserRepo
	Roles           RoleRepo
	Permissions     PermissionRepo
	UserRoles       UserRoleRepo
	RolePermissions RolePermissionRepo
}

func New(db *sqlx.DB, cfg config.Config) Repositories {
	sql := goqu.New(cfg.DB.Type, db)
	return Repositories{
		Users:           UserRepo{db, sql},
		Roles:           RoleRepo{db, sql},
		Permissions:     PermissionRepo{db, sql},
		UserRoles:       UserRoleRepo{db, sql},
		RolePermissions: RolePermissionRepo{db, sql},
	}
}
