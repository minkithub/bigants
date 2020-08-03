package db

import (
	"context"
	"time"
)

type keyBatcher struct{}

func (db *Database) WithBatcher(ctx context.Context, wait time.Duration, maxBatch int) context.Context {
	return context.WithValue(ctx, keyBatcher{}, NewBatcher(ctx, db.Pool, maxBatch, wait))
}

func getBatcher(ctx context.Context) *Batcher {
	val := ctx.Value(keyBatcher{})
	if val == nil {
		return nil
	}
	return val.(*Batcher)
}
