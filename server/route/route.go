package route

import (
	"tobi-api/server/controllers"
	"tobi-api/server/middleware"

	"github.com/gin-gonic/gin"
)

// Init :
func Init() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	{
		controllers.StockInit(api)
		controllers.UserInit(api)
		// controllers.Init(api) // 保留控制器的彈性，故註解
	}

	auth := router.Group("auth")
	{
		auth.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"auth": true,
			})
		})

		auth.GET("mw", middleware.Auth(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"auth": true,
			})
		})
	}

	return router

}
