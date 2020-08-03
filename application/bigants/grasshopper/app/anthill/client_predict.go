package anthill

import (
	"context"
)

type PredictRequest struct {
	Code     string      `json:"code"`
	Type     PredictType `json:"type"`
	Start    string      `json:"start"`
	End      string      `json:"end"`
	Holidays []string    `json:"holidays"`
}

type PredictType string

var (
	PredictHigh,
	PredictLow,
	PredictClose PredictType = "high", "low", "close"
)

type PredictResponse struct {
	MAE  float64   `json:"mae"`
	MAPE float64   `json:"mape"`
	Data []float64 `json:"data"`
}

func (c *Client) Predict(ctx context.Context, req PredictRequest) (res *PredictResponse, err error) {
	if err = c.doAPIPostRequest(ctx, "/api/v1/predict", req, &res); err != nil {
		return nil, err
	}
	return res, err
}
