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

type MemberOrgUnitsRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r MemberOrgUnitsRepo) TableName() string {
	return "member_org_units"
}

func (r MemberOrgUnitsRepo) PrimaryKey() string {
	return "member_org_unit_id"
}

func (r MemberOrgUnitsRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.MemberOrgUnit, error) {
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

	var result []models.MemberOrgUnit
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r MemberOrgUnitsRepo) FindByMemberId(ctx context.Context, memberId uuid.UUID, f *filters.Filters) ([]models.MemberOrgUnit, error) {
	where := []goqu.Expression{
		goqu.Ex{"member_id": memberId},
	}
	return r.FindAll(ctx, where, f)
}

func (r MemberOrgUnitsRepo) FindActiveByMemberId(ctx context.Context, memberId uuid.UUID) (models.MemberOrgUnit, error) {
	where := []goqu.Expression{
		goqu.Ex{
			"member_id":  memberId,
			"deleted_at": nil,
			"is_primary": true,
		},
	}
	var result models.MemberOrgUnit
	org_units, err := r.FindAll(ctx, where, nil)
	if err != nil {
		return result, err
	}

	if len(org_units) == 0 {
		return result, ErrNotFound
	}

	result = org_units[0]
	return result, nil
}

func (r MemberOrgUnitsRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.MemberOrgUnit, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.MemberOrgUnit
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MemberOrgUnitsRepo) FindById(ctx context.Context, id uuid.UUID) (models.MemberOrgUnit, error) {
	return r.FindOneBy(ctx, goqu.Ex{r.PrimaryKey(): id})
}

func (r MemberOrgUnitsRepo) Insert(ctx context.Context, u models.MemberOrgUnit) (models.MemberOrgUnit, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.MemberOrgUnit
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MemberOrgUnitsRepo) Update(ctx context.Context, id uuid.UUID, data helpers.Envelope) (models.MemberOrgUnit, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.MemberOrgUnit
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r MemberOrgUnitsRepo) Delete(ctx context.Context, id uuid.UUID) error {
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

func (r MemberOrgUnitsRepo) UpdateMemberOrgUnit(ctx context.Context, memberId uuid.UUID, newOrgUnitId uuid.UUID) (models.MemberOrgUnit, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(map[string]interface{}{"is_primary": false, "deleted_at": time.Now()}).
		Where(goqu.Ex{
			"member_id":  memberId,
			"deleted_at": nil,
			"is_primary": true,
		})

	var result models.MemberOrgUnit
	_, err := sel.Executor().ExecContext(ctx)
	if err != nil {
		return result, err
	}

	u := models.NewMemberOrgUnit(memberId, newOrgUnitId, true)
	result, err = r.Insert(ctx, u)
	if err != nil {
		return result, err
	}

	return result, nil
}
