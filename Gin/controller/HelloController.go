package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (receiver *HelloController) Router(engin *gin.Engine) {
	engin.GET("/hello", receiver.Hello)
}

// 解析
func (receiver *HelloController) Hello(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"code":    20,
		"message": "成功",
		"data":    "Router接口",
	})
}
