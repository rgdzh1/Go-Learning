package model

import "time"

type SmsCode struct {
	// pk autoincr 表示主键,自增
	// xorm:"phone" 这里面的列名用phone表示
	// json:"phone" 这里是用来将Json格式化为结构体用的
	Id         int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Phone      string    `xorm:"phone" json:"phone"`
	BizId      string    `xorm:"biz_id" json:"biz_id"`
	Code       string    `xorm:"code" json:"code"`
	CreateTime time.Time `xorm:"DATETIME"`
}
