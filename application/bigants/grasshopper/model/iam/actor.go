package iam

import "context"

type actorKey struct{}

// WithActor 함수는 주어진 유저를 컨텍스트의 액터로 설정한다.
func WithActor(ctx context.Context, actor *User) context.Context {
	return context.WithValue(ctx, actorKey{}, actor)
}

// GetActor 함수는 컨텍스트에서 액터의 참조를 반환한다.
func GetActor(ctx context.Context) *User {
	if actor, ok := ctx.Value(actorKey{}).(*User); ok {
		return actor
	}
	return nil
}
