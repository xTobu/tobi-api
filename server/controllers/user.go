package controllers

import (
	"tobi-api/server/handlers/user"

	"github.com/gin-gonic/gin"
)

// UserInit : for individual use
func UserInit(router *gin.RouterGroup) {
	groupUser := router.Group("user")
	{
		groupUser.GET(":id", user.GETUser)
		groupUser.POST("", user.POSTUser)
		groupUser.PUT(":id", user.PUTUser)
		groupUser.DELETE(":id", user.DELETEUser)
	}

	groupUsers := router.Group("users")
	{
		groupUsers.GET("", user.GETUsers)
		groupUsers.GET("double", user.GETDoubleUsers)
	}
}
