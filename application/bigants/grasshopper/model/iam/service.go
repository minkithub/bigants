// Package iam (Identity and Access Management)은 인증, 접근 권한과 관련된 비즈니스 로직을 구현한다.
package iam

import (
	"context"
	"fmt"
)

func New(
	finder UserFinder,
	depot UserDepot,
	secret Secret,
) *Service {
	return &Service{finder, depot, secret}
}

// Service 구조체는 인증과 권한에 대한 메서드 셋을 제공한다.
type Service struct {
	UserFinder
	UserDepot
	secret Secret
}

// UserFinder 는 저장장치에서 아이디에 해당하는 유저의 슬라이스를 반환한다.
type UserFinder interface {
	GetUsers(ctx context.Context, s *Service, ids ...string) ([]*User, error)
}

func (s *Service) GetUsers(ctx context.Context, ids ...string) (ents []*User, err error) {
	if ents, err = s.UserFinder.GetUsers(ctx, s, ids...); err != nil {
		return nil, err
	}
	for _, ent := range ents {
		ent.s = s
	}
	return ents, err
}

// UserDepot 는 주어진 유저를 저장장치에 저장한다.
type UserDepot interface {
	SaveUser(ctx context.Context, user *User) error
}

// Secret 은 비밀을 취급하는 각종 함수를 제공한다.
type Secret interface {
	Hash(plain string) (encoded string, err error)
	Compare(encoded string, plain string) (ok bool, err error)
	RandomPassword() (plain string, err error)
}

var (
	// ErrPermissionDenied 오류는 권한 부족 상태를 나타낸다.
	ErrPermissionDenied = fmt.Errorf("Permission Denied")
	// ErrPasswordNotMatch 오류는 패스워드가 매치하지 않는 상태를 나타낸다.
	ErrPasswordNotMatch = fmt.Errorf("Password not match")
)
