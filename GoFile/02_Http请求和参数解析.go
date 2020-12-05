package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	// Handle方法处理
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println("全路劲", path)
		name := context.DefaultQuery("name", "xxx")
		fmt.Println("姓名为", name)
		// 输出
		context.Writer.Write([]byte("hello" + name))
	})
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.PostForm("name")
		pass := context.PostForm("pass")
		fmt.Println(name)
		fmt.Println(pass)
		context.Writer.Write([]byte(name + " " + pass + " 登录"))
	})
	// 分类处理
	engine.GET("h", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		//name:= context.Query("name")
		name := context.DefaultQuery("name", "张三")
		fmt.Println(name)
		context.Writer.Write([]byte(name))
	})
	engine.POST("/l", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name, exist := context.GetPostForm("name")
		if exist {
			fmt.Println(name)
		}
		pass, exist := context.GetPostForm("pass")
		if exist {
			fmt.Println(pass)
		}
		context.Writer.Write([]byte("hello" + name + " " + pass))
	})
	// 删除
	engine.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		userId := context.Param("id")
		fmt.Println(userId)
		context.Writer.Write([]byte("delete 用户id:=" + userId))
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
