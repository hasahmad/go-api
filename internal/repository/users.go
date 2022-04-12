package repository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/models"
	"github.com/hasahmad/go-api/pkg/filters"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB  *sqlx.DB
	sql *goqu.Database
}

func NewUserRepo(db *sqlx.DB, cfg config.Config, sql *goqu.Database) UserRepo {
	if sql != nil {
		sql = goqu.New(cfg.DB.Type, db)
	}
	return UserRepo{db, sql}
}

func (r UserRepo) TableName() string {
	return "users"
}

func (r UserRepo) FindAll(ctx context.Context, wheres []goqu.Expression, f *filters.Filters) ([]models.User, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(goqu.Ex{"is_active": true})

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

	var users []models.User
	err := sel.ScanStructsContext(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepo) FindOneBy(ctx context.Context, where goqu.Ex) (models.User, error) {
	sel := r.sql.
		From(r.TableName()).
		Where(goqu.Ex{"is_active": true}).
		Where(where).
		Limit(1)

	var user models.User
	found, err := sel.ScanStructContext(ctx, &user)
	if err != nil {
		return user, err
	}
	if !found {
		return user, ErrNotFound
	}

	return user, nil
}

func (r UserRepo) FindById(ctx context.Context, userId uuid.UUID) (models.User, error) {
	return r.FindOneBy(ctx, goqu.Ex{"user_id": userId})
}

func (r UserRepo) FindByUsername(ctx context.Context, username string) (models.User, error) {
	return r.FindOneBy(ctx, goqu.Ex{"username": username})
}

func (r UserRepo) FindByEmail(ctx context.Context, email string) (models.User, error) {
	return r.FindOneBy(ctx, goqu.Ex{"email": email})
}

func (r UserRepo) Insert(ctx context.Context, u models.User) (models.User, error) {
	sel := r.sql.
		Insert(r.TableName()).
		Rows(u).
		Returning(goqu.T(r.TableName()).All())

	var user models.User
	found, err := sel.Executor().ScanStructContext(ctx, &user)
	if err != nil {
		return user, err
	}
	if !found {
		return user, ErrNotFound
	}

	return user, nil
}

func (r UserRepo) Update(ctx context.Context, userId uuid.UUID, version int, data helpers.Envelope) (models.User, error) {
	sel := r.sql.
		Update(r.TableName()).
		Set(data).
		Where(goqu.Ex{"is_active": true, "user_id": userId, "version": version}).
		Returning(goqu.T(r.TableName()).All())

	var user models.User
	found, err := sel.Executor().ScanStructContext(ctx, &user)
	if err != nil {
		return user, err
	}
	if !found {
		return user, ErrNotFound
	}

	return user, nil
}

func (r UserRepo) Delete(ctx context.Context, userId uuid.UUID) error {
	sel := r.sql.
		Delete(r.TableName()).
		Where(goqu.Ex{"is_active": true, "user_id": userId}).
		Limit(1)

	_, err := sel.Executor().QueryContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
