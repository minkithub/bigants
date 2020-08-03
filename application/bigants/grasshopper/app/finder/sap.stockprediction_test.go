package finder_test

import (
	"context"
	"testing"

	"github.com/allscape/bigants/grasshopper/app/finder"
	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/assertion"
)

func TestGetStockPredictions(t *testing.T) {
	var err error

	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var fa sap.StockPredictionFinder = &finder.Adapter{tdb}
	ids := []string{
		"5d4e1621-c153-4094-b513-83ec0220be87",
		"5c86780a-2fb6-4921-b6f7-a0624900ad53", // null id
		"5c86780a-2fb6-4921-b6f7-a0624900ae53",
	}
	ents, err := fa.GetStockPredictions(ctx, ids...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ents) != 3 {
		t.Fatal()
	}
	if ents[0].ID != ids[0] {
		t.Fatal()
	}
	if ents[1] != nil {
		t.Fatal()
	}
	if ents[2].ID != ids[2] {
		t.Fatal()
	}
	if err = assertion.TestJSONSnapshot(ents, "snapshots/TestGetStockPredictions.json"); err != nil {
		t.Fatal(err)
	}

}

func TestPaginateStockPredictions(t *testing.T) {
	var err error

	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var fa sap.StockPredictionFinder = &finder.Adapter{tdb}

	c, err := fa.PaginateStockPredictions(ctx, list.PageOption{}, sap.PredictionFilter{})
	if err != nil {
		t.Fatal(err)
	}
	if err = assertion.TestJSONSnapshot(&c, "snapshots/TestPaginateStockPredictions.json"); err != nil {
		t.Fatal(err)
	}
}

func TestPaginateStockPredictionsFilteringUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var fa sap.StockPredictionFinder = &finder.Adapter{tdb}

	userID := "fddc8454-aefc-42ac-a20d-d9723a4c7cb8"

	c, err := fa.PaginateStockPredictions(ctx, list.PageOption{}, sap.PredictionFilter{
		RequesterID: &userID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err = assertion.TestJSONSnapshot(&c, "snapshots/TestPaginateStockPredictionsFilteringUser.json"); err != nil {
		t.Fatal(err)
	}
}
