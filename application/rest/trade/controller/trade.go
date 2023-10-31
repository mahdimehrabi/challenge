package controller

import (
	"challenge/application/rest/response"
	"challenge/application/rest/trade/dto"
	"challenge/internal/service"
	"github.com/gin-gonic/gin"
)

type TradeController struct {
	tradeService service.Trade
}

func NewTradeController(tradeService service.Trade) *TradeController {
	return &TradeController{
		tradeService: tradeService,
	}
}

func (c *TradeController) Create(gc *gin.Context) {
	req := new(dto.TradeReq)
	if err := gc.ShouldBindJSON(req); err != nil {
		response.RespondBadRequest(gc, nil)
		return
	}
	trade := req.ToModel()

	if err := c.tradeService.Create(trade); err != nil {
		response.RespondInternalServerError(gc, nil)
		return
	}

	response.RespondOK(gc, "Trade created successfully", nil)
}
