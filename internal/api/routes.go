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

	oauth2Srv := oauth.SetupOAuthServer(app.DB, app.Config, app.Logger)
	AuthMiddleware := ms.AuthMiddlewareHandler(oauth2Srv)

	r.Use(middleware.Recoverer)
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
		r.Post("/o/authorize", func(w http.ResponseWriter, req *http.Request) {
			err := oauth2Srv.HandleAuthorizeRequest(w, req)
			if err != nil {
				helpers.BadRequestResponse(app.Logger, w, req, err)
			}
		})
		r.Post("/o/token", func(w http.ResponseWriter, req *http.Request) {
			oauth2Srv.HandleTokenRequest(w, req)
		})

		r.Group(func(r chi.Router) {
			r.Use(AuthMiddleware)
			r.Route("/users", func(r chi.Router) {
				r.Get("/", hs.GetAllUsersHandler)
				r.Post("/", hs.CreateUserHandler)
				r.Get("/{id}", hs.GetUserHandler)
				r.Put("/{id}", hs.UpdateUserHandler)
				r.Delete("/{id}", hs.DeleteUserHandler)
				r.Get("/{id}/roles", hs.GetUserRolesHandler)
			})

			// TODO: add middleware check if has permission to view this route
			r.Route("/roles", func(r chi.Router) {
				r.Get("/", hs.GetAllRolesHandler)
				r.Post("/", hs.CreateRoleHandler)
				r.Get("/{id}", hs.GetRoleHandler)
				r.Put("/{id}", hs.UpdateRoleHandler)
				r.Delete("/{id}", hs.DeleteRoleHandler)
				r.Get("/{id}/permissions", hs.GetRolePermissionsHandler)
			})

			// TODO: add middleware check if has permission to view this route
			r.Route("/permissions", func(r chi.Router) {
				r.Get("/", hs.GetAllPermissionsHandler)
				r.Post("/", hs.CreatePermissionHandler)
				r.Get("/{id}", hs.GetPermissionHandler)
				r.Put("/{id}", hs.UpdatePermissionHandler)
				r.Delete("/{id}", hs.DeletePermissionHandler)
				r.Get("/{id}/roles", hs.GetPermissionRolesHandler)
			})
		})
	})

	return r
}
