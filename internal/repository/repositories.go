package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Users              UserRepo
	Roles              RoleRepo
	Permissions        PermissionRepo
	UserRoles          UserRoleRepo
	RolePermissions    RolePermissionRepo
	Departments        DepartmentsRepo
	Offices            OfficesRepo
	OfficeRequests     OfficeRequestsRepo
	OfficeRoles        OfficeRolesRepo
	OrgUnits           OrgUnitsRepo
	Periods            PeriodsRepo
	UserOfficeRequests UserOfficeRequestsRepo
}

func New(db *sqlx.DB, cfg config.Config) Repositories {
	sql := goqu.New(cfg.DB.Type, db)
	return Repositories{
		Users:              UserRepo{db, sql},
		Roles:              RoleRepo{db, sql},
		Permissions:        PermissionRepo{db, sql},
		UserRoles:          UserRoleRepo{db, sql},
		RolePermissions:    RolePermissionRepo{db, sql},
		Departments:        DepartmentsRepo{db, sql},
		Offices:            OfficesRepo{db, sql},
		OfficeRequests:     OfficeRequestsRepo{db, sql},
		OfficeRoles:        OfficeRolesRepo{db, sql},
		OrgUnits:           OrgUnitsRepo{db, sql},
		Periods:            PeriodsRepo{db, sql},
		UserOfficeRequests: UserOfficeRequestsRepo{db, sql},
	}
}
