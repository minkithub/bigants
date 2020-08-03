// Package runctx 는 실행 시간 의존성을 관리
package runctx

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func Now(ctx context.Context) time.Time {
	return time.Now()
}

func UUID(ctx context.Context) string {
	return uuid.New().String()
}
