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

type MembersRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r MembersRepo) TableName() string {
	return "members"
}

func (r MembersRepo) PrimaryKey() string {
	return "member_id"
}

func (r MembersRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.Member, error) {
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

	var result []models.Member
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r MembersRepo) BaseSelectBy(where goqu.Ex) *goqu.SelectDataset {
	return r.sql.
		Select(goqu.I("m.*"), goqu.I("mo.org_unit_id"), goqu.I("me.email")).
		From(goqu.T(r.TableName()).As("m")).
		LeftJoin(
			goqu.T("member_org_units").As("mo"),
			goqu.On(goqu.Ex{
				"mo.member_id":        goqu.I("m.member_id"),
				"mo.primary_org_unit": true,
				"mo.deleted_at":       nil,
			}),
		).
		LeftJoin(
			goqu.T("member_emails").As("me"),
			goqu.On(goqu.Ex{
				"me.member_id":     goqu.I("m.member_id"),
				"me.primary_email": true,
				"me.deleted_at":    nil,
			}),
		).
		Where(goqu.Ex{"m.deleted_at": nil}).
		Where(where)
}

func (r MembersRepo) FindBy(ctx context.Context, where goqu.Ex) ([]models.Member, error) {
	sel := r.BaseSelectBy(where).GroupBy(goqu.I("m.member_id"))
	var result []models.Member
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r MembersRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.Member, error) {
	sel := r.BaseSelectBy(where).GroupBy(goqu.I("m.member_id"))
	var result models.Member
	found, err := sel.ScanStructContext(ctx, &result)

	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MembersRepo) FindById(ctx context.Context, id uuid.UUID) (models.Member, error) {
	return r.FindOneBy(ctx, goqu.Ex{r.PrimaryKey(): id})
}

func (r MembersRepo) FindByOrgUnitId(ctx context.Context, id uuid.UUID) ([]models.Member, error) {
	return r.FindBy(ctx, goqu.Ex{"mo.org_unit_id": id})
}

func (r MembersRepo) FindByEmail(ctx context.Context, email string) ([]models.Member, error) {
	return r.FindBy(ctx, goqu.Ex{"me.email": email})
}

func (r MembersRepo) Insert(ctx context.Context, u models.Member) (models.Member, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.Member
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MembersRepo) Update(ctx context.Context, id uuid.UUID, data helpers.Envelope) (models.Member, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.Member
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MembersRepo) Delete(ctx context.Context, id uuid.UUID) error {
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
