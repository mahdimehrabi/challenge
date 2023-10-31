package pgx

import (
	"context"

	"challenge/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Trade interface {
	Create(trade *entity.Trade) error
}
type trade struct {
	db *pgxpool.Pool
}

func NewTradeRepository(db *pgxpool.Pool) Trade {
	return &trade{
		db: db,
	}
}

func (r *trade) Create(trade *entity.Trade) error {
	sqlStatement := `
        INSERT INTO trade (Id,InstrumentId, DateEn, Open, High, Low, Close)
        VALUES ($1 ,$2, $3, $4, $5, $6, $7)
    `

	_, err := r.db.Exec(context.Background(), sqlStatement, trade.ID, trade.InstrumentID,
		trade.DateEn, trade.Open, trade.High, trade.Low, trade.Close)
	if err != nil {
		return err
	}

	return nil
}
