package model

type Member struct {
	// pk autoincr 表示主键,自增
	// xorm:"phone" 这里面的列名用phone表示
	// json:"phone" 这里是用来将Json格式化为结构体用的
	Id           int    `xorm:"not null pk autoincr INT(11)" json:"id"`
	UserName     string `xorm:"user_name" json:"user_name"`
	Mobile       string `xorm:"mobile" json:"mobile"`
	Password     string `xorm:"password" json:"password"`
	RegisterTime string `xorm:"DATETIME" json:"register_time"`
	Avatar       string `xorm:"avatar" json:"avatar"`
	Balance      string `xorm:"balance" json:"balance"`
	IsActive     string `xorm:"is_active" json:"is_active"`
	City         string `xorm:"city" json:"city"`
}
