package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin

	// root:123qwe@tcp(127.0.0.1) 这一串是固定模式,账户名称root 密码:123qwe tcp链接 127.0.0.1本地数据库 ginsql 创建的数据库名称,如果没有的话还需要手动创建
	connStr := "root:123qwe@tcp(127.0.0.1)/ginsql"
	_, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	engine := gin.Default()

	// 监听8090端口

	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
