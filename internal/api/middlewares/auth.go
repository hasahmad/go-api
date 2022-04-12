package middlewares

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/google/uuid"
	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/hasahmad/go-api/internal/helpers"
)

func (m Middlewares) AuthMiddlewareHandler(oauth2Server *server.Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := oauth2Server.UserAuthorizationHandler(w, r)
			if err != nil {
				helpers.BadRequestResponse(m.Logger, w, r, err)
				return
			}

			user, err := m.Repositories.Users.FindById(r.Context(), uuid.MustParse(userID))
			if err != nil {
				helpers.BadRequestResponse(m.Logger, w, r, err)
				return
			}

			ctx := apicontext.SetUser(r.Context(), &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
