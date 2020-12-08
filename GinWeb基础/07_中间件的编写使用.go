package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	// engine.Use(RequstInfosMiddleware())// 第一种,中间件的使用
	// 第二种中间件的使用,对单个方法的使用
	engine.GET("/hello", RequstInfosMiddleware(), func(context *gin.Context) {
		context.JSON(404, map[string]interface{}{
			"code":    200,
			"message": "请求成功",
			"data":    context.FullPath(),
		})
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}

/**
中间件方法封装
*/
func RequstInfosMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("请求路劲", context.FullPath())
		fmt.Println("请求方法", context.Request.Method)
		fmt.Println("状态码", context.Writer.Status()) // 请求还未完成,状态码此时是200

		context.Next() // 执行到这里暂停往下执行,等网络请求处理完成再继续向下走

		fmt.Println("状态码", context.Writer.Status())
	}
}
