package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	apicontext "github.com/hasahmad/go-api/internal/api/context"
)

func (m Middlewares) AssignRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-REQUEST-ID")
		if len(reqID) == 0 {
			reqID = uuid.NewString()
		}

		r = r.WithContext(apicontext.SetRequestID(r.Context(), reqID))
		w.Header().Set("X-REQUEST-ID", reqID)

		next.ServeHTTP(w, r)
	})
}
