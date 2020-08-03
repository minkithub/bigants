// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package tables

import (
	"context"
	"encoding/json"
	"time"
)

// SeriesDailyPredictionRow represents a row for table "series_daily_prediction"
type SeriesDailyPredictionRow struct {
	SeriesDailyPredictionData
}

type SeriesDailyPredictionData struct {
	StockPredictionID string    `json:"stock_prediction_id"`
	Source            string    `json:"source"`
	Date              time.Time `json:"date"`
	Value             float64   `json:"value"`
}

func NewSeriesDailyPredictionRow(stockPredictionID string, source string, date time.Time, value float64) *SeriesDailyPredictionRow {
	return &SeriesDailyPredictionRow{SeriesDailyPredictionData{stockPredictionID, source, date, value}}
}

func NewSeriesDailyPredictionRows(data ...SeriesDailyPredictionData) SeriesDailyPredictionRows {
	rows := make(SeriesDailyPredictionRows, len(data))
	for i, d := range data {
		rows[i] = &SeriesDailyPredictionRow{d}
	}
	return rows
}

// GetStockPredictionID gets value of column "stock_prediction_id" from "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) GetStockPredictionID() string {
	return r.SeriesDailyPredictionData.StockPredictionID
}

// SetStockPredictionID sets value of column "stock_prediction_id" in "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) SetStockPredictionID(stockPredictionID string) {
	r.SeriesDailyPredictionData.StockPredictionID = stockPredictionID
}

// GetSource gets value of column "source" from "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) GetSource() string { return r.SeriesDailyPredictionData.Source }

// SetSource sets value of column "source" in "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) SetSource(source string) {
	r.SeriesDailyPredictionData.Source = source
}

// GetDate gets value of column "date" from "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) GetDate() time.Time { return r.SeriesDailyPredictionData.Date }

// SetDate sets value of column "date" in "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) SetDate(date time.Time) { r.SeriesDailyPredictionData.Date = date }

// GetValue gets value of column "value" from "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) GetValue() float64 { return r.SeriesDailyPredictionData.Value }

// SetValue sets value of column "value" in "series_daily_prediction" row
func (r *SeriesDailyPredictionRow) SetValue(value float64) { r.SeriesDailyPredictionData.Value = value }

// SeriesDailyPredictionRows represents multiple rows for table "series_daily_prediction"
type SeriesDailyPredictionRows []*SeriesDailyPredictionRow

func (r *SeriesDailyPredictionRow) RefStockPredictionID() StockPredictionID {
	return StockPredictionID{r.GetStockPredictionID()}
}

func (rs SeriesDailyPredictionRows) RefStockPredictionID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.RefStockPredictionID()
	}
	return keys
}

// NewSeriesDailyPredictionTable(h SQLHandle) creates new SeriesDailyPredictionTable
func NewSeriesDailyPredictionTable(h SQLHandle) *SeriesDailyPredictionTable {
	return &SeriesDailyPredictionTable{h}
}

// SeriesDailyPredictionTable provides access methods for table "series_daily_prediction"
type SeriesDailyPredictionTable struct {
	h SQLHandle
}

func (t *SeriesDailyPredictionTable) Find(ctx context.Context, filter SeriesDailyPredictionValues) (SeriesDailyPredictionRows, error) {
	return FindSeriesDailyPredictionRows(ctx, t.h, filter)
}

func (t *SeriesDailyPredictionTable) Count(ctx context.Context, filter SeriesDailyPredictionValues) (int, error) {
	return CountSeriesDailyPredictionRows(ctx, t.h, filter)
}

func (t *SeriesDailyPredictionTable) Update(ctx context.Context, changeset, filter SeriesDailyPredictionValues) (int64, error) {
	return UpdateSeriesDailyPredictionRows(ctx, t.h, changeset, filter)
}

func (t *SeriesDailyPredictionTable) Insert(ctx context.Context, rows ...*SeriesDailyPredictionRow) (int, error) {
	return InsertReturningSeriesDailyPredictionRows(ctx, t.h, rows...)
}

func (t *SeriesDailyPredictionTable) Delete(ctx context.Context, filter SeriesDailyPredictionValues) (int64, error) {
	return DeleteSeriesDailyPredictionRows(ctx, t.h, filter)
}

type SeriesDailyPredictionValues struct {
	StockPredictionID *string    `json:"stock_prediction_id"`
	Source            *string    `json:"source"`
	Date              *time.Time `json:"date"`
	Value             *float64   `json:"value"`
}

// InsertSeriesDailyPredictionRows inserts the rows into table "series_daily_prediction"
func InsertSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, rows ...*SeriesDailyPredictionRow) (numRows int64, err error) {
	if len(rows) == 0 {
		return 0, nil
	}
	numRows, err = execWithJSONArgs(ctx, db, SQLInsertSeriesDailyPredictionRows, rows)
	if err != nil {
		return numRows, formatError("InsertSeriesDailyPredictionRows", err)
	}
	return numRows, nil
}

// InsertReturningSeriesDailyPredictionRows inserts the rows into table "series_daily_prediction" and returns the rows.
func InsertReturningSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, inputs ...*SeriesDailyPredictionRow) (numRows int, err error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	rows := SeriesDailyPredictionRows(inputs)
	numRows, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLInsertReturningSeriesDailyPredictionRows, rows)
	if err != nil {
		return numRows, formatError("InsertReturningSeriesDailyPredictionRows", err)
	}
	return numRows, nil
}

// FindSeriesDailyPredictionRows finds the rows matching the condition from table "series_daily_prediction"
func FindSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, cond SeriesDailyPredictionValues) (rows SeriesDailyPredictionRows, err error) {
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLFindSeriesDailyPredictionRows, cond); err != nil {
		return nil, formatError("FindSeriesDailyPredictionRows", err)
	}
	return rows, nil
}

