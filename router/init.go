package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func New() {
	engine := gin.Default()

	currencyRouter(engine)

	// 如環境參數未帶，則預設8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := engine.Run(":" + port); err != nil {
		panic(err)
	}
}
