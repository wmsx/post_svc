package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name   string
	Status byte
}

func GetAllCategory() ([]*Category, error) {
	var categories []*Category
	err := db.Find(&categories).Error
	if err != nil && err  != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categories, nil
}
