package model

import "time"

type Member struct {
	// pk autoincr 表示主键,自增
	// xorm:"phone" 这里面的列名用phone表示
	// json:"phone" 这里是用来将Json格式化为结构体用的
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	UserName     string    `xorm:"varchar(20) user_name" json:"user_name"`
	Mobile       string    `xorm:"varchar(11) mobile" json:"mobile"`
	Password     string    `xorm:"varchar(255) password" json:"password"`
	RegisterTime time.Time `xorm:"DATETIME" json:"register_time"`
	Avatar       string    `xorm:"varchar(255) avatar" json:"avatar"`
	Balance      float64   `xorm:"double balance" json:"balance"`
	IsActive     int8      `xorm:"tinyint" json:"is_active"`
	City         string    `xorm:"varchar(10) city" json:"city"`
}
