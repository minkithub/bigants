package web

import (
	"net/http"

	"github.com/allscape/bigants/grasshopper/app/auth"
)

// NewHandler 함수는 토큰을 발행하는 핸들러를 생성한다.
func newAuthenticateHandler(auth *auth.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.DecodeToken(r.Context(), r.Header.Get("Authorization"))
		if err != nil {
			handleError(w, err)
			return
		}
		token, err := auth.SignJWT(user.ID)
		if err != nil {
			handleError(w, err)
			return
		}
		w.Write([]byte(*token))
	})
}
