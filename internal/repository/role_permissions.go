package repository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/filters"
	"github.com/jmoiron/sqlx"
)

type RolePermissionRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r RolePermissionRepo) TableName() string {
	return "role_permissions"
}

func (r RolePermissionRepo) PrimaryKey() string {
	return ""
}

func (r RolePermissionRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.RolePermission, error) {
	sel := r.sql.From(r.TableName())
	if wheres != nil {
		sel = sel.Where(wheres...)
	}

	if f != nil {
		if f.Sort != "" {
			if f.SortDirection() == "DESC" {
				sel = sel.Order(goqu.I(f.SortColumn()).Desc())
			} else {
				sel = sel.Order(goqu.I(f.SortColumn()).Asc())
			}
		}

		if f.Limit() > 0 && f.Page > 0 {
			sel = sel.Limit(uint(f.Limit())).
				Offset(uint(f.Offset()))
		} else if f.Limit() > 0 {
			sel = sel.Limit(uint(f.Limit()))
		}
	}

	var result []models.RolePermission
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r RolePermissionRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.RolePermission, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.RolePermission
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RolePermissionRepo) FindById(ctx context.Context, roleId uuid.UUID, permissionId uuid.UUID) (models.RolePermission, error) {
	return r.FindOneBy(ctx, goqu.Ex{"role_id": roleId, "permission_id": permissionId})
}

func (r RolePermissionRepo) Insert(ctx context.Context, u models.RolePermission) (models.RolePermission, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.RolePermission
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RolePermissionRepo) Update(ctx context.Context, key string, id uuid.UUID, data helpers.Envelope) (models.RolePermission, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{key: id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.RolePermission
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RolePermissionRepo) DeleteBy(ctx context.Context, where goqu.Ex) error {
	sel := r.sql.
		Delete(r.TableName()).
		Where(where)

	_, err := sel.Executor().QueryContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r RolePermissionRepo) Delete(ctx context.Context, roleId uuid.UUID, permissionId uuid.UUID) error {
	return r.DeleteBy(ctx, goqu.Ex{"role_id": roleId, "permission_id": permissionId})
}
