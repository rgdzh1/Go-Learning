package service

import (
	"Go-Learning/Gin/dao"
	"Go-Learning/Gin/model"
)

type FoodCategoryService struct {
}

func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	categoryDao := dao.NewFoodCategoryDao()
	return categoryDao.QueryCategories()
}
