package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string
	MengerId int64
	Status   byte
}

func GetAllCategory() ([]*Category, error) {
	var categories []*Category
	err := db.Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(category *Category) error {
	err := db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}
