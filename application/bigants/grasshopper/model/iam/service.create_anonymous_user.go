package iam

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/runctx"
)

// CreateAnonymousUser 메서드는 무작위 비밀번호의 유저를 새로 생성하고 그 참조와 평문 비밀번호를 반환한다.
func (s *Service) CreateAnonymousUser(ctx context.Context) (ent *User, plain string, err error) {
	plain, err = s.secret.RandomPassword()
	if err != nil {
		return nil, "", err
	}
	pw, err := s.secret.Hash(plain)
	if err != nil {
		return nil, "", err
	}
	ent = &User{
		s:        s,
		ID:       runctx.UUID(ctx),
		Created:  runctx.Now(ctx),
		Password: pw,
	}
	return ent, plain, nil
}
