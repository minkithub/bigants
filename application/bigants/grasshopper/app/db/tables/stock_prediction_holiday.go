// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package tables

import (
	"context"
	"encoding/json"
	"time"
)

// StockPredictionHolidayRow represents a row for table "stock_prediction_holiday"
type StockPredictionHolidayRow struct {
	StockPredictionHolidayData
}

type StockPredictionHolidayData struct {
	StockPredictionID string    `json:"stock_prediction_id"`
	Date              time.Time `json:"date"`
}

func NewStockPredictionHolidayRow(stockPredictionID string, date time.Time) *StockPredictionHolidayRow {
	return &StockPredictionHolidayRow{StockPredictionHolidayData{stockPredictionID, date}}
}

func NewStockPredictionHolidayRows(data ...StockPredictionHolidayData) StockPredictionHolidayRows {
	rows := make(StockPredictionHolidayRows, len(data))
	for i, d := range data {
		rows[i] = &StockPredictionHolidayRow{d}
	}
	return rows
}

// GetStockPredictionID gets value of column "stock_prediction_id" from "stock_prediction_holiday" row
func (r *StockPredictionHolidayRow) GetStockPredictionID() string {
	return r.StockPredictionHolidayData.StockPredictionID
}

// SetStockPredictionID sets value of column "stock_prediction_id" in "stock_prediction_holiday" row
func (r *StockPredictionHolidayRow) SetStockPredictionID(stockPredictionID string) {
	r.StockPredictionHolidayData.StockPredictionID = stockPredictionID
}

// GetDate gets value of column "date" from "stock_prediction_holiday" row
func (r *StockPredictionHolidayRow) GetDate() time.Time { return r.StockPredictionHolidayData.Date }

// SetDate sets value of column "date" in "stock_prediction_holiday" row
func (r *StockPredictionHolidayRow) SetDate(date time.Time) { r.StockPredictionHolidayData.Date = date }

// StockPredictionHolidayRows represents multiple rows for table "stock_prediction_holiday"
type StockPredictionHolidayRows []*StockPredictionHolidayRow

func (r *StockPredictionHolidayRow) RefStockPredictionID() StockPredictionID {
	return StockPredictionID{r.GetStockPredictionID()}
}

func (rs StockPredictionHolidayRows) RefStockPredictionID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.RefStockPredictionID()
	}
	return keys
}

// NewStockPredictionHolidayTable(h SQLHandle) creates new StockPredictionHolidayTable
func NewStockPredictionHolidayTable(h SQLHandle) *StockPredictionHolidayTable {
	return &StockPredictionHolidayTable{h}
}

// StockPredictionHolidayTable provides access methods for table "stock_prediction_holiday"
type StockPredictionHolidayTable struct {
	h SQLHandle
}

func (t *StockPredictionHolidayTable) Find(ctx context.Context, filter StockPredictionHolidayValues) (StockPredictionHolidayRows, error) {
	return FindStockPredictionHolidayRows(ctx, t.h, filter)
}

func (t *StockPredictionHolidayTable) Count(ctx context.Context, filter StockPredictionHolidayValues) (int, error) {
	return CountStockPredictionHolidayRows(ctx, t.h, filter)
}

func (t *StockPredictionHolidayTable) Update(ctx context.Context, changeset, filter StockPredictionHolidayValues) (int64, error) {
	return UpdateStockPredictionHolidayRows(ctx, t.h, changeset, filter)
}

func (t *StockPredictionHolidayTable) Insert(ctx context.Context, rows ...*StockPredictionHolidayRow) (int, error) {
	return InsertReturningStockPredictionHolidayRows(ctx, t.h, rows...)
}

func (t *StockPredictionHolidayTable) Delete(ctx context.Context, filter StockPredictionHolidayValues) (int64, error) {
	return DeleteStockPredictionHolidayRows(ctx, t.h, filter)
}

type StockPredictionHolidayValues struct {
	StockPredictionID *string    `json:"stock_prediction_id"`
	Date              *time.Time `json:"date"`
}

// InsertStockPredictionHolidayRows inserts the rows into table "stock_prediction_holiday"
func InsertStockPredictionHolidayRows(ctx context.Context, db SQLHandle, rows ...*StockPredictionHolidayRow) (numRows int64, err error) {
	if len(rows) == 0 {
		return 0, nil
	}
	numRows, err = execWithJSONArgs(ctx, db, SQLInsertStockPredictionHolidayRows, rows)
	if err != nil {
		return numRows, formatError("InsertStockPredictionHolidayRows", err)
	}
	return numRows, nil
}

// InsertReturningStockPredictionHolidayRows inserts the rows into table "stock_prediction_holiday" and returns the rows.
func InsertReturningStockPredictionHolidayRows(ctx context.Context, db SQLHandle, inputs ...*StockPredictionHolidayRow) (numRows int, err error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	rows := StockPredictionHolidayRows(inputs)
	numRows, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLInsertReturningStockPredictionHolidayRows, rows)
	if err != nil {
		return numRows, formatError("InsertReturningStockPredictionHolidayRows", err)
	}
	return numRows, nil
}

