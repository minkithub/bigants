// Package db 는 데이터베이스 커넥션과 트랜잭션 관리를 구현한다.
package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Database 구조체는 데이터베이스 접근에 대한 전체 API를 구현한다.
type Database struct {
	*pgxpool.Pool
	opt     *DatabaseOptions
	batcher *Batcher
}

type DatabaseOptions struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	MaxOpenConns int32
	// Log          func(fmt string, a ...interface{})
}

// New 함수는 환경을 기반으로 새로운 데이터베이스를 만든다.
func New(ctx context.Context, opt DatabaseOptions) (*Database, error) {
	cfg, err := pgxpool.ParseConfig(FormatConnStr(opt.Host, opt.Port, opt.User, opt.Password, opt.Name))
	if err != nil {
		return nil, err
	}
	if opt.MaxOpenConns > 0 {
		cfg.MaxConns = opt.MaxOpenConns
	}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	batcher := NewBatcher(ctx, pool, 200, 30*time.Microsecond)

	go func() {
		<-ctx.Done()
		pool.Close()
		log.Println("Pool closed")
	}()
	return &Database{pool, &opt, batcher}, nil
}

// FormatConnStr 함수는 주어진 설정을 데이터베이스 연결 파라미터로 쓰일 문자열로 포매팅한다.
func FormatConnStr(host, port, user, password, name string) string {
	var connStrs []string
	if host != "" {
		connStrs = append(connStrs, fmt.Sprintf("host=%s", host))
	}
	if port != "" {
		connStrs = append(connStrs, fmt.Sprintf("port=%s", port))
	}
	if user != "" {
		connStrs = append(connStrs, fmt.Sprintf("user=%s", user))
	}
	if password != "" {
		connStrs = append(connStrs, fmt.Sprintf("password=%s", password))
	}
	if name != "" {
		connStrs = append(connStrs, fmt.Sprintf("dbname=%s", name))
	}
	connStrs = append(connStrs, "sslmode=disable")

	return strings.Join(connStrs, " ")
}

func (opt *DatabaseOptions) formatConnStr() string {
	return FormatConnStr(opt.Host, opt.Port, opt.User, opt.Password, opt.Name)
}
