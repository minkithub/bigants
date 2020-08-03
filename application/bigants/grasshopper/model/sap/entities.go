package sap

import (
	"math"
	"time"

	"github.com/allscape/bigants/grasshopper/pkg/date"
)

// Stock Entity
type Stock struct {
	Code                   string          `json:"code"`
	NameKo                 string          `json:"name_ko"`
	RecentHistory          []*StockHistory `json:"recent_history"`
	PredictablePeriodStart date.Date       `json:"predictable_period_start"`
	PredictablePeriodEnd   date.Date       `json:"predictable_period_end"`
}

func (ent *Stock) LatestHistory() *StockHistory {
	if len(ent.RecentHistory) > 0 {
		return ent.RecentHistory[0]
	}
	return nil
}

type StockHistory struct {
	Date     date.Date     `json:"date"`
	Open     float64       `json:"open"`
	Close    float64       `json:"close"`
	High     float64       `json:"high"`
	Low      float64       `json:"low"`
	Volume   int32         `json:"volume"`
	Previous *StockHistory // 상대값(상한가, 하한가, 변화율)을 구하기 위해서
}

// HighLimit 메서드는 히스토리 날짜의 상한가를 구한다.
func (ent *StockHistory) HighLimit() *float64 {
	if ent.Previous == nil {
		return nil
	}
	val := ent.Previous.Close * 1.3
	return &val
}

// LowLimit 메서드는 히스토리 날짜의 하한가를 구한다.
func (ent *StockHistory) LowLimit() *float64 {
	if ent.Previous == nil {
		return nil
	}
	val := ent.Previous.Close * 0.7
	return &val
}

// PriceChange 메서드는 히스토리 날짜의 전일대비 가격차를 구한다.
func (ent *StockHistory) PriceChange() *float64 {
	if ent.Previous == nil {
		return nil
	}
	val := ent.Close - ent.Previous.Close
	return &val
}

// PriceChangeRate 메서드는 히스토리 날짜의 전일대비 가격비를 구한다.
func (ent *StockHistory) PriceChangeRate() *float64 {
	if ent.Previous == nil || ent.Previous.Close == 0 { // 0으로 나눌 수 없음
		return nil
	}
	val := float64(ent.Close) / float64(ent.Previous.Close)
	return &val
}

// StockPrediction Entity
type StockPrediction struct {
	ID          string            `json:"id"`
	Created     time.Time         `json:"created"`
	StartDate   date.Date         `json:"start_date"`
	EndDate     date.Date         `json:"end_date"`
	RequesterID *string           `json:"requester_id"`
	StockCode   string            `json:"stock_code"`
	Holidays    []date.Date       `json:"holidays"`
	LowSeries   *SeriesPrediction `json:"low_series"`
	HighSeries  *SeriesPrediction `json:"high_series"`
	CloseSeries *SeriesPrediction `json:"close_series"`
}

// AverageIncome 메서드는 (최고가 - 최저가) / 평균가 * 100 의 방식으로 예상평균수익률을 계산한다.
func (ent *StockPrediction) AverageIncome() float64 {
	var closeAvg float64
	for _, pred := range ent.CloseSeries.DailyPredictions {
		closeAvg += pred.Value / float64(len(ent.CloseSeries.DailyPredictions))
	}
	var max, min = .0, math.MaxFloat64
	for _, sep := range []*SeriesPrediction{ent.CloseSeries, ent.LowSeries, ent.HighSeries} {
		for _, pred := range sep.DailyPredictions {
			if pred.Value > max {
				max = pred.Value
			}
			if pred.Value < min {
				min = pred.Value
			}
		}
	}
	return (max - min) / closeAvg * 100
}

func (ent *StockPrediction) Accuracy() float64 {
	return 100 - ent.CloseSeries.MAPE
}
func (ent *StockPrediction) MAE() float64 {
	return ent.CloseSeries.MAE
}
func (ent *StockPrediction) DailyPredictions() []*MergedDailyPrediction {
	mpreds := make([]*MergedDailyPrediction, len(ent.CloseSeries.DailyPredictions))
	for i := range ent.CloseSeries.DailyPredictions {
		mpreds[i] = &MergedDailyPrediction{
			Date:  ent.CloseSeries.DailyPredictions[i].Date,
			Close: ent.CloseSeries.DailyPredictions[i].Value,
			High:  ent.HighSeries.DailyPredictions[i].Value,
			Low:   ent.LowSeries.DailyPredictions[i].Value,
		}
	}

	return mpreds
}

// SeriesPrediction 구조체는 한 계열의 예측결과를 나타낸다.
type SeriesPrediction struct {
	// ID               *int32             `json:"id"`
	Created          time.Time          `json:"created"`
	MAE              float64            `json:"mae"`
	MAPE             float64            `json:"mape"`
	DailyPredictions []*DailyPrediction `json:"daily_predictions"`
}

// DailyPrediction 구조체는 날짜별 값을 나타낸다.
type DailyPrediction struct {
	Date  date.Date `json:"date"`
	Value float64   `json:"value"`
}

type MergedDailyPrediction struct {
	Date  date.Date `json:"date"`
	High  float64   `json:"high"`
	Low   float64   `json:"low"`
	Close float64   `json:"close"`
}

func (ent *MergedDailyPrediction) ExpectedPrice() float64 {
	return ent.Close
}

func (ent *MergedDailyPrediction) ExpectedIncome() float64 {
	return (ent.High - ent.Low) / ent.Low * 100
}
