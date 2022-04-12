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

type UserRoleRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r UserRoleRepo) TableName() string {
	return "user_roles"
}

func (r UserRoleRepo) PrimaryKey() string {
	return ""
}

func (r UserRoleRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.UserRole, error) {
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

	var result []models.UserRole
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r UserRoleRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.UserRole, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.UserRole
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserRoleRepo) FindById(ctx context.Context, roleId uuid.UUID, userId uuid.UUID) (models.UserRole, error) {
	return r.FindOneBy(ctx, goqu.Ex{"role_id": roleId, "user_id": userId})
}

func (r UserRoleRepo) Insert(ctx context.Context, u models.UserRole) (models.UserRole, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.UserRole
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserRoleRepo) Update(ctx context.Context, key string, id uuid.UUID, data helpers.Envelope) (models.UserRole, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{key: id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.UserRole
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserRoleRepo) DeleteBy(ctx context.Context, where goqu.Ex) error {
	sel := r.sql.
		Delete(r.TableName()).
		Where(where)

	_, err := sel.Executor().QueryContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRoleRepo) Delete(ctx context.Context, roleId uuid.UUID, userId uuid.UUID) error {
	return r.DeleteBy(ctx, goqu.Ex{"role_id": roleId, "user_id": userId})
}
