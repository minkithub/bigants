package resolver

import (
	"context"
	"time"

	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/ident"
	"github.com/graph-gophers/graphql-go"
)

func (rsv *Resolver) ResolveStockPredictions(ctx context.Context, ids ...string) (nexts []*StockPrediction, err error) {
	srcs, err := rsv.saps.GetStockPredictions(ctx, ids...)
	if err != nil {
		return nil, err
	}
	nexts = make([]*StockPrediction, len(ids))
	for i, src := range srcs {
		if src != nil {
			nexts[i] = &StockPrediction{rsv, src}
		}
	}
	return nexts, nil
}

func (rsv *Resolver) ResolveStockPrediction(ctx context.Context, id string) (next *StockPrediction, err error) {
	nexts, err := rsv.ResolveStockPredictions(ctx, id)
	return nexts[0], err
}

type StockPrediction struct {
	root *Resolver
	src  *sap.StockPrediction
}

func (rsv *StockPrediction) ID() graphql.ID         { return ident.EncodeID("StockPrediction", rsv.src.ID) }
func (rsv *StockPrediction) Created() string        { return rsv.src.Created.Format(time.RFC3339) }
func (rsv *StockPrediction) Start() string          { return rsv.src.StartDate.String() }
func (rsv *StockPrediction) End() string            { return rsv.src.EndDate.String() }
func (rsv *StockPrediction) AverageIncome() float64 { return rsv.src.AverageIncome() }
func (rsv *StockPrediction) Accuracy() float64      { return rsv.src.Accuracy() }
func (rsv *StockPrediction) MAE() float64           { return rsv.src.MAE() }

func (rsv *StockPrediction) Stock(ctx context.Context) (*Stock, error) {
	return rsv.root.ResolveStock(ctx, rsv.src.StockCode)
}

func (rsv *StockPrediction) Requester(ctx context.Context) (*User, error) {
	if rsv.src.RequesterID == nil {
		return nil, nil
	}
	return rsv.root.ResolveUser(ctx, *rsv.src.RequesterID)
}

func (rsv *StockPrediction) DailyPredictions() (nexts []*StockDailyPrediction) {
	srcs := rsv.src.DailyPredictions()
	nexts = make([]*StockDailyPrediction, len(srcs))
	for i, src := range srcs {
		nexts[i] = &StockDailyPrediction{rsv.root, src}
	}
	return nexts
}

type StockDailyPrediction struct {
	root *Resolver
	src  *sap.MergedDailyPrediction
}

func (rsv *StockDailyPrediction) Date() string            { return rsv.src.Date.String() }
func (rsv *StockDailyPrediction) ExpectedPrice() float64  { return rsv.src.ExpectedPrice() }
func (rsv *StockDailyPrediction) ExpectedIncome() float64 { return rsv.src.ExpectedIncome() }
