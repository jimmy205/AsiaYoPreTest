package currencyApi

import (
	"asiayo/business/currencyBin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	From   string  `json:"from" binding:"required"`
	To     string  `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"min=0"`
}

type ErrorResponse struct {
	ErrorText string `json:"error_text"`
}

type TransferResponse struct {
	Amount string `json:"amount"`
}

func Transfer(c *gin.Context) {

	// 利用gin的validator來判斷輸入的值
	// 可依輸入的值回傳不同的錯誤，做錯誤處理
	req := TransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, ErrorResponse{ErrorText: "參數錯誤，請再次確認。"})
		return
	}

	// 視情況 currency 可以考慮用inject進來的方式
	currency := currencyBin.NewCurrency()

	rate, err := currency.FindRate(req.From, req.To)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse{ErrorText: err.Error()})
		return
	}

	amount, err := currency.Transfer(rate, req.Amount)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse{ErrorText: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TransferResponse{Amount: amount})
}
