package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// QueryAndReceive 함수는 컨텍스트에 트랜잭션이 있는 경우 트랜잭션으로 쿼리하고 없으면 풀에 배칭해서 쿼리한다.
func (srv *Database) QueryAndReceive(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error) {
	val := ctx.Value(keyTx{})
	if val == nil {
		batcher := srv.batcher
		if batcher == nil { // 배처가 없는 경우에 암묵적으로 풀을 가지고 쿼리함
			return queryScan(ctx, srv.Pool, receiver, sql, args...)
		}
		return batchQueryScan(batcher, receiver, sql, args...)
	}
	tx := val.(txHandle)
	return queryScan(ctx, tx, receiver, sql, args...)
}

func (srv *Database) ExecAndCount(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	c, err := srv.getTx(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return c.RowsAffected(), nil
}

// batchQueryScan 함수는 주어진 배처, 리시버, SQL, 쿼리 인자를 통해 쿼리를 배칭해 실행하고 그 결과를 리시버에 스캔한다.
// 읽은 행의 수와 오류를 반환한다.
func batchQueryScan(batcher *Batcher, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error) {
	var rowsReceived int
	errc := make(chan error)
	batcher.BatchQuery(sql, args, func(rows pgx.Rows, err error) {
		defer close(errc)
		if err != nil {
			errc <- err
			return
		}
		for rows.Next() {
			err = rows.Scan(receiver(rowsReceived)...)
			if err != nil {
				errc <- err
				return
			}
			rowsReceived++
		}
		if err = rows.Err(); err != nil {
			errc <- err
			return
		}
	})
	return rowsReceived, <-errc
}

// queryScan 함수는 SQL쿼리 후 Receiver로 값을 읽고 읽은 행 수와 오류를 반환한다.
func queryScan(ctx context.Context, querier interface {
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
}, receiver func(int) []interface{}, sql string, args ...interface{}) (rowsReceived int, err error) {
	var rows pgx.Rows

	if rows, err = querier.Query(ctx, sql, args...); err != nil {
		return rowsReceived, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(receiver(rowsReceived)...)
		if err != nil {
			return rowsReceived, fmt.Errorf("error while queryScan(%w)", err)
		}
		rowsReceived++
	}

	return rowsReceived, rows.Err()
}
