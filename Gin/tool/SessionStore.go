package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(app *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}
	app.Use(sessions.Sessions("mysession", store))
}

/**
设置Session
*/
func SetSeess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

/**
获取Session
*/
func GetSeess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
