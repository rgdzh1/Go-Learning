package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	// root:123qwe@tcp(127.0.0.1) 这一串是固定模式,账户名称root 密码:123qwe tcp链接 127.0.0.1本地数据库 ginsql 创建的数据库名称,如果没有的话还需要手动创建
	connStr := "root:123qwe@tcp(127.0.0.1)/ginsql"
	db, err1 := sql.Open("mysql", connStr)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	//创建表
	//_, err2 = db.Exec("CREATE TABLE person( id int auto_increment PRIMARY KEY, name VARCHAR(12) NOT NULL,age INT DEFAULT 1);")
	//插入数据
	//_, err2 := db.Exec("INSERT INTO person(name,age) VALUES(?,?);", "张三", 28)
	//查询数据
	result, err2 := db.Query("SELECT * FROM person;")
	if err2 != nil {
		log.Fatal(err2.Error())
		return
	}
scan:
	if result.Next() {
		person := new(Person)
		// &person.Id,表示取person.Id的地址值.
		err2 = result.Scan(&person.Id, &person.Name, &person.age)
		if err2 != nil {
			log.Fatal(err2.Error())
			return
		}
		fmt.Println("打印查询数据", person.Name, person.age, person.Id)
		goto scan
	}

	//// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	//engine := gin.Default()
	//// 监听8090端口
	//if err := engine.Run(":8090"); err != nil {
	//	log.Fatal(err.Error())
	//}
}

type Person struct {
	Id   int
	Name string
	age  int
}
