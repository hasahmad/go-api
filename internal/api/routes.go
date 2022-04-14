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

			r.Get("/profile", hs.GetProfileHandler)

			r.Route("/users", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					[]string{"BROWSE-USER"},
					true,
					hs.GetAllUsersHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					[]string{"ADD-USER"},
					true,
					hs.CreateUserHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					[]string{"READ-USER"},
					true,
					hs.GetUserHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					[]string{"EDIT-USER"},
					true,
					hs.UpdateUserHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					[]string{"DELETE-USER"},
					true,
					hs.DeleteUserHandler,
				))
				r.Get("/{id}/roles", ms.RequirePermissionHandler(
					[]string{"READ-USER-ROLE"},
					true,
					hs.GetUserRolesHandler,
				))
				r.Post("/{id}/roles/{role_id}", ms.RequirePermissionHandler(
					[]string{"CREATE-USER-ROLE"},
					true,
					hs.CreateUserRoleHandler,
				))
				r.Delete("/{id}/roles/{role_id}", ms.RequirePermissionHandler(
					[]string{"DELETE-USER-ROLE"},
					true,
					hs.DeleteUserRoleHandler,
				))
			})

			// TODO: add middleware check if has permission to view this route
			r.Route("/roles", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					[]string{"BROWSE-ROLE"},
					true,
					hs.GetAllRolesHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					[]string{"CREATE-ROLE"},
					true,
					hs.CreateRoleHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					[]string{"READ-ROLE"},
					true,
					hs.GetRoleHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					[]string{"EDIT-ROLE"},
					true,
					hs.UpdateRoleHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					[]string{"DELETE-ROLE"},
					true,
					hs.DeleteRoleHandler,
				))
				r.Get("/{id}/permissions", ms.RequirePermissionHandler(
					[]string{"READ-ROLE-PERMISSION"},
					true,
					hs.GetRolePermissionsHandler,
				))
				r.Post("/{id}/permissions/{permission_id}", ms.RequirePermissionHandler(
					[]string{"CREATE-ROLE-PERMISSION"},
					true,
					hs.CreateRolePermissionRoleHandler,
				))
				r.Delete("/{id}/permissions/{permission_id}", ms.RequirePermissionHandler(
					[]string{"DELETE-ROLE-PERMISSION"},
					true,
					hs.DeleteRolePermissionHandler,
				))
			})

			// TODO: add middleware check if has permission to view this route
			r.Route("/permissions", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					[]string{"BROWSE-PERMISSION"},
					true,
					hs.GetAllPermissionsHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					[]string{"CREATE-PERMISSION"},
					true,
					hs.CreatePermissionHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					[]string{"READ-PERMISSION"},
					true,
					hs.GetPermissionHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					[]string{"EDIT-PERMISSION"},
					true,
					hs.UpdatePermissionHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					[]string{"DELETE-PERMISSION"},
					true,
					hs.DeletePermissionHandler,
				))
				r.Get("/{id}/roles", ms.RequirePermissionHandler(
					[]string{"READ-ROLE-PERMISSION"},
					true,
					hs.GetPermissionRolesHandler,
				))
			})
		})
	})

	return r
}
