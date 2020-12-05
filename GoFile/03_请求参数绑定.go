package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	//`form:"name"` 很重要,不然没有值
	// Name 字段首字母要大写,不然也无法赋值
	Name string `form:"name"`
	Age  string `form:"age"`
}

func main() {
	// curl http://localhost:8080/hello 在终端输入这个请求,可以看到Gin返回了Hello Gin
	engine := gin.Default()
	// GET绑定数据
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		student := new(Student)
		err := context.ShouldBindQuery(student) // 将请求参数与结构体绑定,注意这里使用的是指针,如果使用值会浪费性能
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name)
		fmt.Println(student.Age)
		context.Writer.Write([]byte("hello" + student.Name + " " + student.Age))
	})
	// POST表单数据绑定
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		student := new(Student) // 结构体大写,form字段等很重要
		if err := context.ShouldBind(student); err != nil {
			log.Fatal(err.Error())
			return
		}
		context.Writer.Write([]byte("register " + student.Name))
	})
	// JOSN格式数据绑定
	engine.POST("/json", func(context *gin.Context) {
		/**
		请求数据格式:
		{
		"name":"小明",
		"age":"29"
		}
		*/
		fmt.Println(context.FullPath())
		student := new(Student)
		if err := context.BindJSON(student); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name)
		fmt.Println(student.Age)
		context.Writer.Write([]byte("返回数据" + student.Name))
	})
	// 监听8090端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
