package controllers

import "github.com/gin-gonic/gin"

// Init : various controller init
func Init(router *gin.RouterGroup) {

	var listControllerInit = []func(router *gin.RouterGroup){
		StockInit,
		UserInit,
	}

	for _, fn := range listControllerInit {
		fn(router)
	}
}
