package web

import (
	"io"
	"net/http"
)

// PostUsers 핸들러는 익명의 유저를 생성하고 유저 ID와 평문 비밀번호를 베이직 토큰으로 인코딩해 반환한다.
func (s *Service) PostUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, plain, err := s.iam.CreateAnonymousUser(ctx)
	if err != nil {
		handleError(w, err)
		return
	}
	if err = s.iam.SaveUser(ctx, user); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, s.auths.EncodeBasic(user.ID, plain))
}
