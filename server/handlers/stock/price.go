package stock

import (
	"github.com/gin-gonic/gin"
)

// GETPrice 取得價錢
func GETPrice(c *gin.Context) {

	c.JSON(200, gin.H{
		"price": 10000,
	})
}
