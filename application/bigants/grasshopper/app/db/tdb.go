package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v4"
)

type TDB struct {
	*Database
	name         string
	maintainance *pgx.Conn
}

// NewTDB 함수는 데이터베이스의 복제판을 만든다.
func NewTDB(
	ctx context.Context,
	opt DatabaseOptions,
) (tdb *TDB, err error) {
	rand.Seed(time.Now().UnixNano())

	tdbName := fmt.Sprintf("__testdb%d", rand.Int())
	maintainance, err := pgx.Connect(ctx, opt.formatConnStr())
	if err != nil {
		return nil, err
	}
	_, err = maintainance.Exec(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, tdbName))
	if err != nil {
		return nil, err
	}
	tdb = &TDB{nil, tdbName, maintainance}
	// 여기부턴 오류 있으면 데이터베이스 드롭해주어야 함
	defer func() {
		if err != nil {
			tdb.Close()
		}
	}()
	tdbOpt := opt
	tdbOpt.Name = tdbName
	if tdb.Database, err = New(ctx, tdbOpt); err != nil {
		return nil, err
	}
	return tdb, nil
}

func (tdb *TDB) Migrate(ctx context.Context, migrationDir string) (err error) {
	var tx pgx.Tx
	if tx, err = tdb.Begin(ctx); err != nil {
		return err
	}
	m := NewMigrator(tx)
	m.Log = log.Printf

	if m.MigrateToLatestVersion(ctx, migrationDir); err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("population error: %w", err)
	}
	return tx.Commit(ctx)
}

func (tdb *TDB) Close() {
	if tdb.Pool != nil {
		tdb.Pool.Close()
	}
	sql := fmt.Sprintf(`DROP DATABASE "%s"`, tdb.name)
	_, err := tdb.maintainance.Exec(context.Background(), sql)
	log.Println(sql)
	if err != nil {
		log.Println(err)
	}
}
