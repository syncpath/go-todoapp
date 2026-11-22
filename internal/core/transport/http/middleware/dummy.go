package core_http_middlware

import (
	"fmt"
	"net/http"

	"github.com/syncpath/go-todoapp/internal/core/logger"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				log := logger.FromContext(ctx)

				log.Debug(fmt.Sprintf("-> before: %s", s))
				next.ServeHTTP(w, r)

				log.Debug(fmt.Sprintf("-> after: %s", s))
			})
	}
}
