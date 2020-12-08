package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("Hello Gin\n"))
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
