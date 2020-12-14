package controller

import (
	"Go-Learning/Gin/service"
	"Go-Learning/Gin/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (receiver FoodCategoryController) Router(engin *gin.Engine) {
	engin.GET("/api/food_category", receiver.foodCategory)
}

func (receiver FoodCategoryController) foodCategory(context *gin.Context) {
	fcService := new(service.FoodCategoryService)
	categories, err := fcService.Categories()
	if err != nil {
		tool.Failed(context, "食品种类数据获取失败")
		return
	}
	// imageUrl
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = "数据进行拼接"
		}
	}
	tool.Success(context, categories)
}
