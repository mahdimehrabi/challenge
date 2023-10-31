package service

import (
	"errors"
	"testing"
	"time"

	"challenge/internal/entity"
	"challenge/mocks/repository"

	"github.com/stretchr/testify/assert"
)

func TestTradeService_Create(t *testing.T) {
	mockRepo := new(repository.MockTradeRepository)

	service := NewTradeService(mockRepo)

	testTrade := &entity.Trade{
		InstrumentID: 4,
		Close:        22,
		ID:           1,
		DateEn:       time.Now(),
	}

	mockRepo.On("Create", testTrade).Return(nil)
	err := service.Create(testTrade)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTradeService_Create_RepoError(t *testing.T) {
	mockRepo := new(repository.MockTradeRepository)

	service := NewTradeService(mockRepo)

	testTrade := &entity.Trade{
		InstrumentID: 4,
		Close:        22,
		ID:           1,
		DateEn:       time.Now(),
	}

	mockRepo.On("Create", testTrade).Return(errors.New("error"))
	err := service.Create(testTrade)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
