package models

import "github.com/jinzhu/gorm"

type Category struct {
	Model
	Name     string `gorm:"not null;type:varchar(16)"`
	ShowName string `gorm:"not null;type:varchar(10)"`
	MengerId int64  `gorm:"not null"`
	Status   byte   `gorm:"default:1"`
}


func (Category) TableName() string  {
	return "t_category"
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
