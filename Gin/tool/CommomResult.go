package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0 // 操作成功
	FAILED  = 1 // 失败
)

// 请求成功返回的抽取封装
func Success(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})
}

// 操作失败返回的内容
func Failed(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}
