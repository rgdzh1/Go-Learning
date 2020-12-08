package main

import (
	"Go-Learning/Gin/controller"
	"Go-Learning/Gin/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := tool.ParseConfig("./Gin/config/app.json")
	// 初始化xorm数据库引擎
	_, err = tool.InitOrmEngine(config)
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	registerRouter(engine)
	err = engine.Run(config.AppHost + ":" + config.AppPort)
	if err != nil {
		panic(err)
	}
}

// 路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
