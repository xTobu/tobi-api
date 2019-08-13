package controllers

import (
	"tobi-api/server/handlers/stock"

	"github.com/gin-gonic/gin"
)

// StockInit : for individual use
func StockInit(router *gin.RouterGroup) {
	group := router.Group("stock")
	{
		group.POST("", stock.POSTStock)
		group.GET("", stock.GETPrice)
	}
}
