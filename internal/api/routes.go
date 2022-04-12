package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/hasahmad/go-api/internal/api/handlers"
	"github.com/hasahmad/go-api/internal/api/middlewares"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/oauth"
	"github.com/rs/cors"
)

func (app *Application) routes() http.Handler {
	r := chi.NewRouter()

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

	oauth2Srv := oauth.SetupOAuthServer(app.DB, app.Config)
	AuthMiddleware := ms.AuthMiddlewareHandler(oauth2Srv)

	r.Use(ms.RecoverPanic)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: app.Config.Cors.TrustedOrigins,
	})
	r.Use(corsMiddleware.Handler)

	r.Use(ms.RateLimit)

	r.NotFound(http.HandlerFunc(helpers.NotFoundResponseHandler(app.Logger)))
	r.MethodNotAllowed(http.HandlerFunc(helpers.MethodNotAllowedResponseHandler(app.Logger)))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", hs.HealthCheckHandler)
		r.Get("/o/authorize", func(w http.ResponseWriter, r *http.Request) {
			err := oauth2Srv.HandleAuthorizeRequest(w, r)
			if err != nil {
				helpers.BadRequestResponse(app.Logger, w, r, err)
			}
		})
		r.Get("/o/token", func(w http.ResponseWriter, r *http.Request) {
			err := oauth2Srv.HandleTokenRequest(w, r)
			if err != nil {
				helpers.BadRequestResponse(app.Logger, w, r, err)
			}
		})

		r.Group(func(gr chi.Router) {
			gr.Use(AuthMiddleware)
			gr.Get("/users", hs.GetAllUsersHandler)
			gr.Get("/users/:id", hs.GetUserHandler)
			gr.Put("/users/:id", hs.UpdateUserHandler)
			gr.Delete("/users/:id", hs.DeleteUserHandler)
		})
	})

	return r
}
