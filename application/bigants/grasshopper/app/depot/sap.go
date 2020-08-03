package depot

import (
	"context"
	"time"

	"github.com/allscape/bigants/grasshopper/app/db/tables"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

// SaveStockPrediction 메서드는 sap.StockPredictionDepot를 구현한다.
func (a *Adapter) SaveStockPrediction(ctx context.Context, ent *sap.StockPrediction) (err error) {
	var (
		stpRow    *tables.StockPredictionRow
		stpTb     = tables.NewStockPredictionTable(a.DB)
		stphdRows tables.StockPredictionHolidayRows
		stphdTb   = tables.NewStockPredictionHolidayTable(a.DB)
		sepRows   tables.SeriesPredictionRows
		sepTb     = tables.NewSeriesPredictionTable(a.DB)
		sedpRows  tables.SeriesDailyPredictionRows
		sedpTb    = tables.NewSeriesDailyPredictionTable(a.DB)
	)
	stpRow = newStockPredictionRow(ent)
	sepRows = append(sepRows,
		newSeriesPredictionRow(ent, ent.CloseSeries, "C"),
		newSeriesPredictionRow(ent, ent.HighSeries, "H"),
		newSeriesPredictionRow(ent, ent.LowSeries, "L"),
	)
	stphdRows = append(stphdRows, newStockPredictionHolidayRows(ent)...)
	sedpRows = append(sedpRows, newSeriesDailyPredictionRows(ent, ent.CloseSeries, "C")...)
	sedpRows = append(sedpRows, newSeriesDailyPredictionRows(ent, ent.HighSeries, "H")...)
	sedpRows = append(sedpRows, newSeriesDailyPredictionRows(ent, ent.LowSeries, "L")...)

	if _, err = sepTb.Delete(ctx, tables.SeriesPredictionValues{StockPredictionID: &stpRow.ID}); err != nil {
		return err
	}
	if _, err = stphdTb.Delete(ctx, tables.StockPredictionHolidayValues{StockPredictionID: &stpRow.ID}); err != nil {
		return err
	}
	if _, err = sedpTb.Delete(ctx, tables.SeriesDailyPredictionValues{StockPredictionID: &stpRow.ID}); err != nil {
		return err
	}
	if err = stpTb.Save(ctx, stpRow); err != nil {
		return err
	}
	if _, err = sepTb.Insert(ctx, sepRows...); err != nil {
		return err
	}
	if _, err = stphdTb.Insert(ctx, stphdRows...); err != nil {
		return err
	}
	if _, err = sedpTb.Insert(ctx, sedpRows...); err != nil {
		return err
	}
	return nil
}

func newStockPredictionRow(ent *sap.StockPrediction) *tables.StockPredictionRow {
	// TODO: time.Time(date)?
	return tables.NewStockPredictionRow(ent.ID, time.Time(ent.StartDate), time.Time(ent.EndDate), ent.StockCode, ent.Created, ent.RequesterID)
}

func newStockPredictionHolidayRows(ent *sap.StockPrediction) tables.StockPredictionHolidayRows {
	rs := make(tables.StockPredictionHolidayRows, len(ent.Holidays))
	for i, h := range ent.Holidays {
		rs[i] = tables.NewStockPredictionHolidayRow(ent.ID, time.Time(h)) // TODO: time.Time(date)?
	}
	return rs
}

func newSeriesPredictionRow(parent *sap.StockPrediction, ent *sap.SeriesPrediction, source string) *tables.SeriesPredictionRow {
	return tables.NewSeriesPredictionRow(parent.ID, source, ent.Created, ent.MAE, ent.MAPE)
}

func newSeriesDailyPredictionRows(parent *sap.StockPrediction, ent *sap.SeriesPrediction, source string) tables.SeriesDailyPredictionRows {
	rs := make(tables.SeriesDailyPredictionRows, len(ent.DailyPredictions))
	for i, d := range ent.DailyPredictions {
		rs[i] = tables.NewSeriesDailyPredictionRow(parent.ID, source, time.Time(d.Date), d.Value) // TODO: time.Time(date)?
	}
	return rs
}
