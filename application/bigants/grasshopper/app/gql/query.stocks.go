package gql

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/resolver"
	"github.com/allscape/bigants/grasshopper/app/gql/schema"
	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

func (rsv *Root) Stocks(ctx context.Context, args schema.QueryStocksArgs) (*resolver.StockConnection, error) {
	conn, err := rsv.saps.PaginateStocks(ctx, list.PageOption{
		First: args.First, After: args.After, Last: args.Last, Before: args.Before,
	}, sap.StockFilter{
		Q: args.Q,
	})
	if err != nil {
		return nil, err
	}
	return getResolver(ctx).ResolveStockConnection(conn), nil
}
