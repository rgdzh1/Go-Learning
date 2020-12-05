package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	// 设置HTML路径
	// .表示根目录
	engine.LoadHTMLGlob("./Gin_Study/html/*")
	// 设置静态资源
	// 当Html中访问/img目录时候,实际访问的是./Gin_Study/img目录,这里的.表示程序的根目录
	engine.Static("/img", "./Gin_Study/img")
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		// 返回HTML文件
		context.HTML(http.StatusOK, "index.html", gin.H{
			"path":  context.FullPath(),
			"title": "Gin 教程", // 网页标题栏改变名称
		})
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
