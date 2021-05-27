package middlewares

import (
	"log"
	"net/http"
)

func (m Middleware) RecoverHandler(hdl http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover from panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		hdl.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
