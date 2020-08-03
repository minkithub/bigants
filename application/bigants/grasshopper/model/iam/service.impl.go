package iam

import (
	"context"
	"fmt"
)

func (s *Service) FindUserOrFail(ctx context.Context, id string) (*User, error) {
	ents, err := s.UserFinder.GetUsers(ctx, s, id)
	if err != nil {
		return nil, err
	}
	if ents[0] == nil {
		return nil, fmt.Errorf("User: Does not exists(%s)", id)
	}
	return ents[0], nil
}
