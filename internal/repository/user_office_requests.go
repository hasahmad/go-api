package repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/filters"
	"github.com/jmoiron/sqlx"
)

type UserOfficeRequestsRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func (r UserOfficeRequestsRepo) TableName() string {
	return "user_office_requests"
}

func (r UserOfficeRequestsRepo) PrimaryKey() string {
	return "user_office_request_id"
}

func (r UserOfficeRequestsRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.UserOfficeRequest, error) {
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

	var result []models.UserOfficeRequest
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r UserOfficeRequestsRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.UserOfficeRequest, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where).
		Limit(1)

	var result models.UserOfficeRequest
	found, err := sel.ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserOfficeRequestsRepo) FindById(ctx context.Context, id uuid.UUID) (models.UserOfficeRequest, error) {
	return r.FindOneBy(ctx, goqu.Ex{r.PrimaryKey(): id})
}

func (r UserOfficeRequestsRepo) FindBy(ctx context.Context, where goqu.Ex) ([]models.UserOfficeRequest, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(where)

	var result []models.UserOfficeRequest
	err := sel.ScanStructsContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r UserOfficeRequestsRepo) FindByOfficeRequestId(ctx context.Context, id uuid.UUID) ([]models.UserOfficeRequest, error) {
	return r.FindBy(ctx, goqu.Ex{
		"office_request_id": id,
		"deleted_at":        nil,
	})
}

func (r UserOfficeRequestsRepo) Insert(ctx context.Context, u models.UserOfficeRequest) (models.UserOfficeRequest, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var result models.UserOfficeRequest
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserOfficeRequestsRepo) Update(ctx context.Context, id uuid.UUID, data helpers.Envelope) (models.UserOfficeRequest, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{r.PrimaryKey(): id}).
		Returning(goqu.T(r.TableName()).All())

	var result models.UserOfficeRequest
	found, err := sel.Executor().ScanStructContext(ctx, &result)
	if err != nil {
		return result, err
	}
	if !found {
		return result, ErrNotFound
	}

	return result, nil
}

func (r UserOfficeRequestsRepo) Delete(ctx context.Context, id uuid.UUID) error {
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

func (r UserOfficeRequestsRepo) GetFullUserOfficeRequest(ctx context.Context, id uuid.UUID) (models.UserOfficeRequest, error) {
	user_cols_map := models.ColsMap(models.UserCols(), "user__", "", "u.", "")
	office_cols_map := models.ColsMap(models.OfficeCols(), "office__", "", "of.", "")
	office_req_cols_map := models.ColsMap(models.OfficeRequestCols(), "office_request__", "", "or.", "")
	user_office_req_cols_map := models.ColsMap(models.UserOfficeRequestCols(), "uor__", "", "uor.", "")
	members_map := models.ColsMap(models.MemberCols(), "member__", "", "m.", "")
	periods_map := models.ColsMap(models.PeriodCols(), "period__", "", "p.", "")
	org_units_map := models.ColsMap(models.OrgUnitCols(), "org_unit__", "", "o.", "")
	cols := []interface{}{}
	for k, v := range user_cols_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range office_cols_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range office_req_cols_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range user_office_req_cols_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range members_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range periods_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}
	for k, v := range org_units_map {
		cols = append(
			cols,
			goqu.I(v).As(k),
		)
	}

	var userOfficeRequest models.UserOfficeRequest

	sel := r.sql.
		Select(cols...).
		From(goqu.I(r.TableName()).As("uor")).
		Where(goqu.Ex{
			fmt.Sprintf("uor.%s", r.PrimaryKey()): id,
			"uor.deleted_at":                      nil,
			"u.deleted_at":                        nil,
			"m.deleted_at":                        nil,
			"or.deleted_at":                       nil,
			"of.deleted_at":                       nil,
			"o.deleted_at":                        nil,
			"p.deleted_at":                        nil,
		}).
		Join(
			goqu.I("users").As("u"),
			goqu.On(goqu.Ex{"u.user_id": goqu.I("uor.user_id")}),
		).
		Join(
			goqu.I("members").As("m"),
			goqu.On(goqu.Ex{"m.member_id": goqu.I("u.member_id")}),
		).
		Join(
			goqu.I("office_requests").As("or"),
			goqu.On(goqu.Ex{"or.office_request_id": goqu.I("uor.office_request_id")}),
		).
		Join(
			goqu.I("offices").As("of"),
			goqu.On(goqu.Ex{"of.office_id": goqu.I("uor.office_id")}),
		).
		Join(
			goqu.I("org_units").As("o"),
			goqu.On(goqu.Ex{"o.org_unit_id": goqu.I("uor.org_unit_id")}),
		).
		Join(
			goqu.I("periods").As("p"),
			goqu.On(goqu.Ex{"p.period_id": goqu.I("uor.period_id")}),
		)

	query, params, err := sel.ToSQL()
	if err != nil {
		return userOfficeRequest, err
	}

	rows, err := r.DB.QueryxContext(ctx, query, params...)
	if err != nil {
		return userOfficeRequest, err
	}

	var user models.User
	var office models.Office
	var officeRequest models.OfficeRequest
	var member models.Member
	var orgUnit models.OrgUnit
	var period models.Period

	defer rows.Close()
	for rows.Next() {
		for j := 0; j < 7; j++ {
			var (
				dest       interface{}
				searchText string
			)
			if j == 0 {
				dest = &user
				searchText = "user__"
			} else if j == 1 {
				dest = &office
				searchText = "office__"
			} else if j == 2 {
				dest = &officeRequest
				searchText = "office_request__"
			} else if j == 3 {
				dest = &userOfficeRequest
				searchText = "uor__"
			} else if j == 4 {
				dest = &member
				searchText = "member__"
			} else if j == 5 {
				dest = &period
				searchText = "period__"
			} else if j == 6 {
				dest = &orgUnit
				searchText = "org_unit__"
			}

			err = helpers.ScanStruct(dest, rows, searchText, "")
			if err != nil {
				return userOfficeRequest, err
			}
		}

		user.Member = &member
		userOfficeRequest.Office = &office
		userOfficeRequest.User = &user
		userOfficeRequest.Period = &period
		userOfficeRequest.OrgUnit = &orgUnit
		userOfficeRequest.OfficeRequest = &officeRequest

		break
	}

	if userOfficeRequest.UserOfficeRequestID.String() == "" || userOfficeRequest.OfficeID.String() == "" || userOfficeRequest.UserID.String() == "" {
		return userOfficeRequest, ErrNotFound
	}

	return userOfficeRequest, nil
}
