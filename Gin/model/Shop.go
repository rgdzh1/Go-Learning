package model

type Shop struct {
	// id
	Id int64 `xorm:"pk autoincr" json:"id"`
	// 店铺名称
	Name string `xorm:"varchar(12) name" json:"name"`
	// 宣传信息
	PromotionInfo string `xorm:"varchar(30) promotion_info" json:"promotion_info"`
	// 地址
	Address string `xorm:"varchar(100) address" json:"address"`
	// 联系电话
	Phone string `xorm:"varchar(11) phone" json:"phone"`
	// 店铺营业状态
	Status string `xorm:"tinyint status" json:"status"`
	// 经度
	Longitude float64 `xorm:"float longitude" json:"longitude"`
	// 纬度
	Latitude float64 `xorm:"float latitude" json:"latitude"`
	// 店铺图标
	ImagePath string `xorm:"varchar(255) image_path" json:"image_path"`

	IsNew bool `xorm:"bool is_new" json:"is_new"`
	//
	IsPremium bool `xorm:"bool is_premium" json:"is_premium"`

	// 商铺评分
	Rating float32 `xorm:"float rating" json:"rating"`
	// 评分总数
	RatingCount int64 `xorm:"int rating_count" json:"rating_count"`
	// 当前订单总数
	RecentOrderNum int64 `xorm:"int recent_order_num" json:"recent_order_num"`

	// 配送起送价
	MinimumOrederAmount int32 `xorm:"int minimum_oreder_amount" json:"minimum_oreder_amount"`
	// 配送费
	DeliveryFee int32 `xorm:"int delivery_fee" json:"delivery_fee"`

	// 营业时间
	OpeningHours string `xorm:"varchar(20) opening_hours" json:"opening_hours"`
}
