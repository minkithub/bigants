package resolver

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/schema"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/ident"
	"github.com/graph-gophers/graphql-go"
)

func (rsv *Resolver) ResolveStocks(ctx context.Context, ids ...string) (nexts []*Stock, err error) {
	srcs, err := rsv.saps.GetStocksByCode(ctx, ids...)
	if err != nil {
		return nil, err
	}
	nexts = make([]*Stock, len(ids))
	for i, src := range srcs {
		if src != nil {
			nexts[i] = &Stock{rsv, src}
		}
	}
	return nexts, nil
}

func (rsv *Resolver) ResolveStock(ctx context.Context, id string) (next *Stock, err error) {
	nexts, err := rsv.ResolveStocks(ctx, id)
	if err != nil {
		return nil, err
	}
	return nexts[0], nil
}

type Stock struct {
	root *Resolver
	src  *sap.Stock
}

func (rsv *Stock) ID() graphql.ID { return ident.EncodeID("Stock", rsv.src.Code) }
func (rsv *Stock) Code() string   { return rsv.src.Code }
func (rsv *Stock) NameKo() string { return rsv.src.NameKo }
func (rsv *Stock) LatestHistory() *StockHistory {
	src := rsv.src.LatestHistory()
	if src == nil {
		return nil
	}
	return &StockHistory{rsv.root, src}
}
func (rsv *Stock) PredictablePeriodStart() string { return rsv.src.PredictablePeriodStart.String() }
func (rsv *Stock) PredictablePeriodEnd() string   { return rsv.src.PredictablePeriodEnd.String() }
func (rsv *Stock) History(ctx context.Context, args schema.StockHistoryArgs) (nexts []*StockHistory, err error) {
	srcs, err := rsv.root.saps.FindStockHistory(ctx, rsv.src.Code, args.Count)
	if err != nil {
		return nil, err
	}
	nexts = make([]*StockHistory, len(srcs))
	for i, src := range srcs {
		nexts[i] = &StockHistory{rsv.root, src}
	}
	return nexts, nil
}

type StockHistory struct {
	root *Resolver
	src  *sap.StockHistory
}

func (rsv *StockHistory) Date() string              { return rsv.src.Date.String() }
func (rsv *StockHistory) Open() float64             { return rsv.src.Open }
func (rsv *StockHistory) Close() float64            { return rsv.src.Close }
func (rsv *StockHistory) High() float64             { return rsv.src.High }
func (rsv *StockHistory) Low() float64              { return rsv.src.Low }
func (rsv *StockHistory) Volume() int32             { return rsv.src.Volume }
func (rsv *StockHistory) HighLimit() *float64       { return rsv.src.HighLimit() }
func (rsv *StockHistory) LowLimit() *float64        { return rsv.src.LowLimit() }
func (rsv *StockHistory) PriceChange() *float64     { return rsv.src.PriceChange() }
func (rsv *StockHistory) PriceChangeRate() *float64 { return rsv.src.PriceChangeRate() }
