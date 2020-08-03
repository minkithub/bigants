package web

import (
	"net/http"

	"github.com/allscape/bigants/grasshopper/app/auth"
	"github.com/allscape/bigants/grasshopper/model/iam"
	// "github.com/allscape/timeandspace/backend/service/web/exception"
)

func withAuth(authService *auth.Service, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()
		val := r.Header.Get("Authorization")
		var user *iam.User
		if val != "" {
			user, err = authService.DecodeToken(ctx, val)
			if err != nil {
				handleError(w, err)
				return
			}
			ctx = iam.WithActor(ctx, user)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
