package stock

import (
	"tobi-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// Info : struct
type Info struct {
	Username string `json:"user"`
	Password int    `json:"pwd"`
}

// POSTStock POST Stock
func POSTStock(c *gin.Context) {
	req := new(Info)
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.JSON(200, handlers.OK(req))
}
