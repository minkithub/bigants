package anthill_test

import (
	"context"
	"log"
	"testing"

	"github.com/allscape/bigants/grasshopper/app/anthill"
	"golang.org/x/sync/errgroup"
)

func TestClientPredict(t *testing.T) {
	var err error
	c := anthill.NewClient("https://anthill-ljnq6oi2dq-an.a.run.app")

	req := anthill.PredictRequest{
		Code:  "034730.KS",
		Start: "2020-01-20",
		End:   "2020-02-28",
		Holidays: []string{
			"2020-01-31",
		},
	}
	ctx := context.Background()
	println()

	runType := func(predictType anthill.PredictType) error {
		req := req
		req.Type = predictType
		log.Printf("Starting request for %s", req.Type)
		resp, err := c.Predict(ctx, req)
		log.Printf("End request for %s (%v)", req.Type, resp)
		return err
	}

	var eg errgroup.Group

	eg.Go(func() error { return runType(anthill.PredictClose) })
	eg.Go(func() error { return runType(anthill.PredictHigh) })
	eg.Go(func() error { return runType(anthill.PredictLow) })

	if err = eg.Wait(); err != nil {
		t.Fatal(err)
	}
}

func TestClientHistory(t *testing.T) {
	c := anthill.NewClient("https://anthill-ljnq6oi2dq-an.a.run.app")
	resp, err := c.History(context.Background(), anthill.HistoryRequest{
		Code:  "034730.KS",
		Start: "2015-01-04",
		End:   "2020-01-03",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(resp.Data))
}

func TestClientStocks(t *testing.T) {
	c := anthill.NewClient("https://anthill-ljnq6oi2dq-an.a.run.app")
	resp, err := c.Stocks(context.Background(), anthill.StocksRequest{
		Q: "SK",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)

	for _, p := range resp.Data {
		t.Logf("%#v", p)
	}
}
