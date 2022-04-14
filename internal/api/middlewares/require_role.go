package middlewares

import (
	"net/http"

	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/hasahmad/go-api/internal/helpers"
)

func (m Middlewares) CheckRolesMiddleware(names []string, checkAll bool, w http.ResponseWriter, r *http.Request) bool {
	user := apicontext.GetUser(r.Context())

	roles, err := m.Repositories.Roles.FindByUserIdAndRoleNames(r.Context(), user.UserID, names)
	if err != nil {
		helpers.ServerErrorResponse(m.Logger, w, r, err)
		return false
	}

	if checkAll && len(roles) != len(names) {
		helpers.NotPermittedResponse(m.Logger, w, r)
		return false
	} else if !checkAll && len(roles) == 0 {
		helpers.NotPermittedResponse(m.Logger, w, r)
		return false
	}

	return true
}

func (m Middlewares) RequireRoleHandler(names []string, checkAll bool, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldContinue := m.CheckRolesMiddleware(names, checkAll, w, r)
		if !shouldContinue {
			return
		}

		next.ServeHTTP(w, r)
	})

	return fn
}

func (m Middlewares) RequireRole(names []string, checkAll bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			shouldContinue := m.CheckRolesMiddleware(names, checkAll, w, r)
			if !shouldContinue {
				return
			}

			next.ServeHTTP(w, r)
		})

		return fn
	}
}
