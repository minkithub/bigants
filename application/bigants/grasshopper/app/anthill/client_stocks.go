package anthill

import "context"

type StocksRequest struct {
	Q string
}

type StocksResponse struct {
	Data []*Stock `json:"data"`
}

type Stock struct {
	Code   string `json:"code"`
	NameKo string `json:"name_ko"`
}

func (c *Client) Stocks(ctx context.Context, req StocksRequest) (resp *StocksResponse, err error) {
	if err = c.doAPIGetRequest(ctx, "/api/v1/stocks", map[string]string{
		"q": req.Q,
	}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
