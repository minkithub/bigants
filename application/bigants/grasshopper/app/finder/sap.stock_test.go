package finder_test

import (
	"context"
	"testing"

	"github.com/allscape/bigants/grasshopper/pkg/assertion"

	"github.com/kr/pretty"

	"github.com/allscape/bigants/grasshopper/app/finder"
	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

func TestGetStocksByCode(t *testing.T) {
	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	var a sap.StockFinder = &finder.Adapter{tdb}

	ents, err := a.GetStocksByCode(ctx, "097950.KS", "108670.KS")
	if err != nil {
		t.Fatal(err)
	}
	pretty.Log(ents)
}

func TestPaginateStocks(t *testing.T) {
	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	var a sap.StockFinder = &finder.Adapter{tdb}

	c, err := a.PaginateStocks(ctx, list.PageOption{}, sap.StockFilter{})
	if err != nil {
		t.Fatal(err)
	}
	if err = assertion.TestJSONSnapshot(c, "snapshots/TestPaginateStocks.json"); err != nil {
		t.Fatal(err)
	}
}

func TestPaginateStocksFilteringQ(t *testing.T) {
	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	q := "제일제당"

	var a sap.StockFinder = &finder.Adapter{tdb}

	c, err := a.PaginateStocks(ctx, list.PageOption{}, sap.StockFilter{Q: &q})
	if err != nil {
		t.Fatal(err)
	}
	if err = assertion.TestJSONSnapshot(c, "snapshots/TestPaginateStocksFilteringQ.json"); err != nil {
		t.Fatal(err)
	}
}

func TestFindStockHistory(t *testing.T) {
	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	var a sap.StockFinder = &finder.Adapter{tdb}

	h, err := a.FindStockHistory(ctx, "097950.KS", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h)
	// if err = assertion.TestJSONSnapshot(h, "snapshots/TestFindStockHistory.json"); err != nil {
	// 	t.Fatal(err)
	// }
}
