package anthill

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/runctx"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/date"
)

var _ sap.Anthill = (*Client)(nil)

// PredictClose 는 sap.Anthill 인터페이스를 구현한다.
func (c *Client) PredictClose(ctx context.Context, req sap.AnthillPredictRequest) (ent *sap.SeriesPrediction, err error) {
	resp, err := c.Predict(ctx, newPredictRequest(req, PredictClose))
	if err != nil {
		return nil, err
	}
	return newSeriesPrediction(ctx, req.End, resp), nil
}

// PredictHigh 는 sap.Anthill 인터페이스를 구현한다.
func (c *Client) PredictHigh(ctx context.Context, req sap.AnthillPredictRequest) (ent *sap.SeriesPrediction, err error) {
	resp, err := c.Predict(ctx, newPredictRequest(req, PredictHigh))
	if err != nil {
		return nil, err
	}
	return newSeriesPrediction(ctx, req.End, resp), nil
}

// PredictLow 는 sap.Anthill 인터페이스를 구현한다.
func (c *Client) PredictLow(ctx context.Context, req sap.AnthillPredictRequest) (ent *sap.SeriesPrediction, err error) {
	resp, err := c.Predict(ctx, newPredictRequest(req, PredictLow))
	if err != nil {
		return nil, err
	}
	return newSeriesPrediction(ctx, req.End, resp), nil
}

func newPredictRequest(req sap.AnthillPredictRequest, t PredictType) PredictRequest {
	pReq := PredictRequest{
		Code:     req.StockCode,
		End:      req.End.String(),
		Start:    req.Start.String(),
		Type:     t,
		Holidays: make([]string, len(req.Holidays)),
	}
	for i, h := range req.Holidays {
		pReq.Holidays[i] = h.String()
	}
	return pReq
}

func newSeriesPrediction(ctx context.Context, end date.Date, resp *PredictResponse) (ent *sap.SeriesPrediction) {
	ent = &sap.SeriesPrediction{
		Created:          runctx.Now(ctx),
		MAE:              resp.MAE,
		MAPE:             resp.MAPE,
		DailyPredictions: make([]*sap.DailyPrediction, len(resp.Data)),
	}
	for i, d := range resp.Data {
		ent.DailyPredictions[i] = &sap.DailyPrediction{
			Date:  end.AddDays(i + 1),
			Value: d,
		}
	}
	return ent
}
