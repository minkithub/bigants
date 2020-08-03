package anthill

import (
	"context"
	"encoding/json"
)

type HistoryRequest struct {
	Code  string `json:"code"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type HistoryResponse struct {
	Data json.RawMessage `json:"data"`
}

func (c *Client) History(ctx context.Context, req HistoryRequest) (res *HistoryResponse, err error) {

	if err = c.doAPIPostRequest(ctx, "/api/v1/history", req, &res); err != nil {
		return nil, err
	}
	return res, nil

}
