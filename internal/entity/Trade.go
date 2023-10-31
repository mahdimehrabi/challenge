package entity

import "time"

type Trade struct {
	ID           int
	InstrumentID int
	DateEn       time.Time
	Open         float64
	High         float64
	Low          float64
	Close        float64
}
