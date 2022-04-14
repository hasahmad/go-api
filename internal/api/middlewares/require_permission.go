package middlewares

import (
	"net/http"

	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/hasahmad/go-api/internal/helpers"
)

func (m Middlewares) RequirePermissionHandler(names []string, checkAll bool, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := apicontext.GetUser(r.Context())

		permissions, err := m.Repositories.Permissions.FindByUserIdAndPermissionNames(r.Context(), user.UserID, names)
		if err != nil {
			helpers.ServerErrorResponse(m.Logger, w, r, err)
			return
		}

		println(permissions)

		if checkAll && len(permissions) != len(names) {
			helpers.NotPermittedResponse(m.Logger, w, r)
			return
		} else if !checkAll && len(permissions) > 0 {
			helpers.NotPermittedResponse(m.Logger, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})

	return fn
}

func (m Middlewares) RequirePermission(names []string, checkAll bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := apicontext.GetUser(r.Context())

			permissions, err := m.Repositories.Permissions.FindByUserIdAndPermissionNames(r.Context(), user.UserID, names)
			if err != nil {
				helpers.ServerErrorResponse(m.Logger, w, r, err)
				return
			}

			if checkAll && len(permissions) != len(names) {
				helpers.NotPermittedResponse(m.Logger, w, r)
				return
			} else if !checkAll && len(permissions) > 0 {
				helpers.NotPermittedResponse(m.Logger, w, r)
				return
			}

			next.ServeHTTP(w, r)
		})

		return fn
	}
}
