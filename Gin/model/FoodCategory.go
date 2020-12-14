package model

type FoodCategory struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Title       string `xorm:"varchar(20) title" json:"title"`
	Description string `xorm:"varchar(255) description" json:"description"`
	ImageUrl    string `xorm:"varchar(255) image_url" json:"image_url"`
	LinkUrl     string `xorm:"varchar(255) link_url" json:"link_url"`
	IsInServing bool   `xorm:"varchar(255) is_in_serving" json:"is_in_serving"`
}
