package sap_test

import (
	"context"
	"testing"

	"github.com/allscape/bigants/grasshopper/app/anthill"
	"github.com/allscape/bigants/grasshopper/app/db"
	"github.com/allscape/bigants/grasshopper/app/depot"
	"github.com/allscape/bigants/grasshopper/app/finder"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/assertion"
	"github.com/allscape/bigants/grasshopper/pkg/date"
)

func TestCreatePrediction(t *testing.T) {
	ctx := context.Background()
	c := anthill.NewClient("https://anthill-ljnq6oi2dq-an.a.run.app")
	tdb, err := db.NewTDB(ctx, db.DatabaseOptions{
		Host:         "localhost",
		MaxOpenConns: 3,
		Name:         "postgres",
		Password:     "OmcOKiMpLJn5N6K4",
		Port:         "15432",
		User:         "postgres",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	if err = tdb.Migrate(ctx, "../../app/db/migrations"); err != nil {
		t.Fatal(err)
	}

	s := sap.Service{
		Anthill:               c,
		StockFinder:           nil,
		StockPredictionDepot:  &depot.Adapter{tdb},
		StockPredictionFinder: &finder.Adapter{tdb},
	}

	var (
		start, _    = date.Parse("2020-01-20")
		end, _      = date.Parse("2020-02-28")
		holiday1, _ = date.Parse("2020-01-31")
	)
	ent, err := s.CreatePrediction(ctx, sap.CreatePredictionRequest{
		Stock:    &sap.Stock{Code: "034730.KS"},
		Start:    start,
		End:      end,
		Holidays: []date.Date{holiday1},
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = assertion.TestJSONSnapshot(ent, "snapshots/TestCreatePrediction.json"); err != nil {
		t.Fatal(err)
	}

	// if err = s.SaveStockPrediction(ctx, ent); err != nil {
	// 	t.Fatal(err)
	// }

	// gotEnts, err := s.GetStockPredictions(ctx, ent.ID)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if gotEnts[0].ID != ent.ID {
	// 	t.Fatal()
	// }
}
