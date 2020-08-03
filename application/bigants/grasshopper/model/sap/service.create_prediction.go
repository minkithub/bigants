package sap

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/allscape/bigants/grasshopper/model/runctx"
	"github.com/allscape/bigants/grasshopper/pkg/date"
	"golang.org/x/sync/errgroup"
)

type CreatePredictionRequest struct {
	Stock     *Stock
	Start     date.Date
	End       date.Date
	Holidays  []date.Date
	Requester *iam.User
}

// CreatePrediction 함수는 CHL의 3가지 분석을 실행하고 한 StockPrediction으로 모아서 아이디를 부여, 데이터베이스에 저장해 둔다.
func (s *Service) CreatePrediction(ctx context.Context, req CreatePredictionRequest) (ent *StockPrediction, err error) {

	ent = &StockPrediction{
		ID:          runctx.UUID(ctx),
		Created:     runctx.Now(ctx),
		StartDate:   req.Start,
		EndDate:     req.End,
		RequesterID: nil,
		StockCode:   req.Stock.Code,
		Holidays:    req.Holidays,
	}
	if req.Requester != nil {
		ent.RequesterID = &req.Requester.ID
	}

	var eg errgroup.Group

	apreq := AnthillPredictRequest{req.Stock.Code, req.Start, req.End, req.Holidays}

	eg.Go(func() (err error) {
		ent.CloseSeries, err = s.Anthill.PredictClose(ctx, apreq)
		return err
	})
	eg.Go(func() (err error) {
		ent.HighSeries, err = s.Anthill.PredictHigh(ctx, apreq)
		return err
	})
	eg.Go(func() (err error) {
		ent.LowSeries, err = s.Anthill.PredictLow(ctx, apreq)
		return err
	})

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return ent, nil
}
