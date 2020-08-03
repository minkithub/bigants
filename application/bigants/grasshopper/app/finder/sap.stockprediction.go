package finder

import (
	"context"

	"github.com/allscape/bigants/grasshopper/pkg/pgcc"

	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

// GetStockPredictions 메서드는 sap.StockPredictionFinder를 구현한다.
func (a *Adapter) GetStockPredictions(ctx context.Context, ids ...string) (ents []*sap.StockPrediction, err error) {
	ents = make([]*sap.StockPrediction, len(ids))
	keys := make([]interface{}, len(ids))
	for i, id := range ids {
		keys[i] = struct {
			ID string `json:"id" `
		}{id}
	}

	if _, err = a.DB.QueryAndReceive(ctx, func(i int) []interface{} {
		ent := new(sap.StockPrediction)
		ents[i] = ent
		return []interface{}{
			&ent.ID, &ent.Created, &ent.StartDate, &ent.EndDate, &ent.RequesterID, &ent.StockCode,
			&ent.Holidays, &ent.CloseSeries, &ent.HighSeries, &ent.LowSeries,
		}
	}, `
		WITH __k AS (
			SELECT id, ROW_NUMBER() over() __idx
			FROM json_populate_recordset(null::gh.stock_prediction, $1)
		), sedp AS (
			SELECT stock_prediction_id, source,
				json_build_object('date', date, 'value', value) json_item
			FROM gh.series_daily_prediction sedp JOIN __k ON __k.id = sedp.stock_prediction_id
		), sep AS (
			SELECT
				stock_prediction_id, source, json_build_object(
					'created', created,
					'mae', mae,
					'mape', mape,
					'daily_predictions', (
						SELECT json_agg(json_item) FROM sedp
						WHERE sedp.stock_prediction_id = sep.stock_prediction_id
							AND sedp.source = sep.source
					)
				) json_item
			FROM gh.series_prediction sep JOIN __k ON __k.id = sep.stock_prediction_id
		)
		SELECT
			id, created, start_date, end_date, requester_id, stock_code,
			(
				SELECT json_agg(date) FROM gh.stock_prediction_holiday WHERE stock_prediction_id = stp.id
			) AS holidays,
			(
				SELECT json_item FROM sep WHERE sep.stock_prediction_id = stp.id AND sep.source = 'C'
			) AS close_series,
			(
				SELECT json_item FROM sep WHERE sep.stock_prediction_id = stp.id AND sep.source = 'H'
			) AS high_series,
			(
				SELECT json_item FROM sep WHERE sep.stock_prediction_id = stp.id AND sep.source = 'L'
			) AS low_series
		FROM gh.stock_prediction stp JOIN __k USING (id)
		ORDER BY __idx
	`, keys); err != nil {
		return nil, err
	}

	for i := 0; i < len(ids); i++ {
		if ents[i] == nil {
			break
		} else if ents[i].ID != ids[i] {
			copy(ents[i+1:], ents[i:])
			ents[i] = nil
		}
	}

	return ents, nil
}

var sqlPaginateStockPredictions = pgcc.RenderString(pgcc.Options{
	Cursor: "node.id",
	From: `(
		SELECT id, created FROM gh.stock_prediction stp
		WHERE ($5::UUID IS NULL OR stp.requester_id = $5)
	) AS node`,
	SortKeys: []pgcc.SortKey{
		{"node.created", "DESC"},
	},
})

// PaginateStockPredictions 메서드는 sap.StockPredictionFinder를 구현한다.
func (a *Adapter) PaginateStockPredictions(ctx context.Context, page list.PageOption, filter sap.PredictionFilter) (res list.Connection, err error) {
	_, err = a.DB.QueryAndReceive(ctx, res.Receive, sqlPaginateStockPredictions,
		page.First, page.After, page.Last, page.Before,
		filter.RequesterID,
	)
	return res, err
}
