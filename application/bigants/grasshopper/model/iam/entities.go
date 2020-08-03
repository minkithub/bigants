package iam

import (
	"time"
)

func (s *Service) NewUser() *User {
	return &User{s: s}
}

// User 구조체는 어플리케이션의 사용자이다.
type User struct {
	s        *Service
	ID       string    `json:"id"`
	Created  time.Time `json:"created"`
	Password string    `json:"password"`
}

func (ent *User) MatchPassword(plain string) error {
	if ok, err := ent.s.secret.Compare(ent.Password, plain); err != nil {
		return err
	} else if !ok {
		return ErrPasswordNotMatch
	}
	return nil
}
