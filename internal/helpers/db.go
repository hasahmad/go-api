package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
)

func OpenDbConnection(cfg config.DbConfig) (*sqlx.DB, error) {
	connUrl := BuildDbConnString(cfg)
	db, err := sqlx.Connect(cfg.Type, connUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// Create a context with a 5-second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext to establish a new connection to the database, passing in the
	// context created above. If connection can't be established successfully
	// within 5 second deadline, it'll return an error
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func BuildDbConnString(cfg config.DbConfig) string {
	var connUrl string = ""
	connUrl = fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Type,
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSlMode,
	)

	return connUrl
}
