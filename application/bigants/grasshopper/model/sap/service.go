// Package sap (Stocks and Predictions)
// 주식정보와 예측을 다룬다.
package sap

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/pkg/date"
)

// Service 구조체는 주식정보와 예측에 대한 기능을 제공한다.
type Service struct {
	StockFinder
	StockPredictionFinder
	StockPredictionDepot
	Anthill Anthill
}

// StockFinder 인터페이스로 주식 정보를 읽어온다.
type StockFinder interface {
	GetStocksByCode(ctx context.Context, codes ...string) ([]*Stock, error)
	PaginateStocks(ctx context.Context, page list.PageOption, filter StockFilter) (list.Connection, error)
	// 종목의 가격 히스토리 조회 (범위는 적절하게)
	FindStockHistory(ctx context.Context, code string, count *int32) ([]*StockHistory, error)
}

// StockPredictionFinder 인터페이스로 예측 정보를 읽어온다.
type StockPredictionFinder interface {
	GetStockPredictions(ctx context.Context, ids ...string) ([]*StockPrediction, error)
	PaginateStockPredictions(ctx context.Context, page list.PageOption, filter PredictionFilter) (list.Connection, error)
}

type StockPredictionDepot interface {
	SaveStockPrediction(ctx context.Context, ent *StockPrediction) error
}

type Anthill interface { //
	PredictClose(ctx context.Context, req AnthillPredictRequest) (*SeriesPrediction, error)
	PredictHigh(ctx context.Context, req AnthillPredictRequest) (*SeriesPrediction, error)
	PredictLow(ctx context.Context, req AnthillPredictRequest) (*SeriesPrediction, error)
}

// StockFilter 구조체는 주식을 필터링하는 조건이다.
type StockFilter struct {
	Q *string
}

type StockHistoryFilter struct {
	StockCode string
}

// PredictionFilter 구조체는 예측을 필터링하는 조건이다.
type PredictionFilter struct {
	RequesterID *string
}

type AnthillPredictRequest struct {
	StockCode string
	Start     date.Date
	End       date.Date
	Holidays  []date.Date
}