// FindStockPredictionHolidayRows finds the rows matching the condition from table "stock_prediction_holiday"
func FindStockPredictionHolidayRows(ctx context.Context, db SQLHandle, cond StockPredictionHolidayValues) (rows StockPredictionHolidayRows, err error) {
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLFindStockPredictionHolidayRows, cond); err != nil {
		return nil, formatError("FindStockPredictionHolidayRows", err)
	}
	return rows, nil
}

// DeleteStockPredictionHolidayRows deletes the rows matching the condition from table "stock_prediction_holiday"
func DeleteStockPredictionHolidayRows(ctx context.Context, db SQLHandle, cond StockPredictionHolidayValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLDeleteStockPredictionHolidayRows, cond); err != nil {
		return numRows, formatError("DeleteStockPredictionHolidayRows", err)
	}
	return numRows, nil
}

func UpdateStockPredictionHolidayRows(ctx context.Context, db SQLHandle, changeset, filter StockPredictionHolidayValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLUpdateStockPredictionHolidayRows, changeset, filter); err != nil {
		return numRows, formatError("UpdateStockPredictionHolidayRows", err)
	}
	return numRows, nil
}

// CountStockPredictionHolidayRows counts the number of rows matching the condition from table "stock_prediction_holiday"
func CountStockPredictionHolidayRows(ctx context.Context, db SQLHandle, cond StockPredictionHolidayValues) (count int, err error) {
	if _, err = queryWithJSONArgs(ctx, db, func(int) []interface{} { return []interface{}{&count} }, SQLCountStockPredictionHolidayRows, cond); err != nil {
		return 0, formatError("CountStockPredictionHolidayRows", err)
	}
	return count, nil
}

// ReceiveRow returns all pointers of the column values for scanning
func (r *StockPredictionHolidayRow) ReceiveRow() []interface{} {
	return []interface{}{&r.StockPredictionHolidayData.StockPredictionID, &r.StockPredictionHolidayData.Date}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *StockPredictionHolidayRows) ReceiveRows(i int) []interface{} {
	if len(*rs) <= i {
		*rs = append(*rs, new(StockPredictionHolidayRow))
	} else if (*rs)[i] == nil {
		(*rs)[i] = new(StockPredictionHolidayRow)
	}
	return (*rs)[i].ReceiveRow()
}

func (r *StockPredictionHolidayRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.StockPredictionHolidayData)
}

var (
	SQLFindStockPredictionHolidayRows = `
		WITH __f AS (SELECT "stock_prediction_id", "date" FROM json_populate_record(null::"gh"."stock_prediction_holiday", $1))
		SELECT __t.stock_prediction_id, __t.date
		FROM "gh"."stock_prediction_holiday" AS __t
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))`
	SQLCountStockPredictionHolidayRows = `
		WITH __f AS (SELECT "stock_prediction_id", "date" FROM json_populate_record(null::"gh"."stock_prediction_holiday", $1))
		SELECT count(*) FROM "gh"."stock_prediction_holiday" AS __t
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))`
	SQLReturningStockPredictionHolidayRows = `
		RETURNING "stock_prediction_id", "date"`
	SQLInsertStockPredictionHolidayRows = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"gh"."stock_prediction_holiday", $1))
		INSERT INTO "gh"."stock_prediction_holiday" AS __t ("stock_prediction_id", "date")
		SELECT 
			__v."stock_prediction_id", 
			__v."date" FROM __v`
	SQLInsertReturningStockPredictionHolidayRows = SQLInsertStockPredictionHolidayRows + SQLReturningStockPredictionHolidayRows
	SQLDeleteStockPredictionHolidayRows          = `
		DELETE FROM "gh"."stock_prediction_holiday" AS __t
		WHERE TRUE
			AND (($1::json->>'stock_prediction_id' IS NULL) OR CAST($1::json->>'stock_prediction_id' AS uuid) = __t."stock_prediction_id")
			AND (($1::json->>'date' IS NULL) OR CAST($1::json->>'date' AS date) = __t."date")`
	SQLDeleteReturningStockPredictionHolidayRows = SQLDeleteStockPredictionHolidayRows + SQLReturningStockPredictionHolidayRows
	SQLUpdateStockPredictionHolidayRows          = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"gh"."stock_prediction_holiday", $1)),
			__f AS (SELECT * FROM json_populate_record(null::"gh"."stock_prediction_holiday", $2))
		UPDATE "gh"."stock_prediction_holiday" AS __t
		SET ("stock_prediction_id", "date") = (SELECT 
			COALESCE(__v."stock_prediction_id", __t."stock_prediction_id"), 
			COALESCE(__v."date", __t."date") FROM __v)
		WHERE ((SELECT __f."stock_prediction_id" IS NULL FROM __f) OR (SELECT __f."stock_prediction_id" = __t."stock_prediction_id" FROM __f))
			AND ((SELECT __f."date" IS NULL FROM __f) OR (SELECT __f."date" = __t."date" FROM __f))`
	SQLUpdateReturningStockPredictionHolidayRows = SQLUpdateStockPredictionHolidayRows + SQLReturningStockPredictionHolidayRows
)
