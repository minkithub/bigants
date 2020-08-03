package finder

import (
	"context"

	"github.com/allscape/bigants/grasshopper/pkg/pgcc"

	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

// GetStocksByCode 메서드는 sap.StockFinder를 구현한다.
func (a *Adapter) GetStocksByCode(ctx context.Context, codes ...string) (ents []*sap.Stock, err error) {
	ents = make([]*sap.Stock, len(codes))
	keys := make([]interface{}, len(codes))
	for i, id := range codes {
		keys[i] = struct {
			Code string `json:"code" `
		}{id}
	}

	if _, err = a.DB.QueryAndReceive(ctx, func(i int) []interface{} {
		ent := new(sap.Stock)
		ents[i] = ent
		return []interface{}{
			&ent.Code, &ent.NameKo, &ent.PredictablePeriodStart, &ent.PredictablePeriodEnd, &ent.RecentHistory,
		}
	}, `
		WITH __k AS (
			SELECT code, ROW_NUMBER() over() __idx
			FROM json_populate_recordset(null::stock, $1)
		), s AS (
			SELECT * FROM public.stock s JOIN __k USING (code)
			WHERE EXISTS(SELECT * FROM price WHERE price.stock_code_id  = code)
		)
		SELECT
			code, name_ko,
			(
				SELECT min(record_date) FROM public.price p WHERE p.stock_code_id = s.code
			) predictable_period_start,
			(
				SELECT max(record_date) FROM public.price p WHERE p.stock_code_id = s.code
			) predictable_period_end,
			(
				WITH latest_history AS (
					SELECT json_build_object(
						'date', record_date,
						'open', open_value,
						'close', close_value,
						'high', high_value,
						'low', low_value,
						'volume', volume
					) json_item
					FROM public.price WHERE stock_code_id = s.code
					ORDER BY record_date DESC
					LIMIT 2
				)
				SELECT json_agg(json_item) FROM latest_history
			) latest_history
		FROM s
		ORDER BY __idx
	`, keys); err != nil { // TODO: 히스토리가 없는 주식은 불러오지 않게 했음. 이게 맞는지 검토할 것
		return nil, err
	}

	for _, ent := range ents {
		for i, h := range ent.RecentHistory {
			if i < len(ent.RecentHistory)-1 {
				h.Previous = ent.RecentHistory[i+1]
			}
		}
	}

	for i := 0; i < len(codes); i++ {
		if ents[i] == nil {
			break
		} else if ents[i].Code != codes[i] {
			copy(ents[i+1:], ents[i:])
			ents[i] = nil
		}
	}

	return ents, nil
}

var sqlPaginateStocks = pgcc.RenderString(pgcc.Options{
	Cursor: "node.code",
	From: `(
		SELECT code FROM public.stock stock
		WHERE ($5::TEXT IS NULL OR stock.code ILIKE '%' || $5 || '%' OR stock.name_ko ILIKE '%' || $5 || '%')
	) AS node`,
	SortKeys: []pgcc.SortKey{
		{"node.code", "ASC"},
	},
})

// PaginateStocks 메서드는 sap.StockFinder를 구현한다.
func (a *Adapter) PaginateStocks(ctx context.Context, page list.PageOption, filter sap.StockFilter) (res list.Connection, err error) {
	_, err = a.DB.QueryAndReceive(ctx, res.Receive, sqlPaginateStocks,
		page.First, page.After, page.Last, page.Before,
		filter.Q,
	)
	return res, err
}

func (a *Adapter) FindStockHistory(ctx context.Context, code string, count *int32) (ents []*sap.StockHistory, err error) {
	if _, err = a.DB.QueryAndReceive(ctx, func(i int) []interface{} {
		ents = append(ents, new(sap.StockHistory))
		ent := ents[i]
		return []interface{}{
			&ent.Date, &ent.Open, &ent.Close, &ent.High, &ent.Low, &ent.Volume,
		}
	}, `
		WITH reversed_history AS (
			SELECT record_date, open_value, close_value, high_value, low_value, volume FROM public.price
			WHERE stock_code_id = $1
			ORDER BY record_date DESC
			LIMIT CASE WHEN $2::INT IS NULL THEN 500 ELSE LEAST(500, $2) END
		)
		SELECT * FROM reversed_history ORDER BY record_date ASC
	`, code, count); err != nil {
		return nil, err
	}
	return ents, nil
}
