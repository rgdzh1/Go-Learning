package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// curl http://localhost:8090/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	// 返回Json数据
	engine.GET("/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println("请求路径", path)
		// 1. 普通的方式返回数据
		//context.Writer.Write([]byte(path))
		// 2. 将map数据转为json格式返回数据
		// 2xx:代表请求已经成功
		// 404:表示请求失败
		// 500:服务器内部异常
		// map[string]interface{}: 表示键为string类型值为任意类型的map
		//context.JSON(200, map[string]interface{}{
		//	"code":    2000,
		//	"message": "Ok",
		//	"data":    path,
		//})
		// 3. 将结构体转为Json数据返回
		context.JSON(http.StatusOK, Response{200, "成功", path})
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
