package api

import (
	"net/http"

	"github.com/hasahmad/go-api/internal/api/handlers"
	"github.com/hasahmad/go-api/internal/api/middlewares"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(helpers.NotFoundResponseHandler(app.Logger))
	router.MethodNotAllowed = http.HandlerFunc(helpers.MethodNotAllowedResponseHandler(app.Logger))

	ms := middlewares.Middlewares{
		Config:       app.Config,
		Logger:       app.Logger,
		DB:           app.DB,
		Repositories: app.Repositories,
	}

	hs := handlers.Handlers{
		Config:       app.Config,
		Logger:       app.Logger,
		DB:           app.DB,
		Repositories: app.Repositories,
	}
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", hs.HealthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users", hs.GetAllUsersHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", hs.GetUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/:id", hs.UpdateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", hs.DeleteUserHandler)

	return ms.RecoverPanic(
		ms.EnableCORS(
			ms.RateLimit(
				ms.AssignRequestID(
					ms.RequestLogger(router),
				),
			),
		),
	)
}
