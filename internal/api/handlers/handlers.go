package handlers

import (
	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/repository"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Handlers struct {
	Config       config.Config
	Logger       *log.Logger
	DB           *sqlx.DB
	Repositories repository.Repositories
}
