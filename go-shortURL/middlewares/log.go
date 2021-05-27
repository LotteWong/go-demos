package middlewares

import (
	"log"
	"net/http"
	"time"
)

func (m Middleware) LogHandler(hdl http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		startAt := time.Now()
		hdl.ServeHTTP(w, r)
		endAt := time.Now()

		log.Printf("[%s] %s - %v", r.Method, r.URL.String(), endAt.Sub(startAt))
	}
	return http.HandlerFunc(fn)
}
