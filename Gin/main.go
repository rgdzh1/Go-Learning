package main

import (
	"Go-Learning/Gin/controller"
	"Go-Learning/Gin/tool"
	"github.com/gin-gonic/gin"
)

/**
跨域访问:
	如果通信协议,主机和端口这三部分内容中得任意一个与原页面的不相同,就被称为跨域访问.
*/
func main() {
	config, err := tool.ParseConfig("./Gin/config/app.json")
	// 初始化xorm数据库引擎
	_, err = tool.InitOrmEngine(config)
	if err != nil {
		panic(err)
	}
	// 初始化Redis配置
	tool.InitRedisStore()
	engine := gin.Default()
	registerRouter(engine)
	// 集成Session
	tool.InitSession(engine)
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
