package middlewares

import (
	"net/http"
	"time"

	apicontext "github.com/hasahmad/go-api/internal/api/context"
	"github.com/sirupsen/logrus"
)

func (m Middlewares) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			m.Logger.WithFields(logrus.Fields{
				"request_id": apicontext.GetRequestID(r.Context()),
				"method":     r.Method,
				"duration":   time.Since(start).String(),
				"host":       r.Host,
				"path":       r.URL.Path,
				"query":      r.URL.RawQuery,
			}).Info("request")
		}()

		next.ServeHTTP(w, r)
	})
}