// DeleteSeriesDailyPredictionRows deletes the rows matching the condition from table "series_daily_prediction"
func DeleteSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, cond SeriesDailyPredictionValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLDeleteSeriesDailyPredictionRows, cond); err != nil {
		return numRows, formatError("DeleteSeriesDailyPredictionRows", err)
	}
	return numRows, nil
}

func UpdateSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, changeset, filter SeriesDailyPredictionValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLUpdateSeriesDailyPredictionRows, changeset, filter); err != nil {
		return numRows, formatError("UpdateSeriesDailyPredictionRows", err)
	}
	return numRows, nil
}

// CountSeriesDailyPredictionRows counts the number of rows matching the condition from table "series_daily_prediction"
func CountSeriesDailyPredictionRows(ctx context.Context, db SQLHandle, cond SeriesDailyPredictionValues) (count int, err error) {
	if _, err = queryWithJSONArgs(ctx, db, func(int) []interface{} { return []interface{}{&count} }, SQLCountSeriesDailyPredictionRows, cond); err != nil {
		return 0, formatError("CountSeriesDailyPredictionRows", err)
	}
	return count, nil
}

// ReceiveRow returns all pointers of the column values for scanning
func (r *SeriesDailyPredictionRow) ReceiveRow() []interface{} {
	return []interface{}{&r.SeriesDailyPredictionData.StockPredictionID, &r.SeriesDailyPredictionData.Source, &r.SeriesDailyPredictionData.Date, &r.SeriesDailyPredictionData.Value}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *SeriesDailyPredictionRows) ReceiveRows(i int) []interface{} {
	if len(*rs) <= i {
		*rs = append(*rs, new(SeriesDailyPredictionRow))
	} else if (*rs)[i] == nil {
		(*rs)[i] = new(SeriesDailyPredictionRow)
	}
	return (*rs)[i].ReceiveRow()
}

func (r *SeriesDailyPredictionRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.SeriesDailyPredictionData)
}

var (
	SQLFindSeriesDailyPredictionRows = `
		WITH __f AS (SELECT "stock_prediction_id", "source", "date", "value" FROM json_populate_record(null::"gh"."series_daily_prediction", $1))
		SELECT __t.stock_prediction_id, __t.source, __t.date, __t.value
		FROM "gh"."series_daily_prediction" AS __t
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."source" IS NULL FROM __f) OR (SELECT __f."source" = __t."source" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))
			AND ((SELECT __f."value" IS NULL FROM __f) OR (SELECT __f."value" = __t."value" FROM __f))`
	SQLCountSeriesDailyPredictionRows = `
		WITH __f AS (SELECT "stock_prediction_id", "source", "date", "value" FROM json_populate_record(null::"gh"."series_daily_prediction", $1))
		SELECT count(*) FROM "gh"."series_daily_prediction" AS __t
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."source" IS NULL FROM __f) OR (SELECT __f."source" = __t."source" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))
			AND ((SELECT __f."value" IS NULL FROM __f) OR (SELECT __f."value" = __t."value" FROM __f))`
	SQLReturningSeriesDailyPredictionRows = `
		RETURNING "stock_prediction_id", "source", "date", "value"`
	SQLInsertSeriesDailyPredictionRows = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"gh"."series_daily_prediction", $1))
		INSERT INTO "gh"."series_daily_prediction" AS __t ("stock_prediction_id", "source", "date", "value")
		SELECT 
			__v."stock_prediction_id", 
			__v."source", 
			__v."date", 
			__v."value" FROM __v`
	SQLInsertReturningSeriesDailyPredictionRows = SQLInsertSeriesDailyPredictionRows + SQLReturningSeriesDailyPredictionRows
	SQLDeleteSeriesDailyPredictionRows          = `
		DELETE FROM "gh"."series_daily_prediction" AS __t
		WHERE TRUE
			AND (($1::json->>'stock_prediction_id' IS NULL) OR CAST($1::json->>'stock_prediction_id' AS uuid) = __t."stock_prediction_id")
			AND (($1::json->>'source' IS NULL) OR CAST($1::json->>'source' AS character varying) = __t."source")
			AND (($1::json->>'date' IS NULL) OR CAST($1::json->>'date' AS date) = __t."date")
			AND (($1::json->>'value' IS NULL) OR CAST($1::json->>'value' AS numeric) = __t."value")`
	SQLDeleteReturningSeriesDailyPredictionRows = SQLDeleteSeriesDailyPredictionRows + SQLReturningSeriesDailyPredictionRows
	SQLUpdateSeriesDailyPredictionRows          = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"gh"."series_daily_prediction", $1)),
			__f AS (SELECT * FROM json_populate_record(null::"gh"."series_daily_prediction", $2))
		UPDATE "gh"."series_daily_prediction" AS __t
		SET ("stock_prediction_id", "source", "date", "value") = (SELECT 
			COALESCE(__v."stock_prediction_id", __t."stock_prediction_id"), 
			COALESCE(__v."source", __t."source"), 
			COALESCE(__v."date", __t."date"), 
			COALESCE(__v."value", __t."value") FROM __v)
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."source" IS NULL FROM __f) OR (SELECT __f."source" = __t."source" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))
			AND ((SELECT __f."value" IS NULL FROM __f) OR (SELECT __f."value" = __t."value" FROM __f))`
	SQLUpdateReturningSeriesDailyPredictionRows = SQLUpdateSeriesDailyPredictionRows + SQLReturningSeriesDailyPredictionRows
)
