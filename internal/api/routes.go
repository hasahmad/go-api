package api

import (
	"net/http"

	"github.com/hasahmad/go-api/internal/api/handlers"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(helpers.NotFoundResponseHandler(app.Logger))
	router.MethodNotAllowed = http.HandlerFunc(helpers.MethodNotAllowedResponseHandler(app.Logger))

	hs := handlers.Handlers{
		Config:       app.Config,
		Logger:       app.Logger,
		DB:           app.DB,
		Repositories: app.Repositories,
	}
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", hs.HealthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users", hs.GetAllUsersHandler)

	return router
}
