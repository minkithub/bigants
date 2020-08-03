package resolver

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/list"
)

type StockPredictionConnection struct {
	root *Resolver
	src  list.Connection
}

func (rsv *StockPredictionConnection) Edges(ctx context.Context) []StockPredictionEdge {
	return rsv.root.ResolveStockPredictionEdges(rsv.src.Cursors)
}

func (rsv *StockPredictionConnection) PageInfo(ctx context.Context) PageInfo {
	return rsv.root.ResolvePageInfo(rsv.src.Cursors, rsv.src.PageInfo)
}

func (rsv *Resolver) ResolveStockPredictionEdges(cursors []string) []StockPredictionEdge {
	next := make([]StockPredictionEdge, len(cursors))
	for i := range cursors {
		next[i] = rsv.ResolveStockPredictionEdge(cursors[i])
	}
	return next
}

func (rsv *Resolver) ResolveStockPredictionEdge(cursor string) StockPredictionEdge {
	return StockPredictionEdge{rsv, cursor}
}

type StockPredictionEdge struct {
	root *Resolver
	ID   string
}

func (rsv StockPredictionEdge) Cursor() string { return rsv.ID }
func (rsv StockPredictionEdge) Node(ctx context.Context) (*StockPrediction, error) {
	return rsv.root.ResolveStockPrediction(ctx, rsv.ID)
}
