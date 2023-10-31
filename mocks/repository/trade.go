package repository

import (
	"challenge/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockTradeRepository struct {
	mock.Mock
}

func (m *MockTradeRepository) Create(trade *entity.Trade) error {
	args := m.Called(trade)
	return args.Error(0)
}
