package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	// 使用路由组
	routeGroup := engine.Group("/user")
	routeGroup.GET("/register", registerUserHandle)
	routeGroup.POST("/login", loginUserHandle)
	routeGroup.DELETE("/:id", deleteUserHandle)
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}

func deleteUserHandle(context *gin.Context) {
	userId := context.Param("id")
	context.Writer.WriteString("用户id" + userId)
}
func loginUserHandle(context *gin.Context) {
	path := context.FullPath()
	context.Writer.WriteString(path)
}

func registerUserHandle(context *gin.Context) {
	path := context.FullPath()
	context.Writer.WriteString(path)
}
