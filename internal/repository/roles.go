package repository

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/filters"
	"github.com/jmoiron/sqlx"
)

type RoleRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r RoleRepo) TableName() string {
	return "roles"
}

func (r RoleRepo) PrimaryKey() string {
	return "role_id"
}

func (r RoleRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.Role, error) {
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

	var result []models.Role
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r RoleRepo) FindByUserId(ctx context.Context, userId uuid.UUID) ([]models.Role, error) {
	sel := r.sql.Select(goqu.I("r.*")).From(goqu.T(r.TableName()).As("r")).
		Join(
			goqu.T("user_roles").As("ur"),
			goqu.On(goqu.Ex{"ur.role_id": goqu.I("r.role_id")}),
		).
		Join(
			goqu.T("office_roles").As("or"),
			goqu.On(goqu.Ex{"or.role_id": goqu.I("r.role_id")}),
		).
		Join(
			goqu.T("user_office_requests").As("uor"),
			goqu.On(goqu.Ex{
				"uor.office_id": goqu.I("or.office_id"),
				"uor.user_id":   goqu.I("ur.user_id"),
			}),
		).
		Where(goqu.Ex{
			"ur.user_id":         userId,
			"uor.user_id":        userId,
			"or.deleted_at":      nil,
			"ur.deleted_at":      nil,
			"uor.deleted_at":     nil,
			"uor.is_default":     true,
			"uor.request_status": models.UserOfficeRequestStatuses["APPROVED"],
		}, goqu.Or(
			goqu.C("uo.end_date").IsNull(),
			goqu.C("uo.end_date").Gte(time.Now()),
		), goqu.Or(
			goqu.C("uor.request_type").IsNull(),
			goqu.C("uor.request_type").Eq(""),
			goqu.C("uor.request_type").In([]string{
				models.UserOfficeRequestTypes["ADD"],
				models.UserOfficeRequestTypes["ENABLE"],
				models.UserOfficeRequestTypes["ADD_TEMP"],
				models.UserOfficeRequestTypes["UPDATE"],
			}),
		)).
		Order(goqu.I("r.role_name").Desc())

	var result []models.Role
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r RoleRepo) FindByUserIdAndRoleNames(ctx context.Context, userId uuid.UUID, names []string) ([]models.Role, error) {
	sel := r.sql.Select(goqu.I("r.*")).From(goqu.T(r.TableName()).As("r")).
		Join(
			goqu.T("user_roles").As("ur"),
			goqu.On(goqu.Ex{"ur.role_id": goqu.I("r.role_id")}),
		).
		Where(goqu.Ex{
			"ur.user_id":  userId,
			"r.role_name": names,
		}).
		Order(goqu.I("r.role_name").Desc())

	var result []models.Role
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r RoleRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.Role, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.Role
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RoleRepo) FindById(ctx context.Context, id uuid.UUID) (models.Role, error) {
	return r.FindOneBy(ctx, goqu.Ex{r.PrimaryKey(): id})
}

func (r RoleRepo) Insert(ctx context.Context, u models.Role) (models.Role, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.Role
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RoleRepo) Update(ctx context.Context, id uuid.UUID, data helpers.Envelope) (models.Role, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.Role
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r RoleRepo) Delete(ctx context.Context, id uuid.UUID) error {
	sel := r.sql.
		Delete(r.TableName()).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Limit(1)

	_, err := sel.Executor().QueryContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
