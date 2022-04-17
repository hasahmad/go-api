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

type DepartmentsRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r DepartmentsRepo) TableName() string {
	return "departments"
}

func (r DepartmentsRepo) PrimaryKey() string {
	return "department_id"
}

func (r DepartmentsRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.Department, error) {
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

	var result []models.Department
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r DepartmentsRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.Department, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.Department
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r DepartmentsRepo) FindById(ctx context.Context, id uuid.UUID) (models.Department, error) {
	return r.FindOneBy(ctx, goqu.Ex{r.PrimaryKey(): id})
}

func (r DepartmentsRepo) Insert(ctx context.Context, u models.Department) (models.Department, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.Department
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r DepartmentsRepo) Update(ctx context.Context, id uuid.UUID, data helpers.Envelope) (models.Department, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.Department
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r DepartmentsRepo) Delete(ctx context.Context, id uuid.UUID) error {
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
