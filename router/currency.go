package router

import (
	"asiayo/api/currencyApi"

	"github.com/gin-gonic/gin"
)

func currencyRouter(e *gin.Engine) {
	group := e.Group("/api/currency")

	group.POST("/transfer", currencyApi.Transfer)
}
