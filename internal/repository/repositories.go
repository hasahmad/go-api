package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Users UserRepo
}

func New(db *sqlx.DB, cfg config.Config) Repositories {
	sql := goqu.New(cfg.DB.Type, db)
	return Repositories{
		Users: UserRepo{db, sql},
	}
}
