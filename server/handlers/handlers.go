package handlers

import (
	"github.com/gin-gonic/gin"
)

// OK 統一用回傳
func OK(data interface{}) *gin.H {
	return &gin.H{
		"status": "success",
		"data":   data,
	}
}

// Fail 統一錯誤回傳
func Fail(msg interface{}) *gin.H {
	return &gin.H{
		"status": "fail",
		"msg":    msg,
	}
}
