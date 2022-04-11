package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/repository"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	// Import the pq driver so that it can register itself with the database/sql
	// package.
	_ "github.com/lib/pq"
)

type Application struct {
	Config       config.Config
	Logger       *log.Logger
	DB           *sqlx.DB
	Repositories repository.Repositories
	Wg           sync.WaitGroup
}

func StartServer() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	db, err := helpers.OpenDbConnection(cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	repos := repository.New(db, cfg)

	app := Application{
		Config:       cfg,
		Logger:       logger,
		DB:           db,
		Repositories: repos,
	}

	err = app.Serve()
	if err != nil {
		logger.Fatal(err)
	}
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// shutdownError channel will be used to receive any errors returned
	// by the graceful Shutdown() function.
	shutdownError := make(chan error)

	go func() {
		// create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)

		// Use signal.Notify() to ligten for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Read Read the signal from the quit channel. The code will block until
		// the signal is received
		s := <-quit

		app.Logger.WithFields(log.Fields{
			"signal": s.String(),
		}).Info("caught signal")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Call Shutdown() to the server with above context
		// Shutdown() will return nil if graceful shutdown was successfull
		// or an error (which may happen because of a problem closing the
		// listeners, or because the shutdown didn't complete before the 5-second
		// context deadline is hit).
		// Relay the return value to shutdownError channel if has error
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		app.Logger.WithFields(log.Fields{
			"add": srv.Addr,
		}).Info("completing background tasks")
		app.Wg.Wait()
		shutdownError <- nil
	}()

	app.Logger.WithFields(log.Fields{
		"addr": srv.Addr,
	}).Info("starting server")

	// Calling Shutdown() on the server will cause ListenAndServe() to immediately
	// return a http.ErrServerClosed error. So if we see this error, it is actually a
	// good thing and an indication that the graceful shutdown has started. So we check
	// specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// Otherwise, wait to receive the return value from Shutdown() on the
	// shutdownError channel. If return value is an error, we know that there was a
	// problem with the graceful shutdown and we return the error.
	err = <-shutdownError
	if err != nil {
		return err
	}

	// we know that the graceful shutdown completed successfully and we
	// log a "stopped server" message.
	app.Logger.WithFields(log.Fields{
		"addr": srv.Addr,
	}).Info("stopped server")

	return nil
}
