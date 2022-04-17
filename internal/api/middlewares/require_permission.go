package middlewares

import (
	"net/http"

	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/hasahmad/go-api/internal/helpers"
)

func (m Middlewares) CheckPermissionsMiddleware(names []string, checkAll bool, w http.ResponseWriter, r *http.Request) bool {
	user := apicontext.GetUser(r.Context())

	permissions, err := m.Repositories.Permissions.FindByUserIdAndPermissionNames(r.Context(), user.UserID, names)
	if err != nil {
		helpers.ServerErrorResponse(m.Logger, w, r, err)
		return false
	}

	if (checkAll && len(permissions) != len(names)) || !checkAll && len(permissions) == 0 {
		helpers.NotPermittedResponse(m.Logger, w, r)
		return false
	}

	return true
}

func (m Middlewares) RequirePermissionHandler(permName string, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldContinue := m.CheckPermissionsMiddleware([]string{permName}, true, w, r)
		if !shouldContinue {
			return
		}

		next.ServeHTTP(w, r)
	})

	return fn
}

func (m Middlewares) RequirePermissionsHandler(names []string, checkAll bool, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldContinue := m.CheckPermissionsMiddleware(names, checkAll, w, r)
		if !shouldContinue {
			return
		}

		next.ServeHTTP(w, r)
	})

	return fn
}

func (m Middlewares) RequiresPermission(names []string, checkAll bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			shouldContinue := m.CheckPermissionsMiddleware(names, checkAll, w, r)
			if !shouldContinue {
				return
			}

			next.ServeHTTP(w, r)
		})

		return fn
	}
}
