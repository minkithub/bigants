// Package depot 는 엔티티의 영속을 구현한다.
package depot

import "context"

type Adapter struct {
	DB
}

type DB interface {
	WithTx(ctx context.Context, action func(ctx context.Context) error) (err error)
	QueryAndReceive(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	ExecAndCount(ctx context.Context, sql string, args ...interface{}) (int64, error)
}
