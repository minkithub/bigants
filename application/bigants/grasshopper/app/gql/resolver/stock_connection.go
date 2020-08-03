package resolver

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/list"
)

func (rsv *Resolver) ResolveStockConnection(src list.Connection) *StockConnection {
	return &StockConnection{rsv, src}
}

type StockConnection struct {
	root *Resolver
	src  list.Connection
}

func (rsv *StockConnection) Edges(ctx context.Context) []StockEdge {
	return rsv.root.ResolveStockEdges(rsv.src.Cursors)
}

func (rsv *StockConnection) PageInfo(ctx context.Context) PageInfo {
	return rsv.root.ResolvePageInfo(rsv.src.Cursors, rsv.src.PageInfo)
}

func (rsv *Resolver) ResolveStockEdges(cursors []string) []StockEdge {
	next := make([]StockEdge, len(cursors))
	for i := range cursors {
		next[i] = rsv.ResolveStockEdge(cursors[i])
	}
	return next
}

func (rsv *Resolver) ResolveStockEdge(cursor string) StockEdge {
	return StockEdge{rsv, cursor}
}

type StockEdge struct {
	root    *Resolver
	StockID string
}

func (rsv StockEdge) Cursor() string { return rsv.StockID }
func (rsv StockEdge) Node(ctx context.Context) (*Stock, error) {
	return rsv.root.ResolveStock(ctx, rsv.StockID)
}
