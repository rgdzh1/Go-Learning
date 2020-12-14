package dao

import (
	"Go-Learning/Gin/model"
	"Go-Learning/Gin/tool"
)

type FoodCategoryDao struct {
	Engin *tool.OrmTool
}

// 实例化Dao对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DBEngine}
}

func (fcd FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	categories := new([]model.FoodCategory)
	err := fcd.Engin.Find(categories)
	if err != nil {
		return nil, err
	}
	return *categories, nil
}
