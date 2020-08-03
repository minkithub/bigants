package web

import (
	"net/http"
)

func withCors(allowedOrigins []string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		var isAllowed bool

		for i := range allowedOrigins {
			if origin == allowedOrigins[i] {
				isAllowed = true
				break
			}
		}
		wh := w.Header()

		if isAllowed {
			wh.Set("Access-Control-Allow-Origin", origin)
		}
		if r.Method == http.MethodOptions {
			if isAllowed {
				wh.Set("Access-Control-Allow-Credentials", "true")
				wh.Set("Access-Control-Allow-Headers", "DeviceID, Content-Type")
				wh.Set("Access-Control-Max-Age", "86400") // seconds
			}
			w.WriteHeader(200)
			return
		}
		h.ServeHTTP(w, r)
	})
}
