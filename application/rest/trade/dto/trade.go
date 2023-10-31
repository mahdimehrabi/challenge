package dto

import (
	"time"

	"challenge/internal/entity"
)

type TradeReq struct {
	ID           int       `json:"id"`
	InstrumentID int       `json:"instrumentId"`
	DateEn       time.Time `json:"dateEn"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
}

func (t *TradeReq) ToModel() *entity.Trade {
	return &entity.Trade{
		ID:           t.ID,
		InstrumentID: t.InstrumentID,
		DateEn:       t.DateEn,
		Open:         t.Open,
		High:         t.High,
		Low:          t.Low,
		Close:        t.Close,
	}
}
