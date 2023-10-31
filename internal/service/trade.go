package service

import (
	"fmt"

	"challenge/internal/entity"
	"challenge/internal/repository/trade/pgx"
)

type Trade interface {
	Create(trade *entity.Trade) error
}

type trade struct {
	tradeRepo pgx.Trade
}

func NewTradeService(tradeRepo Trade) Trade {
	return &trade{
		tradeRepo: tradeRepo,
	}
}

func (s *trade) Create(trade *entity.Trade) error {
	if err := s.tradeRepo.Create(trade); err != nil {
		// we must use a logging package like zap here
		fmt.Printf("log error:%s", err)
		return err
	}
	return nil
}
