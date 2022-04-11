package middlewares

import (
	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/repository"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Middlewares struct {
	Config       config.Config
	Logger       *log.Logger
	DB           *sqlx.DB
	Repositories repository.Repositories
}
