package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// WithTx 메서드는 domain Gateway를 구현한다.
func (srv *Database) WithTx(ctx context.Context, action func(ctx context.Context) error) (err error) {
	bg, ok := ctx.Value(keyTx{}).(beginner)
	if !ok {
		bg = srv.Pool
	}

	return withTx(ctx, bg, func(tx pgx.Tx) error {
		return action(context.WithValue(ctx, keyTx{}, tx))
	})
}

type txHandle interface {
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Begin(context.Context) (pgx.Tx, error)
}

type keyTx struct{}

func (srv *Database) getTx(ctx context.Context) txHandle {
	val, ok := ctx.Value(keyTx{}).(txHandle)
	if !ok {
		return srv.Pool
	}
	return val
}

type beginner interface {
	Begin(context.Context) (pgx.Tx, error)
}

func withTx(ctx context.Context, db beginner, action func(tx pgx.Tx) error) (err error) {
	var tx pgx.Tx

	if tx, err = db.Begin(ctx); err != nil {
		return err
	}
	if err = action(tx); err != nil {
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}
