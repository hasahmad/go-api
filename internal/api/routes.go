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
					"BROWSE-USER",
					hs.GetAllUsersHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"ADD-USER",
					hs.CreateUserHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-USER",
					hs.GetUserHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-USER",
					hs.UpdateUserHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-USER",
					hs.DeleteUserHandler,
				))
				r.Get("/{id}/roles", ms.RequirePermissionHandler(
					"READ-USER-ROLE",
					hs.GetUserRolesHandler,
				))
				r.Post("/{id}/roles/{role_id}", ms.RequirePermissionHandler(
					"CREATE-USER-ROLE",
					hs.CreateUserRoleHandler,
				))
				r.Delete("/{id}/roles/{role_id}", ms.RequirePermissionHandler(
					"DELETE-USER-ROLE",
					hs.DeleteUserRoleHandler,
				))
			})

			r.Route("/roles", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-ROLE",
					hs.GetAllRolesHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-ROLE",
					hs.CreateRoleHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-ROLE",
					hs.GetRoleHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-ROLE",
					hs.UpdateRoleHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-ROLE",
					hs.DeleteRoleHandler,
				))
				r.Get("/{id}/permissions", ms.RequirePermissionHandler(
					"READ-ROLE-PERMISSION",
					hs.GetRolePermissionsHandler,
				))
				r.Post("/{id}/permissions/{permission_id}", ms.RequirePermissionHandler(
					"CREATE-ROLE-PERMISSION",
					hs.CreateRolePermissionRoleHandler,
				))
				r.Delete("/{id}/permissions/{permission_id}", ms.RequirePermissionHandler(
					"DELETE-ROLE-PERMISSION",
					hs.DeleteRolePermissionHandler,
				))
			})

			r.Route("/permissions", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-PERMISSION",
					hs.GetAllPermissionsHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-PERMISSION",
					hs.CreatePermissionHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-PERMISSION",
					hs.GetPermissionHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-PERMISSION",
					hs.UpdatePermissionHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-PERMISSION",
					hs.DeletePermissionHandler,
				))
				r.Get("/{id}/roles", ms.RequirePermissionHandler(
					"READ-ROLE-PERMISSION",
					hs.GetPermissionRolesHandler,
				))
			})

			r.Route("/departments", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-DEPARTMENT",
					hs.GetAllDepartmentsHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-DEPARTMENT",
					hs.CreateDepartmentHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-DEPARTMENT",
					hs.GetDepartmentHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-DEPARTMENT",
					hs.UpdateDepartmentHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-DEPARTMENT",
					hs.DeleteDepartmentHandler,
				))
			})

			r.Route("/org-units", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-ORG-UNIT",
					hs.GetAllOrgUnitsHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-ORG-UNIT",
					hs.CreateOrgUnitHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-ORG-UNIT",
					hs.GetOrgUnitHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-ORG-UNIT",
					hs.UpdateOrgUnitHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-ORG-UNIT",
					hs.DeleteOrgUnitHandler,
				))
			})

			r.Route("/offices", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-OFFICE",
					hs.GetAllOfficesHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-OFFICE",
					hs.CreateOfficeHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-OFFICE",
					hs.GetOfficeHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-OFFICE",
					hs.UpdateOfficeHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-OFFICE",
					hs.DeleteOfficeHandler,
				))
				r.Get("/{id}/requests", ms.RequirePermissionHandler(
					"BROWSE-OFFICE-REQUEST",
					hs.GetOfficeRequestsHandler,
				))
			})

			r.Route("/office-requests", func(r chi.Router) {
				r.Get("/", ms.RequirePermissionHandler(
					"BROWSE-OFFICE-REQUEST",
					hs.GetAllOfficeRequestsHandler,
				))
				r.Post("/", ms.RequirePermissionHandler(
					"CREATE-OFFICE-REQUEST",
					hs.CreateOfficeRequestHandler,
				))
				r.Get("/{id}", ms.RequirePermissionHandler(
					"READ-OFFICE-REQUEST",
					hs.GetOfficeRequestHandler,
				))
				r.Put("/{id}", ms.RequirePermissionHandler(
					"EDIT-OFFICE-REQUEST",
					hs.UpdateOfficeRequestHandler,
				))
				r.Delete("/{id}", ms.RequirePermissionHandler(
					"DELETE-OFFICE-REQUEST",
					hs.DeleteOfficeRequestHandler,
				))
			})
		})
	})

	return r
}
